package uuuid

import (
	"github.com/gofrs/uuid"
)

// FromBytes returns a UUID generated from the raw byte slice input.
// It will return an error if the slice isn't 16 bytes long.
func FromBytes(input []byte) (uuid.UUID, error) {
	return uuid.FromBytes(input)
}

// FromBytesOrNil returns a UUID generated from the raw byte slice input.
// Same behavior as FromBytes(), but returns uuid.Nil instead of an error.
func FromBytesOrNil(input []byte) uuid.UUID {
	return uuid.FromBytesOrNil(input)
}

// FromString returns a UUID parsed from the input string.
// Input is expected in a form accepted by UnmarshalText.
func FromString(input string) (uuid.UUID, error) {
	return uuid.FromString(input)
}

// FromStringOrNil returns a UUID parsed from the input string.
// Same behavior as FromString(), but returns uuid.Nil instead of an error.
func FromStringOrNil(input string) uuid.UUID {
	return uuid.FromStringOrNil(input)
}
