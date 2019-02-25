package datatypes

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/ff14wed/xivnet/v2"
)

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)
var outTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	// Opcodes that change rarely
	AddStatusOpcode    = 0x141 // Unchanged 4.5
	InitZoneOpcode     = 0x19A // Unchanged 4.5
	Notify142Opcode    = 0x142 // Unchanged 4.5
	Notify143Opcode    = 0x143 // Unchanged 4.5
	Notify144Opcode    = 0x144 // Unchanged 4.5
	RemoveEntityOpcode = 0x191 // Unchanged 4.5
	UpdateHPMPTPOpcode = 0x145 // Unchanged 4.5

	UpdateStatusesOpcode = 0x151 // Updated for 4.5
	ActionOpcode         = 0x154 // Updated for 4.5
	AoEAction8Opcode     = 0x157 // Updated for 4.5
	AoEAction16Opcode    = 0x158 // Updated for 4.5
	AoEAction24Opcode    = 0x159 // Updated for 4.5
	AoEAction32Opcode    = 0x15A // Updated for 4.5

	PlayerSpawnOpcode = 0x175 // Updated for 4.5
	NPCSpawnOpcode    = 0x176 // Updated for 4.5
	NPCSpawn2Opcode   = 0x177 // Updated for 4.5

	MovementOpcode = 0x178 // Updated for 4.5
	SetPosOpcode   = 0x17A // Updated for 4.5

	CastingOpcode = 0x17C // Updated for 4.5

	HateRankingOpcode = 0x17F // Updated for 4.5
	HateListOpcode    = 0x180 // Updated for 4.5

	EquipChangeOpcode = 0x18B // Updated for 4.5

	EventPlayOpcode   = 0x1AB // Updated for 4.5
	EventPlay2Opcode  = 0x1AC // Updated for 4.5
	DirectorPlayScene = 0x1AF // Updated for 4.5

	MountOpcode = 0x1E8 // Updated for 4.5

	WeatherChangeOpcode = 0x205 // Updated for 4.5

	WaymarkOpcode = 0x267 // Updated for 4.5

	PrepareZoningOpcode = 0x299 // Updated for 4.5

	GaugeOpcode   = 0x29A // Updated for 4.5
	PerformOpcode = 0x2A5 // Updated for 4.5

	XWorldPartyListOpcode = 0xA1 // Updated 4.18
)

func init() {
	registerInBlockFactory(AddStatusOpcode, func() xivnet.BlockData { return new(AddStatus) })
	registerInBlockFactory(InitZoneOpcode, func() xivnet.BlockData { return new(InitZone) })
	registerInBlockFactory(Notify142Opcode, func() xivnet.BlockData { return new(Notify) })
	registerInBlockFactory(Notify143Opcode, func() xivnet.BlockData { return new(Notify3) })
	registerInBlockFactory(Notify144Opcode, func() xivnet.BlockData { return new(Notify4) })
	registerInBlockFactory(RemoveEntityOpcode, func() xivnet.BlockData { return new(RemoveEntity) })
	registerInBlockFactory(UpdateHPMPTPOpcode, func() xivnet.BlockData { return new(UpdateHPMPTP) })

	registerInBlockFactory(UpdateStatusesOpcode, func() xivnet.BlockData { return new(UpdateStatuses) })
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

	registerInBlockFactory(DirectorPlayScene, func() xivnet.BlockData { return new(CraftState) })

	registerInBlockFactory(MountOpcode, func() xivnet.BlockData { return new(Mount) })

	registerInBlockFactory(WeatherChangeOpcode, func() xivnet.BlockData { return new(WeatherChange) })

	registerInBlockFactory(WaymarkOpcode, func() xivnet.BlockData { return new(Marker) })

	registerInBlockFactory(PrepareZoningOpcode, func() xivnet.BlockData { return new(MapChange) })

	registerInBlockFactory(GaugeOpcode, func() xivnet.BlockData { return new(Gauge) })
	registerInBlockFactory(PerformOpcode, func() xivnet.BlockData { return new(Perform) })

	// registerInBlockFactory(XWorldPartyListOpcode, new(XWorldPartyList))
}

// Opcodes that define the datatypes of outgoing (to server) network blocks
const (
	EgressClientTriggerOpcode = 0x138 // Updated for 4.5

	EgressMovementOpcode         = 0x13F // Updated for 4.5
	EgressInstanceMovementOpcode = 0x17E // Updated for 4.5

	EgressPerformOpcode        = 0x188 // Updated for 4.5
	EgressIngredientListOpcode = 0x15D // Updated for 4.5
)

func init() {
	registerOutBlockFactory(EgressClientTriggerOpcode, func() xivnet.BlockData { return new(EgressClientTrigger) })

	registerOutBlockFactory(EgressMovementOpcode, func() xivnet.BlockData { return new(EgressMovement) })
	registerOutBlockFactory(EgressInstanceMovementOpcode, func() xivnet.BlockData { return new(EgressInstanceMovement) })

	registerOutBlockFactory(EgressPerformOpcode, func() xivnet.BlockData { return new(Perform) })

	registerOutBlockFactory(EgressIngredientListOpcode, func() xivnet.BlockData { return new(EgressIngredientList) })
}

// NewBlockData is a factory for BlockData that uses the opcode to
// determine which BlockData to create
func NewBlockData(opcode uint16, isOut bool) xivnet.BlockData {
	r := inTypeRegistry
	if isOut {
		r = outTypeRegistry
	}
	factory, ok := r[opcode]
	if !ok {
		return nil
	}
	return factory()
}

// UnmarshalBlockBytes converts raw bytes to this block data struct
func UnmarshalBlockBytes(data []byte, block xivnet.BlockData) error {
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
	bd := NewBlockData(block.Opcode, isOut)
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
