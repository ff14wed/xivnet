package datatypes

import "github.com/ff14wed/xivnet/v3"

var outTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of outgoing (to server) network blocks
const (
	EgressClientTriggerOpcode = 0x1BE // Updated for 6.38

	EgressChatZoneOpcode = 0x1F2 // Updated for 6.38

	EgressMovementOpcode         = 0x25F // Updated for 6.38
	EgressInstanceMovementOpcode = 0x1A0 // Updated for 6.38

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
