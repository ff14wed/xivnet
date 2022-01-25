package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0xDE  // Updated for 6.08
	InitZoneOpcode      = 0x1EB // Updated for 6.08
	ControlOpcode       = 0x22F // Updated for 6.08
	ControlSelfOpcode   = 0x6B  // Updated for 6.08
	ControlTargetOpcode = 0x191 // Updated for 6.08
	RemoveEntityOpcode  = 0x82  // Updated for 6.08
	UpdateHPMPTPOpcode  = 0x2C9 // Updated for 6.08

	ChatZoneOpcode = 0x148 // Updated for 6.08

	UpdateStatusesOpcode       = 0xBC  // Updated for 6.08
	UpdateStatusesEurekaOpcode = 0x2D8 // Updated for 6.08
	UpdateStatusesBossOpcode   = 0x7E  // Updated for 6.08

	ActionOpcode      = 0x3C7 // Updated for 6.08
	AoEAction8Opcode  = 0x149 // Updated for 6.08
	AoEAction16Opcode = 0xC1  // Updated for 6.08
	AoEAction24Opcode = 0x213 // Updated for 6.08
	AoEAction32Opcode = 0x38B // Updated for 6.08

	ObjectSpawnOpcode = 0x3A3 // Updated for 6.08
	PlayerSpawnOpcode = 0x226 // Updated for 6.08
	NPCSpawnOpcode    = 0x32C // Updated for 6.08
	NPCSpawn2Opcode   = 0x8F  // Updated for 6.08

	MovementOpcode = 0x370 // Updated for 6.08
	SetPosOpcode   = 0x395 // Updated for 6.08

	CastingOpcode = 0x104 // Updated for 6.08

	HateRankingOpcode = 0x3C3 // Updated for 6.08
	HateListOpcode    = 0x32B // Updated for 6.08

	EquipChangeOpcode = 0xE4 // Updated for 6.08

	EventPlayOpcode    = 0x113 // Updated for 6.08
	EventPlay4Opcode   = 0x302 // Updated for 6.08
	EventPlay8Opcode   = 0x78  // Updated for 6.08
	EventPlay16Opcode  = 0x223 // Updated for 6.08
	EventPlay32Opcode  = 0x2F2 // Updated for 6.08
	EventPlay64Opcode  = 0x3BC // Updated for 6.08
	EventPlay128Opcode = 0x33E // Updated for 6.08
	EventPlay255Opcode = 0x79  // Updated for 6.08

	MountOpcode = 0x373 // Updated for 6.08

	WeatherChangeOpcode = 0xED // Updated for 6.08

	PrepareZoningOpcode = 0x39A // Updated for 6.08

	GaugeOpcode = 0x3B5 // Updated for 6.08

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

func registerInBlockFactory(opcode uint16, factory func() xivnet.BlockData) {
	inTypeRegistry[opcode] = factory
}
