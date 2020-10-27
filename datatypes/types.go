package datatypes

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/ff14wed/xivnet/v3"
)

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)
var outTypeRegistry = make(map[uint16]func() xivnet.BlockData)

const UndefinedOpcode = 0xFFFF

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x35E // Updated for 5.35
	InitZoneOpcode      = 0x303 // Updated for 5.35
	ControlOpcode       = 0x2DC // Updated for 5.35
	ControlSelfOpcode   = 0x32C // Updated for 5.35
	ControlTargetOpcode = 0x369 // Updated for 5.35
	RemoveEntityOpcode  = 0xBC  // Updated for 5.35
	UpdateHPMPTPOpcode  = 0x153 // Updated for 5.35

	ChatZoneOpcode = 0x39F // Updated for 5.35

	UpdateStatusesOpcode       = 0x3A8 // Updated for 5.35
	UpdateStatusesEurekaOpcode = 0x2D3 // Updated for 5.35
	UpdateStatusesBossOpcode   = 0x28C // Updated for 5.35

	ActionOpcode      = 0x3A9 // Updated for 5.35
	AoEAction8Opcode  = 0x2B3 // Updated for 5.35
	AoEAction16Opcode = 0x3D7 // Updated for 5.35
	AoEAction24Opcode = 0x1AB // Updated for 5.35
	AoEAction32Opcode = 0x258 // Updated for 5.35

	PlayerSpawnOpcode = 0x38E // Updated for 5.35
	NPCSpawnOpcode    = 0x1DA // Updated for 5.35
	NPCSpawn2Opcode   = 0x346 // Updated for 5.35

	MovementOpcode = 0x2C5 // Updated for 5.35
	SetPosOpcode   = 0x1D4 // Updated for 5.35

	CastingOpcode = 0x37B // Updated for 5.35

	HateRankingOpcode = 0xA9  // Updated for 5.35
	HateListOpcode    = 0x3AE // Updated for 5.35

	EquipChangeOpcode = 0x1A2 // Updated for 5.35

	EventPlayOpcode   = 0x39A // Updated for 5.35
	EventPlay4Opcode  = 0x382 // Updated for 5.35
	EventPlay32Opcode = 0x115 // Updated for 5.35

	MountOpcode = 0x3D4 // Updated for 5.35

	WeatherChangeOpcode = 0x1AA // Updated for 5.35

	PrepareZoningOpcode = 0x160 // Updated for 5.35

	GaugeOpcode = 0xEB // Updated for 5.35

	WaymarkOpcode         = UndefinedOpcode
	PerformOpcode         = UndefinedOpcode
	XWorldPartyListOpcode = UndefinedOpcode
)

func init() {
	registerInBlockFactory(EffectResultOpcode, func() xivnet.BlockData { return new(EffectResult) })
	registerInBlockFactory(InitZoneOpcode, func() xivnet.BlockData { return new(InitZone) })
	registerInBlockFactory(ControlOpcode, func() xivnet.BlockData { return new(Control) })
	registerInBlockFactory(ControlSelfOpcode, func() xivnet.BlockData { return new(ControlSelf) })
	registerInBlockFactory(ControlTargetOpcode, func() xivnet.BlockData { return new(ControlTarget) })
	registerInBlockFactory(RemoveEntityOpcode, func() xivnet.BlockData { return new(RemoveEntity) })
	registerInBlockFactory(UpdateHPMPTPOpcode, func() xivnet.BlockData { return new(UpdateHPMPTP) })

	registerInBlockFactory(ChatZoneOpcode, func() xivnet.BlockData { return new(ChatZone) })

	registerInBlockFactory(UpdateStatusesOpcode, func() xivnet.BlockData { return new(UpdateStatuses) })
	registerInBlockFactory(UpdateStatusesEurekaOpcode, func() xivnet.BlockData { return new(UpdateStatusesEureka) })
	registerInBlockFactory(UpdateStatusesBossOpcode, func() xivnet.BlockData { return new(UpdateStatusesBoss) })

	registerInBlockFactory(ActionOpcode, func() xivnet.BlockData { return new(Action) })
	registerInBlockFactory(AoEAction8Opcode, func() xivnet.BlockData { return new(AoEAction8) })
	registerInBlockFactory(AoEAction16Opcode, func() xivnet.BlockData { return new(AoEAction16) })
	registerInBlockFactory(AoEAction24Opcode, func() xivnet.BlockData { return new(AoEAction24) })
	registerInBlockFactory(AoEAction32Opcode, func() xivnet.BlockData { return new(AoEAction32) })

	registerInBlockFactory(PlayerSpawnOpcode, func() xivnet.BlockData { return new(PlayerSpawn) })
	registerInBlockFactory(NPCSpawnOpcode, func() xivnet.BlockData { return new(NPCSpawn) })
	registerInBlockFactory(NPCSpawn2Opcode, func() xivnet.BlockData { return new(NPCSpawn2) })

	registerInBlockFactory(MovementOpcode, func() xivnet.BlockData { return new(Movement) })
	registerInBlockFactory(SetPosOpcode, func() xivnet.BlockData { return new(SetPos) })
	registerInBlockFactory(CastingOpcode, func() xivnet.BlockData { return new(Casting) })

	registerInBlockFactory(HateRankingOpcode, func() xivnet.BlockData { return new(HateRanking) })
	registerInBlockFactory(HateListOpcode, func() xivnet.BlockData { return new(HateList) })

	registerInBlockFactory(EquipChangeOpcode, func() xivnet.BlockData { return new(EquipChange) })

	registerInBlockFactory(EventPlayOpcode, func() xivnet.BlockData { return new(EventPlay) })
	registerInBlockFactory(EventPlay4Opcode, func() xivnet.BlockData { return new(EventPlay4) })
	registerInBlockFactory(EventPlay32Opcode, func() xivnet.BlockData { return new(EventPlay32) })

	registerInBlockFactory(MountOpcode, func() xivnet.BlockData { return new(Mount) })

	registerInBlockFactory(WeatherChangeOpcode, func() xivnet.BlockData { return new(WeatherChange) })

	registerInBlockFactory(WaymarkOpcode, func() xivnet.BlockData { return new(Marker) })
	registerInBlockFactory(PrepareZoningOpcode, func() xivnet.BlockData { return new(PrepareZoning) })

	registerInBlockFactory(GaugeOpcode, func() xivnet.BlockData { return new(Gauge) })
	registerInBlockFactory(PerformOpcode, func() xivnet.BlockData { return new(Perform) })

	registerInBlockFactory(XWorldPartyListOpcode, func() xivnet.BlockData { return new(XWorldPartyList) })
}

// Opcodes that define the datatypes of outgoing (to server) network blocks
const (
	EgressClientTriggerOpcode = 0x1DF // Updated for 5.35

	EgressChatZoneOpcode = 0x21C // Updated for 5.35

	EgressMovementOpcode         = 0xBD // Updated for 5.35
	EgressInstanceMovementOpcode = 0x81 // Updated for 5.35

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

var inChatTypeRegistry = make(map[uint16]func() xivnet.BlockData)
var outChatTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming network blocks for chat
const (
	ChatFromOpcode          = 0x64
	ChatOpcode              = 0x65
	ChatFromXWorldOpcode    = 0x6F
	FreeCompanyResultOpcode = 0x12C
	ChatXWorldOpcode        = 0x72
)

func init() {
	registerInChatBlockFactory(ChatFromOpcode, func() xivnet.BlockData { return new(ChatFrom) })
	registerInChatBlockFactory(ChatOpcode, func() xivnet.BlockData { return new(Chat) })
	registerInChatBlockFactory(ChatFromXWorldOpcode, func() xivnet.BlockData { return new(ChatFromXWorld) })
	registerInChatBlockFactory(FreeCompanyResultOpcode, func() xivnet.BlockData { return new(FreeCompanyResult) })
	registerInChatBlockFactory(ChatXWorldOpcode, func() xivnet.BlockData { return new(ChatXWorld) })
}

// Opcodes that define the datatypes of outgoing network blocks for chat
const (
	ChatToOpcode           = 0x64
	EgressChatOpcode       = 0x65
	EgressChatXWorldOpcode = 0x6d
)

func init() {
	registerOutChatBlockFactory(ChatToOpcode, func() xivnet.BlockData { return new(ChatTo) })
	registerOutChatBlockFactory(EgressChatOpcode, func() xivnet.BlockData { return new(EgressChat) })
	registerOutChatBlockFactory(EgressChatXWorldOpcode, func() xivnet.BlockData { return new(EgressChatXWorld) })
}

func newBlockData(opcode uint16, registry map[uint16]func() xivnet.BlockData) xivnet.BlockData {
	factory, ok := registry[opcode]
	if !ok {
		return nil
	}
	return factory()
}

// NewBlockData is a factory for BlockData that uses the opcode to
// determine which BlockData to create
func NewBlockData(opcode uint16, isOut bool) xivnet.BlockData {
	r := inTypeRegistry
	if isOut {
		r = outTypeRegistry
	}
	return newBlockData(opcode, r)
}

// NewChatBlockData is a factory for BlockData that uses the opcode to
// determine which Chat BlockData to create
func NewChatBlockData(opcode uint16, isOut bool) xivnet.BlockData {
	r := inChatTypeRegistry
	if isOut {
		r = outChatTypeRegistry
	}
	return newBlockData(opcode, r)
}

type BlockUnmarshaler interface {
	UnmarshalBytes(data []byte) error
}

// UnmarshalBlockBytes converts raw bytes to this block data struct
func UnmarshalBlockBytes(data []byte, block xivnet.BlockData) error {
	if b, matches := block.(BlockUnmarshaler); matches {
		return b.UnmarshalBytes(data)
	}
	if binary.Size(block) != len(data) {
		return fmt.Errorf("length mismatch: %d != %d", len(data), binary.Size(block))
	}
	return binary.Read(bytes.NewReader(data), binary.LittleEndian, block)
}

func registerInBlockFactory(opcode uint16, factory func() xivnet.BlockData) {
	inTypeRegistry[opcode] = factory
}

func registerOutBlockFactory(opcode uint16, factory func() xivnet.BlockData) {
	outTypeRegistry[opcode] = factory
}

func registerInChatBlockFactory(opcode uint16, factory func() xivnet.BlockData) {
	inChatTypeRegistry[opcode] = factory
}

func registerOutChatBlockFactory(opcode uint16, factory func() xivnet.BlockData) {
	outChatTypeRegistry[opcode] = factory
}

// ParseBlock takes in raw, unparsed blocks and returns a parsed block if
// possible.
// isOut toggles whether we should parse this block as an outgoing block (sent
// packet) or as an incoming block (recv'd packet)
func ParseBlock(block *xivnet.Block, isOut bool) (*xivnet.Block, error) {
	data, ok := block.Data.(*xivnet.GenericBlockData)
	if !ok {
		return block, nil
	}
	blockBytes, _ := data.MarshalBytes()

	var bd xivnet.BlockData

	// This is kind of a hack to check if a packet is a chat type since it's
	// impossible to otherwise obtain this information without capturing the open
	// of the game connections.
	if block.ServerID == 0 && block.SubjectID != 0 {
		bd = NewChatBlockData(block.Opcode, isOut)
	} else {
		bd = NewBlockData(block.Opcode, isOut)
	}

	if bd == nil {
		return block, nil
	}
	err := UnmarshalBlockBytes(blockBytes, bd)
	if err != nil {
		return block, err
	}
	newB := *block
	newB.Data = bd
	return &newB, nil
}
