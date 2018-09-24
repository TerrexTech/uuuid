package uuuid

import (
	"testing"

	"github.com/gocql/gocql"
	"github.com/gofrs/uuid"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUUID(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UUID Suite")
}

// uuidType is a mock gocql-type for testing CQL marshalling/Unmarshalling
type uuidType struct{}

func (u uuidType) Type() gocql.Type {
	return gocql.TypeUUID
}

func (u uuidType) Version() byte {
	return 1
}
func (u uuidType) Custom() string {
	return "custom"
}

func (u uuidType) New() interface{} {
	return uuidType{}
}

var _ = Describe("UUID", func() {
	Context("new V4 UUID is requested", func() {
		It("should return new UUID", func() {
			u, err := NewV4()
			Expect(err).ToNot(HaveOccurred())

			uid := u.String()
			Expect(uid).ToNot(BeEmpty())
			_, err = uuid.FromString(uid)
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Describe("Marshalling", func() {
		It("should marshal to CQL", func() {
			uid, err := NewV4()
			Expect(err).ToNot(HaveOccurred())

			marshalled, err := gocql.Marshal(uuidType{}, uid)
			Expect(err).ToNot(HaveOccurred())

			newUUID, err := uuid.FromBytes(marshalled)
			Expect(err).ToNot(HaveOccurred())
			Expect(newUUID.String()).To(Equal(uid.String()))
		})
	})

	Describe("Unmarshalling", func() {
		It("should unmarshal to CQL", func() {
			uid, err := NewV4()
			Expect(err).ToNot(HaveOccurred())

			uidBytes := uid.Bytes()

			unmarshal := UUID{}
			err = gocql.Unmarshal(uuidType{}, uidBytes, &unmarshal)
			Expect(err).ToNot(HaveOccurred())
			Expect(unmarshal.String()).To(Equal(uid.String()))
		})
	})

	Context("parsing to UUID", func() {
		Describe("FromBytes", func() {
			It("should parse bytes to UUID", func() {
				u, err := uuid.NewV4()
				Expect(err).ToNot(HaveOccurred())
				bytes := u.Bytes()

				uid, err := FromBytes(bytes)
				Expect(err).ToNot(HaveOccurred())
				_, err = uuid.FromBytes(uid.Bytes())
				Expect(err).ToNot(HaveOccurred())
			})

			It("should return any errors that occur", func() {
				invalidUUID := "invalid"

				_, err := FromBytes([]byte(invalidUUID))
				Expect(err).To(HaveOccurred())
			})
		})

		Describe("FromBytesOrNil", func() {
			It("should parse bytes to UUID", func() {
				u, err := uuid.NewV4()
				Expect(err).ToNot(HaveOccurred())
				bytes := u.Bytes()

				uid := FromBytesOrNil(bytes)
				Expect(uid).ToNot(BeNil())
				_, err = uuid.FromBytes(uid.Bytes())
				Expect(err).ToNot(HaveOccurred())
			})

			It("should return nil if any errors occur", func() {
				invalidUUID := "invalid"

				uid := FromBytesOrNil([]byte(invalidUUID))
				emptyUUID := (UUID{}).String()
				Expect(uid.String()).To(Equal(emptyUUID))
			})
		})

		Describe("FromString", func() {
			It("should parse string to UUID", func() {
				u, err := uuid.NewV4()
				Expect(err).ToNot(HaveOccurred())
				str := u.String()

				uid, err := FromString(str)
				Expect(uid).ToNot(BeNil())
				_, err = uuid.FromString(uid.String())
				Expect(err).ToNot(HaveOccurred())
			})

			It("should return any errors that occur", func() {
				invalidUUID := "invalid"

				_, err := FromString(invalidUUID)
				Expect(err).To(HaveOccurred())
			})
		})

		Describe("FromStringOrNil", func() {
			It("should parse string to UUID", func() {
				u, err := uuid.NewV4()
				Expect(err).ToNot(HaveOccurred())
				str := u.String()

				uid, err := FromString(str)
				Expect(uid).ToNot(BeNil())
				_, err = uuid.FromString(uid.String())
				Expect(err).ToNot(HaveOccurred())
			})

			It("should return nil if any errors occur", func() {
				invalidUUID := "invalid"

				uid := FromStringOrNil(invalidUUID)
				emptyUUID := (UUID{}).String()
				Expect(uid.String()).To(Equal(emptyUUID))
			})
		})
	})
})
