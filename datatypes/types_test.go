package datatypes_test

import (
	"github.com/ff14wed/xivnet/v3"
	"github.com/ff14wed/xivnet/v3/datatypes"
	. "github.com/onsi/ginkgo/v2"
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
	Describe("UnmarshalBlockBytes (EgressMovement opcode)", func() {
		It("successfully parses unmarshals block data into a EgressMovement struct", func() {
			bd := datatypes.NewBlockData(datatypes.EgressMovementOpcode, true)
			Expect(bd).ToNot(BeNil())
			Expect(bd).To(BeAssignableToTypeOf(new(datatypes.EgressMovement)))
			Expect(datatypes.UnmarshalBlockBytes(egressMovementBlockBytes, bd)).To(Succeed())
			Expect(bd).To(Equal(expectedEgressMovementBlockData))
		})
	})

	Describe("ParseBlock", func() {
		It("correctly parses raw incoming blocks", func() {
			b := &xivnet.Block{
				Length:    12345,
				SubjectID: 12345,
				CurrentID: 67890,
				IPCHeader: xivnet.IPCHeader{
					ServerID: 123,
					Opcode:   datatypes.MovementOpcode,
				},
				Data: xivnet.GenericBlockDataFromBytes(movementBlockBytes),
			}

			newB, err := datatypes.ParseBlock(b, false)
			Expect(err).ToNot(HaveOccurred())
			Expect(newB).ToNot(Equal(b))
			Expect(newB).To(Equal(&xivnet.Block{
				Length:    12345,
				SubjectID: 12345,
				CurrentID: 67890,
				IPCHeader: xivnet.IPCHeader{
					ServerID: 123,
					Opcode:   datatypes.MovementOpcode,
				},
				Data: expectedMovementBlockData,
			}))
		})

		It("correctly parses raw outgoing blocks", func() {
			b := &xivnet.Block{
				Length:    12345,
				SubjectID: 12345,
				CurrentID: 67890,
				IPCHeader: xivnet.IPCHeader{
					ServerID: 123,
					Opcode:   datatypes.EgressMovementOpcode,
				},
				Data: xivnet.GenericBlockDataFromBytes(egressMovementBlockBytes),
			}

			newB, err := datatypes.ParseBlock(b, true)
			Expect(err).ToNot(HaveOccurred())
			Expect(newB).ToNot(Equal(b))
			Expect(newB).To(Equal(&xivnet.Block{
				Length:    12345,
				SubjectID: 12345,
				CurrentID: 67890,
				IPCHeader: xivnet.IPCHeader{
					ServerID: 123,
					Opcode:   datatypes.EgressMovementOpcode,
				},
				Data: expectedEgressMovementBlockData,
			}))
		})

		It("correctly parses ServerID == 0 blocks as Chat blocks", func() {
			By("ensuring that what would otherwise be parsed as a movement block is not parsed")
			b := &xivnet.Block{
				Length:    12345,
				SubjectID: 12345,
				CurrentID: 67890,
				IPCHeader: xivnet.IPCHeader{
					Opcode: datatypes.MovementOpcode,
				},
				Data: xivnet.GenericBlockDataFromBytes(movementBlockBytes),
			}

			newB, err := datatypes.ParseBlock(b, true)
			Expect(err).ToNot(HaveOccurred())
			Expect(newB).To(Equal(b))

			By("ensuring chat blocks are parsed correctly")
			b2 := &xivnet.Block{
				Length:    12345,
				SubjectID: 12345,
				CurrentID: 67890,
				IPCHeader: xivnet.IPCHeader{
					Opcode: datatypes.ChatOpcode,
				},
				Data: xivnet.GenericBlockDataFromBytes(chatBlockBytes),
			}

			newB2, err := datatypes.ParseBlock(b2, false)
			Expect(err).ToNot(HaveOccurred())
			Expect(newB2).ToNot(Equal(b))
			Expect(newB2).To(Equal(&xivnet.Block{
				Length:    12345,
				SubjectID: 12345,
				CurrentID: 67890,
				IPCHeader: xivnet.IPCHeader{
					Opcode: datatypes.ChatOpcode,
				},
				Data: expectedChatData,
			}))
		})

		It("leaves unknown blocks alone", func() {
			b := &xivnet.Block{
				Length:    12345,
				SubjectID: 12345,
				CurrentID: 67890,
				IPCHeader: xivnet.IPCHeader{
					ServerID: 123,
					Opcode:   0x9999,
				},
				Data: xivnet.GenericBlockDataFromBytes([]byte("abcdefg")),
			}

			newB, err := datatypes.ParseBlock(b, true)
			Expect(err).ToNot(HaveOccurred())
			Expect(newB).To(Equal(b))
		})

		It("leaves already parsed blocks alone", func() {
			b := &xivnet.Block{
				Length:    12345,
				SubjectID: 12345,
				CurrentID: 67890,
				IPCHeader: xivnet.IPCHeader{
					ServerID: 123,
					Opcode:   datatypes.MovementOpcode,
				},
				Data: expectedMovementBlockData,
			}

			newB, err := datatypes.ParseBlock(b, false)
			Expect(err).ToNot(HaveOccurred())
			Expect(newB).To(Equal(b))
		})

		It("returns the untouched block and an error if the parsing failed", func() {
			b := &xivnet.Block{
				Length:    12345,
				SubjectID: 12345,
				CurrentID: 67890,
				IPCHeader: xivnet.IPCHeader{
					ServerID: 123,
					Opcode:   datatypes.MovementOpcode,
				},
				Data: xivnet.GenericBlockDataFromBytes([]byte("abcdefg")),
			}

			newB, err := datatypes.ParseBlock(b, false)
			Expect(err).To(MatchError("length mismatch: 7 != 16"))
			Expect(newB).To(Equal(b))
		})
	})
})
