package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x21A // Updated for 6.08 hotfix
	InitZoneOpcode      = 0x17F // Updated for 6.08 hotfix
	ControlOpcode       = 0x202 // Updated for 6.08 hotfix
	ControlSelfOpcode   = 0x301 // Updated for 6.08 hotfix
	ControlTargetOpcode = 0x333 // Updated for 6.08 hotfix
	RemoveEntityOpcode  = 0x1D3 // Updated for 6.08 hotfix
	UpdateHPMPTPOpcode  = 0x240 // Updated for 6.08 hotfix

	ChatZoneOpcode = 0x2C6 // Updated for 6.08 hotfix

	UpdateStatusesOpcode       = 0x275 // Updated for 6.08 hotfix
	UpdateStatusesEurekaOpcode = 0x170 // Updated for 6.08 hotfix
	UpdateStatusesBossOpcode   = 0x114 // Updated for 6.08 hotfix

	ActionOpcode      = 0x35E // Updated for 6.08 hotfix
	AoEAction8Opcode  = 0x2BA // Updated for 6.08 hotfix
	AoEAction16Opcode = 0x10D // Updated for 6.08 hotfix
	AoEAction24Opcode = 0xF7  // Updated for 6.08 hotfix
	AoEAction32Opcode = 0x1CA // Updated for 6.08 hotfix

	ObjectSpawnOpcode = 0x6A  // Updated for 6.08 hotfix
	PlayerSpawnOpcode = 0x3DC // Updated for 6.08 hotfix
	NPCSpawnOpcode    = 0x32F // Updated for 6.08 hotfix
	NPCSpawn2Opcode   = 0x380 // Updated for 6.08 hotfix

	MovementOpcode = 0x3CB // Updated for 6.08 hotfix
	SetPosOpcode   = 0x33C // Updated for 6.08 hotfix

	CastingOpcode = 0x1F4 // Updated for 6.08 hotfix

	HateRankingOpcode = 0x3B0 // Updated for 6.08 hotfix
	HateListOpcode    = 0x376 // Updated for 6.08 hotfix

	EquipChangeOpcode = 0xCF // Updated for 6.08 hotfix

	EventPlayOpcode    = 0x27A // Updated for 6.08 hotfix
	EventPlay4Opcode   = 0x12E // Updated for 6.08 hotfix
	EventPlay8Opcode   = 0x2E6 // Updated for 6.08 hotfix
	EventPlay16Opcode  = 0x21E // Updated for 6.08 hotfix
	EventPlay32Opcode  = 0x1F9 // Updated for 6.08 hotfix
	EventPlay64Opcode  = 0x360 // Updated for 6.08 hotfix
	EventPlay128Opcode = 0x33B // Updated for 6.08 hotfix
	EventPlay255Opcode = 0x34D // Updated for 6.08 hotfix

	MountOpcode = 0x321 // Updated for 6.08 hotfix

	WeatherChangeOpcode = 0x337 // Updated for 6.08 hotfix

	PrepareZoningOpcode = 0x1CB // Updated for 6.08 hotfix

	GaugeOpcode = 0x374 // Updated for 6.08 hotfix

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
