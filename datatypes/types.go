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
	EffectResultOpcode  = 0x30F // Updated for 5.25
	InitZoneOpcode      = 0x1CD // Updated for 5.25
	ControlOpcode       = 0x1DD // Updated for 5.25
	ControlSelfOpcode   = 0x165 // Updated for 5.25
	ControlTargetOpcode = 0x1EB // Updated for 5.25
	RemoveEntityOpcode  = 0x194 // Updated for 5.25
	UpdateHPMPTPOpcode  = 0xF8  // Updated for 5.25

	ChatZoneOpcode = 0x161 // Updated for 5.25

	UpdateStatusesOpcode       = 0x1D7 // Updated for 5.25
	UpdateStatusesEurekaOpcode = 0x25A // Updated for 5.25
	UpdateStatusesBossOpcode   = 0x1F5 // Updated for 5.25

	ActionOpcode      = 0x87  // Updated for 5.25
	AoEAction8Opcode  = 0x1A6 // Updated for 5.25
	AoEAction16Opcode = 0xBB  // Updated for 5.25
	AoEAction24Opcode = 0x3B6 // Updated for 5.25
	AoEAction32Opcode = 0x3C2 // Updated for 5.25

	PlayerSpawnOpcode = 0xB9  // Updated for 5.25
	NPCSpawnOpcode    = 0x38D // Updated for 5.25
	NPCSpawn2Opcode   = 0x2D3 // Updated for 5.25

	MovementOpcode = 0x290 // Updated for 5.25
	SetPosOpcode   = 0x28E // Updated for 5.25

	CastingOpcode = 0xC4 // Updated for 5.25

	HateRankingOpcode = 0x12E // Updated for 5.25
	HateListOpcode    = 0x2C4 // Updated for 5.25

	EquipChangeOpcode = 0x1E5 // Updated for 5.25

	EventPlayOpcode   = 0x2C3 // Updated for 5.25
	EventPlay4Opcode  = 0x29A // Updated for 5.25
	EventPlay32Opcode = 0x2D9 // Updated for 5.25

	MountOpcode = 0x2BE // Updated for 5.25

	WeatherChangeOpcode = 0x157 // Updated for 5.25

	// WaymarkOpcode = 0x272 // Updated for 5.0
	WaymarkOpcode = UndefinedOpcode

	PrepareZoningOpcode = 0x3B3 // Updated for 5.25

	GaugeOpcode = 0x7D // Updated for 5.25
	// PerformOpcode = 0x2A5 // Updated for 4.5
	PerformOpcode = UndefinedOpcode

	// XWorldPartyListOpcode = 0xA1 // Updated 4.18
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
	EgressClientTriggerOpcode = 0x17D // Updated for 5.25

	EgressChatZoneOpcode = 0x189 // Updated for 5.25

	EgressMovementOpcode         = 0x14D // Updated for 5.25
	EgressInstanceMovementOpcode = 0x290 // Updated for 5.25

	// EgressPerformOpcode    = 0x18B // Updated for 5.0
	EgressPerformOpcode = UndefinedOpcode

	EgressCraftEventOpcode = 0x32E // Updated for 5.25
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
	FreeCompanyResultOpcode = 0x12C
	ChatXWorldOpcode        = 0x72
)

func init() {
	registerInChatBlockFactory(ChatFromOpcode, func() xivnet.BlockData { return new(ChatFrom) })
	registerInChatBlockFactory(ChatOpcode, func() xivnet.BlockData { return new(Chat) })
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
