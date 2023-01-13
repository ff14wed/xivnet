package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x22C // Updated for 6.30h
	InitZoneOpcode      = 0x222 // Updated for 6.30h
	ControlOpcode       = 0x179 // Updated for 6.30h
	ControlSelfOpcode   = 0x26F // Updated for 6.30h
	ControlTargetOpcode = 0x220 // Updated for 6.30h
	RemoveEntityOpcode  = 0x28A // Updated for 6.30h
	UpdateHPMPTPOpcode  = 0x383 // Updated for 6.30h

	ChatZoneOpcode = 0x10A // Updated for 6.30h

	UpdateStatusesOpcode       = 0x2BC // Updated for 6.30h
	UpdateStatusesEurekaOpcode = 0x353 // Updated for 6.30h
	UpdateStatusesBossOpcode   = 0x1EE // Updated for 6.30h

	ActionOpcode      = 0x1C9 // Updated for 6.30h
	AoEAction8Opcode  = 0x24A // Updated for 6.30h
	AoEAction16Opcode = 0x38A // Updated for 6.30h
	AoEAction24Opcode = 0xC8  // Updated for 6.30h
	AoEAction32Opcode = 0x32B // Updated for 6.30h

	// ObjectSpawnOpcode = 0x31B // Updated for 6.30h
	PlayerSpawnOpcode = 0x7F  // Updated for 6.30h
	NPCSpawnOpcode    = 0x39E // Updated for 6.30h
	NPCSpawn2Opcode   = 0x2E5 // Updated for 6.30h

	MovementOpcode = 0x1DB // Updated for 6.30h
	SetPosOpcode   = 0x18C // Updated for 6.30h

	CastingOpcode = 0x29C // Updated for 6.30h

	HateRankingOpcode = 0x134 // Updated for 6.30h
	HateListOpcode    = 0x2F9 // Updated for 6.30h

	PlayerStatsOpcode = 0x1B8 // Updated for 6.30h

	EquipChangeOpcode = 0x286 // Updated for 6.30h

	EventPlayOpcode    = 0x2DE // Updated for 6.30h
	EventPlay4Opcode   = 0x317 // Updated for 6.30h
	EventPlay8Opcode   = 0x1CD // Updated for 6.30h
	EventPlay16Opcode  = 0x1FE // Updated for 6.30h
	EventPlay32Opcode  = 0x2FC // Updated for 6.30h
	EventPlay64Opcode  = 0x7C  // Updated for 6.30h
	EventPlay128Opcode = 0x337 // Updated for 6.30h
	EventPlay255Opcode = 0x1D2 // Updated for 6.30h

	MountOpcode = 0x322 // Updated for 6.30h

	WeatherChangeOpcode = 0xC7 // Updated for 6.30h

	PrepareZoningOpcode = 0x195 // Updated for 6.30h

	GaugeOpcode = 0x171 // Updated for 6.30h

	WaymarkOpcode         = 0x38E // Updated for 6.30h
	PerformOpcode         = 0xE9  // Updated for 6.30h
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
