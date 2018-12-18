package datatypes_test

import (
	"github.com/ff14wed/xivnet"
	"github.com/ff14wed/xivnet/datatypes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Types", func() {
	Describe("NewBlockData", func() {
		It("returns nil if the type doesn't exist", func() {
			bd := datatypes.NewBlockData(0, false)
			Expect(bd).To(BeNil())
		})
	})

	Describe("UnmarshalBlockBytes (Movement opcode)", func() {
		It("successfully parses unmarshals block data into a Movement struct", func() {
			bd := datatypes.NewBlockData(datatypes.MovementOpcode, false)
			Expect(bd).ToNot(BeNil())
			Expect(bd).To(BeAssignableToTypeOf(new(datatypes.Movement)))
			Expect(datatypes.UnmarshalBlockBytes(movementBlockBytes, bd)).To(Succeed())
			Expect(bd).To(Equal(expectedMovementBlockData))
		})

		Context("when the length of the data doesn't match the target struct", func() {
			It("returns a failure", func() {
				bd := datatypes.NewBlockData(datatypes.MovementOpcode, false)
				Expect(bd).ToNot(BeNil())
				Expect(bd).To(BeAssignableToTypeOf(new(datatypes.Movement)))
				err := datatypes.UnmarshalBlockBytes(append(movementBlockBytes, 0x56), bd)
				Expect(err).To(MatchError("length mismatch: 17 != 16"))
			})
		})
	})
	Describe("UnmarshalBlockBytes (MyMovement opcode)", func() {
		It("successfully parses unmarshals block data into a MyMovement struct", func() {
			bd := datatypes.NewBlockData(datatypes.MyMovementOpcode, true)
			Expect(bd).ToNot(BeNil())
			Expect(bd).To(BeAssignableToTypeOf(new(datatypes.MyMovement)))
			Expect(datatypes.UnmarshalBlockBytes(myMovementBlockBytes, bd)).To(Succeed())
			Expect(bd).To(Equal(expectedMyMovementBlockData))
		})
	})

	Describe("ParseBlock", func() {
		It("correctly parses raw incoming blocks", func() {
			b := &xivnet.Block{
				Length: 12345,
				Header: xivnet.BlockHeader{
					SubjectID: 12345,
					CurrentID: 67890,
					Opcode:    datatypes.MovementOpcode,
				},
				Data: xivnet.GenericBlockDataFromBytes(movementBlockBytes),
			}

			newB, err := datatypes.ParseBlock(b, false)
			Expect(err).ToNot(HaveOccurred())
			Expect(newB).ToNot(Equal(b))
			Expect(newB).To(Equal(&xivnet.Block{
				Length: 12345,
				Header: xivnet.BlockHeader{
					SubjectID: 12345,
					CurrentID: 67890,
					Opcode:    datatypes.MovementOpcode,
				},
				Data: expectedMovementBlockData,
			}))
		})

		It("correctly parses raw outgoing blocks", func() {
			b := &xivnet.Block{
				Length: 12345,
				Header: xivnet.BlockHeader{
					SubjectID: 12345,
					CurrentID: 67890,
					Opcode:    datatypes.MyMovementOpcode,
				},
				Data: xivnet.GenericBlockDataFromBytes(myMovementBlockBytes),
			}

			newB, err := datatypes.ParseBlock(b, true)
			Expect(err).ToNot(HaveOccurred())
			Expect(newB).ToNot(Equal(b))
			Expect(newB).To(Equal(&xivnet.Block{
				Length: 12345,
				Header: xivnet.BlockHeader{
					SubjectID: 12345,
					CurrentID: 67890,
					Opcode:    datatypes.MyMovementOpcode,
				},
				Data: expectedMyMovementBlockData,
			}))
		})

		It("leaves unknown blocks alone", func() {
			b := &xivnet.Block{
				Length: 12345,
				Header: xivnet.BlockHeader{
					SubjectID: 12345,
					CurrentID: 67890,
					Opcode:    0x9999,
				},
				Data: xivnet.GenericBlockDataFromBytes([]byte("abcdefg")),
			}

			newB, err := datatypes.ParseBlock(b, true)
			Expect(err).ToNot(HaveOccurred())
			Expect(newB).To(Equal(b))
		})

		It("leaves already parsed blocks alone", func() {
			b := &xivnet.Block{
				Length: 12345,
				Header: xivnet.BlockHeader{
					SubjectID: 12345,
					CurrentID: 67890,
					Opcode:    datatypes.MovementOpcode,
				},
				Data: expectedMovementBlockData,
			}

			newB, err := datatypes.ParseBlock(b, false)
			Expect(err).ToNot(HaveOccurred())
			Expect(newB).To(Equal(b))
		})

		It("returns the untouched block and an error if the parsing failed", func() {
			b := &xivnet.Block{
				Length: 12345,
				Header: xivnet.BlockHeader{
					SubjectID: 12345,
					CurrentID: 67890,
					Opcode:    datatypes.MovementOpcode,
				},
				Data: xivnet.GenericBlockDataFromBytes([]byte("abcdefg")),
			}

			newB, err := datatypes.ParseBlock(b, false)
			Expect(err).To(MatchError("length mismatch: 7 != 16"))
			Expect(newB).To(Equal(b))
		})
	})
})
