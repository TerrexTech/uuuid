package uuuid

import (
	"crypto/rand"
	"net"
	"sync"

	"github.com/gocql/gocql"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

// uuidGen is the generator for UUID.
var uuidGen *uuid.Gen

// uuidMutex is used for thread-safe uuidGen assignment
var uuidMutex sync.RWMutex

// UUID represents a v4 UUID.
type UUID struct {
	uuid.UUID
}

// NewV4 returns a new v4 UUID
func NewV4() (UUID, error) {
	if uuidGen == nil {
		// We don't place this outside "if" condition because even if the condition is executed
		// two times, it doesn't have any negative impact as far as uuinGen doesn't get
		// concurrent writes. Our aim is just to make sure this is not nil.
		uuidMutex.Lock()
		uuidGen = uuid.NewGenWithHWAF(randomHWAddrFunc)
		uuidMutex.Unlock()
	}

	uid, err := uuidGen.NewV4()
	if err != nil {
		return UUID{}, err
	}
	return UUID{uid}, nil
}

// NewV1 returns a new v1 UUID
func NewV1() (UUID, error) {
	if uuidGen == nil {
		// We don't place this outside "if" condition because even if the condition is executed
		// two times, it doesn't have any negative impact as far as uuinGen doesn't get
		// concurrent writes. Our aim is just to make sure this is not nil.
		uuidMutex.Lock()
		uuidGen = uuid.NewGenWithHWAF(randomHWAddrFunc)
		uuidMutex.Unlock()
	}

	uid, err := uuidGen.NewV1()
	if err != nil {
		return UUID{}, err
	}
	return UUID{uid}, nil
}

// MarshalCQL converts the uuid into GoCql-compatible []byte.
func (u UUID) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	return u.MarshalBinary()
}

// UnmarshalCQL converts GoCql UUID-type to local UUID.
func (u *UUID) UnmarshalCQL(info gocql.TypeInfo, data []byte) error {
	return u.UnmarshalBinary(data)
}

// TimestampFromV1 returns the Timestamp embedded within a V1 UUID.
// Returns an error if the UUID is any version other than 1.
func TimestampFromV1(u UUID) (uuid.Timestamp, error) {
	return uuid.TimestampFromV1(u.UUID)
}

// randomHWAddrFunc generates a random MAC address for V1-UUID.
func randomHWAddrFunc() (net.HardwareAddr, error) {
	// From: https://stackoverflow.com/questions/21018729/generate-mac-address-in-go
	addr := make([]byte, 6)
	_, err := rand.Read(addr)
	if err != nil {
		err = errors.Wrap(err, "Error generating random MAC-Addr")
		return []byte{}, err
	}
	// Set local bit, ensure unicast address
	addr[0] = (addr[0] | 2) & 0xfe
	return addr, nil
}
