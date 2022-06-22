package datatypes_test

import (
	"math"

	"github.com/ff14wed/xivnet/v3/datatypes"
)

var movementBlockBytes = []byte{
	0x12, 0x12, 0x67, 0x45, 0x01, 0x02, // Direction, U1, U2
	0xAB, 0x89, 0xAB, 0x89, 0xAB, 0x89, // PackedPosition
	0x67, 0x45, 0x00, 0x00, // U3
}

var expectedMovementBlockData = &datatypes.Movement{
	HeadRotation:    0x12,
	Direction:       0x12,
	AnimationType:   0x67,
	AnimationState:  0x45,
	AnimationSpeed:  0x01,
	UnknownRotation: 0x02,
	Position:        datatypes.PackedPosition{X: 0x89AB, Y: 0x89AB, Z: 0x89AB},
	U3:              0x4567,
}

var egressMovementBlockBytes = []byte{
	219, 15, 73, 64, // Direction
	0x67, 0x45, 0x00, 0x00, // U1
	0, 0, 250, 67, // X
	0, 0, 22, 68, // Y
	0, 0, 47, 68, // Z
	0xAB, 0x89, 0x00, 0x00, // U2
}

var expectedEgressMovementBlockData = &datatypes.EgressMovement{
	Direction: math.Pi,
	U1:        0x4567,
	U2:        0x89AB,
	X:         500,
	Y:         600,
	Z:         700,
}

var eventPlay32BlockBytes = []byte{
	0x15, 0xCD, 0x5B, 0x07, 0x00, 0x00, 0x00, 0x00, // ActorID
	0x01, 0x00, 0x00, 0x00, // EventID
	0x02, 0x00, 0x03, 0x00, // Scene and Pad1
	0x04, 0x00, 0x00, 0x00, // Flags
	0x05, 0x00, 0x00, 0x00, // P1
	0x06, 0x00, 0x00, 0x00, // ParamCount and Pad2
	0x07, 0x00, 0x00, 0x00, // P2

	// EventPlay32Data
	0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0xFF, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
	0xFF, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00,
	0x00, 0x01, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00,
	0x00, 0x03, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00,
	0x05, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00,
	0xF9, 0xFF, 0xFF, 0xFF, 0x08, 0x00, 0x00, 0x00,
	0x09, 0x00, 0x00, 0x00,

	0x01, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00,
	0x03, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00,
	0x05, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00,
	0x07, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00,
	0x09, 0x00, 0x00, 0x00, 0x0A, 0x00, 0x00, 0x00,
	0x0B, 0x00, 0x00, 0x00, 0x0C, 0x00, 0x00, 0x00,
	0x0D, 0x00, 0x00, 0x00, 0x0E, 0x00, 0x00, 0x00,
	0x0F, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00,
	0x11, 0x00, 0x00, 0x00,
}

var expectedEventPlay32BlockData = &datatypes.EventPlay32{
	EventPlayHeader: datatypes.EventPlayHeader{
		ActorID: 123456789,
		EventID: 1,
		Scene:   2, Pad1: 3,
		Flags: 4, P1: 5, ParamCount: 6, P2: 7,
	},
	Data: datatypes.GenericEventPlay32Data([32]uint32{
		8, 0, 0xFF, 1, 0xFF, 2, 0x100, 0x200, 0x300,
		0x400, 5, 6, 0xFFFFFFF9, 8, 9,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 0xA, 0xB, 0xC, 0xD,
		0xE, 0xF, 0x10, 0x11,
	}),
}

var expectedEventPlay32BlockDataJSON = `
{
	"ActorID":123456789,
	"EventID": 1,
	"Scene": 2, "Pad1": 3,
	"Flags": 4, "P1": 5, "ParamCount": 6,
	"Pad2": [0, 0, 0], "P2": 7,
	"Data": [8, 0, 255, 1, 255, 2, 256, 512, 768, 1024, 5, 6, 4294967289, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17]
}`

var expectedCraftState = datatypes.CraftState{
	U1:                8,
	U3:                0,
	U4:                255,
	CraftAction:       1,
	U2:                255,
	StepNum:           2,
	Progress:          256,
	ProgressDelta:     512,
	Quality:           768,
	QualityDelta:      1024,
	HQChance:          5,
	Durability:        6,
	DurabilityDelta:   -7,
	CurrentCondition:  8,
	PreviousCondition: 9,
	Flags:             1,
	U6:                [16]uint32{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17},
}

var expectedCraftStateBlockDataJSON = `
{
	"ActorID":123456789,
	"EventID": 655361,
	"Scene": 2, "Pad1": 3,
	"Flags": 4, "P1": 5, "ParamCount": 6,
	"Pad2": [0, 0, 0], "P2": 7,
	"Data": {
		"U1":                8,
		"U3":                0,
		"U4":                255,
		"CraftAction":       1,
		"U2":                255,
		"StepNum":           2,
		"Progress":          256,
		"ProgressDelta":     512,
		"Quality":           768,
		"QualityDelta":      1024,
		"HQChance":          5,
		"Durability":        6,
		"DurabilityDelta":   -7,
		"CurrentCondition":  8,
		"PreviousCondition": 9,
		"Flags": 1,
		"U6": [2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17]
	}
}`

var chatBlockBytes = func() []byte {
	first := []byte{
		0x89, 0x67, 0x45, 0x23, 0x01, 0x00, 0x00, 0x00, // ChannelID
		0x9A, 0x78, 0x56, 0x34, 0x02, 0x00, 0x00, 0x00, // SpeakerCharacterID
		0x9A, 0x78, 0x56, 0x34, // SpeakerEntityID
		0x12, 0x00, // WorldID
		0x34, // Flags
	}
	entName := datatypes.StringToEntityName("Test Char")
	chatMsg := datatypes.StringToChatMessage("hello")
	return append(append(append(first, entName[:]...), chatMsg[:]...), 0)
}()

var expectedChatData = &datatypes.Chat{
	ChannelID:          0x123456789,
	SpeakerCharacterID: 0x23456789A,
	SpeakerEntityID:    0x3456789A,
	WorldID:            0x12,
	Flags:              0x34,
	SpeakerName:        datatypes.StringToEntityName("Test Char"),
	Message:            datatypes.StringToChatMessage("hello"),
}
