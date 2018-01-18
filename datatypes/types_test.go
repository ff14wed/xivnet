package datatypes_test

import (
	"math"

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
		var (
			movementBytes    []byte
			expectedMovement *datatypes.Movement
		)
		BeforeEach(func() {
			movementBytes = []byte{
				0x12, 0x12, 0x67, 0x45, 0x00, 0x00, // Direction, U1, U2
				0xAB, 0x89, 0xAB, 0x89, 0xAB, 0x89, // PackedPosition
				0x67, 0x45, 0x00, 0x00, // U3
			}
			expectedMovement = &datatypes.Movement{
				Direction: 0x12,
				U1:        0x12,
				U2:        0x4567,
				Position:  datatypes.PackedPosition{X: 0x89AB, Z: 0x89AB, Y: 0x89AB},
				U3:        0x4567,
			}
		})
		It("successfully parses unmarshals block data into a Movement struct", func() {
			bd := datatypes.NewBlockData(datatypes.MovementOpcode, false)
			Expect(bd).ToNot(BeNil())
			Expect(bd).To(BeAssignableToTypeOf(new(datatypes.Movement)))
			Expect(datatypes.UnmarshalBlockBytes(movementBytes, bd)).To(Succeed())
			Expect(bd).To(Equal(expectedMovement))
		})
		Context("when the length of the data doesn't match the target struct", func() {
			It("returns a failure", func() {
				bd := datatypes.NewBlockData(datatypes.MovementOpcode, false)
				Expect(bd).ToNot(BeNil())
				Expect(bd).To(BeAssignableToTypeOf(new(datatypes.Movement)))
				err := datatypes.UnmarshalBlockBytes(append(movementBytes, 0x56), bd)
				Expect(err).To(MatchError("length mismatch: 17 != 16"))
			})
		})
	})
	Describe("UnmarshalBlockBytes (MyMovement opcode)", func() {
		var (
			myMovementBytes    []byte
			expectedMyMovement *datatypes.MyMovement
		)
		BeforeEach(func() {
			myMovementBytes = []byte{
				219, 15, 73, 64, // Direction
				0x67, 0x45, 0x00, 0x00, // U1
				0xAB, 0x89, 0x00, 0x00, // U2
				0, 0, 250, 67, // X
				0, 0, 22, 68, // Z
				0, 0, 47, 68, // Y
			}
			expectedMyMovement = &datatypes.MyMovement{
				Direction: math.Pi,
				U1:        0x4567,
				U2:        0x89AB,
				X:         500,
				Z:         600,
				Y:         700,
			}
		})
		It("successfully parses unmarshals block data into a MyMovement struct", func() {
			bd := datatypes.NewBlockData(datatypes.MyMovementOpcode, true)
			Expect(bd).ToNot(BeNil())
			Expect(bd).To(BeAssignableToTypeOf(new(datatypes.MyMovement)))
			Expect(datatypes.UnmarshalBlockBytes(myMovementBytes, bd)).To(Succeed())
			Expect(bd).To(Equal(expectedMyMovement))
		})
	})
})
