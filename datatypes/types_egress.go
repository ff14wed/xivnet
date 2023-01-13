package datatypes

import "github.com/ff14wed/xivnet/v3"

var outTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of outgoing (to server) network blocks
const (
	EgressClientTriggerOpcode = 0xFC // Updated for 6.3

	EgressChatZoneOpcode = 0x140 // Updated for 6.3

	EgressMovementOpcode         = 0x19C // Updated for 6.3
	EgressInstanceMovementOpcode = 0x1F6 // Updated for 6.3

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
