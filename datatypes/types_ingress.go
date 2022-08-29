package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x263 // Updated for 6.2 Hotfix
	InitZoneOpcode      = 0xE1  // Updated for 6.2 Hotfix
	ControlOpcode       = 0x2A7 // Updated for 6.2 Hotfix
	ControlSelfOpcode   = 0x23C // Updated for 6.2 Hotfix
	ControlTargetOpcode = 0x118 // Updated for 6.2 Hotfix
	RemoveEntityOpcode  = 0xAA  // Updated for 6.2 Hotfix
	UpdateHPMPTPOpcode  = 0x102 // Updated for 6.2 Hotfix

	ChatZoneOpcode = 0xB7 // Updated for 6.2 Hotfix

	UpdateStatusesOpcode       = 0x265 // Updated for 6.2 Hotfix
	UpdateStatusesEurekaOpcode = 0x363 // Updated for 6.2 Hotfix
	UpdateStatusesBossOpcode   = 0x264 // Updated for 6.2 Hotfix

	ActionOpcode      = 0x94  // Updated for 6.2 Hotfix
	AoEAction8Opcode  = 0x2BB // Updated for 6.2 Hotfix
	AoEAction16Opcode = 0x267 // Updated for 6.2 Hotfix
	AoEAction24Opcode = 0x373 // Updated for 6.2 Hotfix
	AoEAction32Opcode = 0x3AC // Updated for 6.2 Hotfix

	ObjectSpawnOpcode = 0x2F7 // Updated for 6.2 Hotfix
	PlayerSpawnOpcode = 0x334 // Updated for 6.2 Hotfix
	NPCSpawnOpcode    = 0x19B // Updated for 6.2 Hotfix
	NPCSpawn2Opcode   = 0x31A // Updated for 6.2 Hotfix

	MovementOpcode = 0xB3  // Updated for 6.2 Hotfix
	SetPosOpcode   = 0x1BA // Updated for 6.2 Hotfix

	CastingOpcode = 0x26C // Updated for 6.2 Hotfix

	HateRankingOpcode = 0xF3  // Updated for 6.2 Hotfix
	HateListOpcode    = 0x1B7 // Updated for 6.2 Hotfix

	PlayerStatsOpcode = 0x26B // Updated for 6.2 Hotfix

	EquipChangeOpcode = 0xA1 // Updated for 6.2 Hotfix

	EventPlayOpcode    = 0x2FD // Updated for 6.2 Hotfix
	EventPlay4Opcode   = 0x380 // Updated for 6.2 Hotfix
	EventPlay8Opcode   = 0x107 // Updated for 6.2 Hotfix
	EventPlay16Opcode  = 0x2A4 // Updated for 6.2 Hotfix
	EventPlay32Opcode  = 0xC1  // Updated for 6.2 Hotfix
	EventPlay64Opcode  = 0x2FB // Updated for 6.2 Hotfix
	EventPlay128Opcode = 0x129 // Updated for 6.2 Hotfix
	EventPlay255Opcode = 0x2CD // Updated for 6.2 Hotfix

	MountOpcode = 0x253 // Updated for 6.2 Hotfix

	WeatherChangeOpcode = 0x281 // Updated for 6.2 Hotfix

	PrepareZoningOpcode = 0xA0 // Updated for 6.2 Hotfix

	GaugeOpcode = 0x2AB // Updated for 6.2 Hotfix

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
