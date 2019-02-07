package datatypes_test

import (
	"math"

	"github.com/ff14wed/xivnet/v2/datatypes"
)

var movementBlockBytes = []byte{
	0x12, 0x12, 0x67, 0x45, 0x00, 0x00, // Direction, U1, U2
	0xAB, 0x89, 0xAB, 0x89, 0xAB, 0x89, // PackedPosition
	0x67, 0x45, 0x00, 0x00, // U3
}

var expectedMovementBlockData = &datatypes.Movement{
	Direction: 0x12,
	U1:        0x12,
	U2:        0x4567,
	Position:  datatypes.PackedPosition{X: 0x89AB, Y: 0x89AB, Z: 0x89AB},
	U3:        0x4567,
}

var myMovementBlockBytes = []byte{
	219, 15, 73, 64, // Direction
	0x67, 0x45, 0x00, 0x00, // U1
	0xAB, 0x89, 0x00, 0x00, // U2
	0, 0, 250, 67, // X
	0, 0, 22, 68, // Y
	0, 0, 47, 68, // Z
}

var expectedMyMovementBlockData = &datatypes.MyMovement{
	Direction: math.Pi,
	U1:        0x4567,
	U2:        0x89AB,
	X:         500,
	Y:         600,
	Z:         700,
}
