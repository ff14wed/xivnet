package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x214 // Updated for 6.31h
	InitZoneOpcode      = 0x118 // Updated for 6.31h
	ControlOpcode       = 0x1A4 // Updated for 6.31h
	ControlSelfOpcode   = 0x203 // Updated for 6.31h
	ControlTargetOpcode = 0x7E  // Updated for 6.31h
	RemoveEntityOpcode  = 0x282 // Updated for 6.31h
	UpdateHPMPTPOpcode  = 0x119 // Updated for 6.31h

	ChatZoneOpcode = 0xC5 // Updated for 6.31h

	UpdateStatusesOpcode       = 0x305 // Updated for 6.31h
	UpdateStatusesEurekaOpcode = 0x3A6 // Updated for 6.31h
	UpdateStatusesBossOpcode   = 0x1E4 // Updated for 6.31h

	ActionOpcode      = 0x100 // Updated for 6.31h
	AoEAction8Opcode  = 0x2B9 // Updated for 6.31h
	AoEAction16Opcode = 0x390 // Updated for 6.31h
	AoEAction24Opcode = 0x22A // Updated for 6.31h
	AoEAction32Opcode = 0x120 // Updated for 6.31h

	ObjectSpawnOpcode = 0x277 // Updated for 6.31h
	PlayerSpawnOpcode = 0xF9  // Updated for 6.31h
	NPCSpawnOpcode    = 0x3D5 // Updated for 6.31h
	NPCSpawn2Opcode   = 0x3B6 // Updated for 6.31h

	MovementOpcode = 0x155 // Updated for 6.31h
	SetPosOpcode   = 0x99  // Updated for 6.31h

	CastingOpcode = 0x185 // Updated for 6.31h

	HateRankingOpcode = 0x1DD // Updated for 6.31h
	HateListOpcode    = 0x3A5 // Updated for 6.31h

	PlayerStatsOpcode = 0x2B6 // Updated for 6.31h

	EquipChangeOpcode = 0xE1 // Updated for 6.31h

	EventPlayOpcode    = 0x3B8 // Updated for 6.31h
	EventPlay4Opcode   = 0x1EC // Updated for 6.31h
	EventPlay8Opcode   = 0x333 // Updated for 6.31h
	EventPlay16Opcode  = 0x3AE // Updated for 6.31h
	EventPlay32Opcode  = 0x160 // Updated for 6.31h
	EventPlay64Opcode  = 0x2F2 // Updated for 6.31h
	EventPlay128Opcode = 0x8B  // Updated for 6.31h
	EventPlay255Opcode = 0x10B // Updated for 6.31h

	MountOpcode = 0x116 // Updated for 6.31h

	WeatherChangeOpcode = 0x148 // Updated for 6.31h

	PrepareZoningOpcode = 0x27C // Updated for 6.31h

	GaugeOpcode = 0x238 // Updated for 6.31h

	WaymarkOpcode         = 0x3CD // Updated for 6.31h
	PerformOpcode         = 0x1E1 // Updated for 6.31h
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
	registerInBlockFactory(EventPlay8Opcode, func() xivnet.BlockData { return new(EventPlay8) })
	registerInBlockFactory(EventPlay16Opcode, func() xivnet.BlockData { return new(EventPlay16) })
	registerInBlockFactory(EventPlay32Opcode, func() xivnet.BlockData { return new(EventPlay32) })
	registerInBlockFactory(EventPlay64Opcode, func() xivnet.BlockData { return new(EventPlay64) })
	registerInBlockFactory(EventPlay128Opcode, func() xivnet.BlockData { return new(EventPlay128) })
	registerInBlockFactory(EventPlay255Opcode, func() xivnet.BlockData { return new(EventPlay255) })

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
