package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x34B // Updated for 6.28
	InitZoneOpcode      = 0x3E5 // Updated for 6.28
	ControlOpcode       = 0x2D3 // Updated for 6.28
	ControlSelfOpcode   = 0x29B // Updated for 6.28
	ControlTargetOpcode = 0x253 // Updated for 6.28
	RemoveEntityOpcode  = 0xD4  // Updated for 6.28
	UpdateHPMPTPOpcode  = 0x2DC // Updated for 6.28

	ChatZoneOpcode = 0x1A8 // Updated for 6.28

	UpdateStatusesOpcode       = 0x90  // Updated for 6.28
	UpdateStatusesEurekaOpcode = 0xC5  // Updated for 6.28
	UpdateStatusesBossOpcode   = 0x36D // Updated for 6.28

	ActionOpcode      = 0x130 // Updated for 6.28
	AoEAction8Opcode  = 0x1DF // Updated for 6.28
	AoEAction16Opcode = 0x202 // Updated for 6.28
	AoEAction24Opcode = 0xBD  // Updated for 6.28
	AoEAction32Opcode = 0x14C // Updated for 6.28

	ObjectSpawnOpcode = 0xBE  // Updated for 6.28
	PlayerSpawnOpcode = 0x2C2 // Updated for 6.28
	NPCSpawnOpcode    = 0x3CD // Updated for 6.28
	NPCSpawn2Opcode   = 0x7C  // Updated for 6.28

	MovementOpcode = 0xC2  // Updated for 6.28
	SetPosOpcode   = 0x22E // Updated for 6.28

	CastingOpcode = 0x398 // Updated for 6.28

	HateRankingOpcode = 0x1AE // Updated for 6.28
	HateListOpcode    = 0x26A // Updated for 6.28

	PlayerStatsOpcode = 0x313 // Updated for 6.28

	EquipChangeOpcode = 0x360 // Updated for 6.28

	EventPlayOpcode    = 0xc9  // Updated for 6.28
	EventPlay4Opcode   = 0x321 // Updated for 6.28
	EventPlay8Opcode   = 0x9b  // Updated for 6.28
	EventPlay16Opcode  = 0x1c2 // Updated for 6.28
	EventPlay32Opcode  = 0x28d // Updated for 6.28
	EventPlay64Opcode  = 0x6b  // Updated for 6.28
	EventPlay128Opcode = 0x200 // Updated for 6.28
	EventPlay255Opcode = 0x314 // Updated for 6.28

	MountOpcode = 0x2A8 // Updated for 6.28

	WeatherChangeOpcode = 0xF9 // Updated for 6.28

	PrepareZoningOpcode = 0x281 // Updated for 6.28

	GaugeOpcode = 0x11E // Updated for 6.28

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
