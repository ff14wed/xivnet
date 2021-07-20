package datatypes

import "github.com/ff14wed/xivnet/v3"

var outTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of outgoing (to server) network blocks
const (
	EgressClientTriggerOpcode = 0x8B // Updated for 5.58

	EgressChatZoneOpcode = 0x2F4 // Updated for 5.58

	EgressMovementOpcode         = 0x212 // Updated for 5.58
	EgressInstanceMovementOpcode = 0x1A3 // Updated for 5.58

	EgressPerformOpcode    = UndefinedOpcode
	EgressCraftEventOpcode = UndefinedOpcode
)

func init() {
	registerOutBlockFactory(EgressClientTriggerOpcode, func() xivnet.BlockData { return new(EgressClientTrigger) })

	registerOutBlockFactory(EgressChatZoneOpcode, func() xivnet.BlockData { return new(EgressChatZone) })

	registerOutBlockFactory(EgressMovementOpcode, func() xivnet.BlockData { return new(EgressMovement) })
	registerOutBlockFactory(EgressInstanceMovementOpcode, func() xivnet.BlockData { return new(EgressInstanceMovement) })

	registerOutBlockFactory(EgressPerformOpcode, func() xivnet.BlockData { return new(Perform) })

	registerOutBlockFactory(EgressCraftEventOpcode, func() xivnet.BlockData { return new(EgressCraftEvent) })
}

func registerOutBlockFactory(opcode uint16, factory func() xivnet.BlockData) {
	outTypeRegistry[opcode] = factory
}
