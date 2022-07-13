package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x21E // Updated for 6.18 hotfix
	InitZoneOpcode      = 0x10F // Updated for 6.18 hotfix
	ControlOpcode       = 0xCB  // Updated for 6.18 hotfix
	ControlSelfOpcode   = 0x3CD // Updated for 6.18 hotfix
	ControlTargetOpcode = 0x174 // Updated for 6.18 hotfix
	RemoveEntityOpcode  = 0x309 // Updated for 6.18 hotfix
	UpdateHPMPTPOpcode  = 0xE7  // Updated for 6.18 hotfix

	ChatZoneOpcode = 0x2FC // Updated for 6.18 hotfix

	UpdateStatusesOpcode       = 0x1D0 // Updated for 6.18 hotfix
	UpdateStatusesEurekaOpcode = 0x3D9 // Updated for 6.18 hotfix
	UpdateStatusesBossOpcode   = 0x2A4 // Updated for 6.18 hotfix

	ActionOpcode      = 0x353 // Updated for 6.18 hotfix
	AoEAction8Opcode  = 0x2E0 // Updated for 6.18 hotfix
	AoEAction16Opcode = 0x14D // Updated for 6.18 hotfix
	AoEAction24Opcode = 0x65  // Updated for 6.18 hotfix
	AoEAction32Opcode = 0x12D // Updated for 6.18 hotfix

	ObjectSpawnOpcode = 0x313 // Updated for 6.18 hotfix
	PlayerSpawnOpcode = 0x20C // Updated for 6.18 hotfix
	NPCSpawnOpcode    = 0x1A4 // Updated for 6.18 hotfix
	NPCSpawn2Opcode   = 0x3A2 // Updated for 6.18 hotfix

	MovementOpcode = 0x227 // Updated for 6.18 hotfix
	SetPosOpcode   = 0x240 // Updated for 6.18 hotfix

	CastingOpcode = 0x77 // Updated for 6.18 hotfix

	HateRankingOpcode = 0x23F // Updated for 6.18 hotfix
	HateListOpcode    = 0x32A // Updated for 6.18 hotfix

	PlayerStatsOpcode = 0x310 // Updated for 6.18 hotfix

	EquipChangeOpcode = 0xD8 // Updated for 6.18 hotfix

	EventPlayOpcode    = 0xB6  // Updated for 6.18 hotfix
	EventPlay4Opcode   = 0xEF  // Updated for 6.18 hotfix
	EventPlay8Opcode   = 0x2BE // Updated for 6.18 hotfix
	EventPlay16Opcode  = 0x268 // Updated for 6.18 hotfix
	EventPlay32Opcode  = 0x1EC // Updated for 6.18 hotfix
	EventPlay64Opcode  = 0x2B4 // Updated for 6.18 hotfix
	EventPlay128Opcode = 0x1DA // Updated for 6.18 hotfix
	EventPlay255Opcode = 0x35C // Updated for 6.18 hotfix

	MountOpcode = 0x1CB // Updated for 6.18 hotfix

	WeatherChangeOpcode = 0x3BF // Updated for 6.18 hotfix

	PrepareZoningOpcode = 0x2F7 // Updated for 6.18 hotfix

	GaugeOpcode = 0x3A3 // Updated for 6.18 hotfix

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
