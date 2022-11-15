package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x1F0 // Updated for 6.28a
	InitZoneOpcode      = 0xF0  // Updated for 6.28a
	ControlOpcode       = 0x197 // Updated for 6.28a
	ControlSelfOpcode   = 0x238 // Updated for 6.28a
	ControlTargetOpcode = 0x28B // Updated for 6.28a
	RemoveEntityOpcode  = 0x212 // Updated for 6.28a
	UpdateHPMPTPOpcode  = 0x2A2 // Updated for 6.28a

	ChatZoneOpcode = 0x39B // Updated for 6.28a

	UpdateStatusesOpcode       = 0x19B // Updated for 6.28a
	UpdateStatusesEurekaOpcode = 0x18C // Updated for 6.28a
	UpdateStatusesBossOpcode   = 0x34D // Updated for 6.28a

	ActionOpcode      = 0x395 // Updated for 6.28a
	AoEAction8Opcode  = 0x311 // Updated for 6.28a
	AoEAction16Opcode = 0x351 // Updated for 6.28a
	AoEAction24Opcode = 0x3C2 // Updated for 6.28a
	AoEAction32Opcode = 0x6A  // Updated for 6.28a

	ObjectSpawnOpcode = 0x265 // Updated for 6.28a
	PlayerSpawnOpcode = 0xDD  // Updated for 6.28a
	NPCSpawnOpcode    = 0x359 // Updated for 6.28a
	NPCSpawn2Opcode   = 0x190 // Updated for 6.28a

	MovementOpcode = 0x384 // Updated for 6.28a
	SetPosOpcode   = 0x329 // Updated for 6.28a

	CastingOpcode = 0x2AD // Updated for 6.28a

	HateRankingOpcode = 0x328 // Updated for 6.28a
	HateListOpcode    = 0x26C // Updated for 6.28a

	PlayerStatsOpcode = 0x86 // Updated for 6.28a

	EquipChangeOpcode = 0x170 // Updated for 6.28a

	EventPlayOpcode    = 0x38C // Updated for 6.28a
	EventPlay4Opcode   = 0x274 // Updated for 6.28a
	EventPlay8Opcode   = 0x2D9 // Updated for 6.28a
	EventPlay16Opcode  = 0x1AB // Updated for 6.28a
	EventPlay32Opcode  = 0x203 // Updated for 6.28a
	EventPlay64Opcode  = 0x38D // Updated for 6.28a
	EventPlay128Opcode = 0x35F // Updated for 6.28a
	EventPlay255Opcode = 0x11D // Updated for 6.28a

	MountOpcode = 0xA2 // Updated for 6.28a

	WeatherChangeOpcode = 0x31B // Updated for 6.28a

	PrepareZoningOpcode = 0x2DB // Updated for 6.28a

	GaugeOpcode = 0x382 // Updated for 6.28a

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

	registerInBlockFactory(PlayerStatsOpcode, func() xivnet.BlockData { return new(PlayerStats) })

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
