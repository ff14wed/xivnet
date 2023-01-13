package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x1C8 // Updated for 6.3
	InitZoneOpcode      = 0x28A // Updated for 6.3
	ControlOpcode       = 0x131 // Updated for 6.3
	ControlSelfOpcode   = 0x1A8 // Updated for 6.3
	ControlTargetOpcode = 0x8F  // Updated for 6.3
	RemoveEntityOpcode  = 0x208 // Updated for 6.3
	UpdateHPMPTPOpcode  = 0x15B // Updated for 6.3

	ChatZoneOpcode = 0x9B // Updated for 6.3

	UpdateStatusesOpcode       = 0x192 // Updated for 6.3
	UpdateStatusesEurekaOpcode = 0x21F // Updated for 6.3
	UpdateStatusesBossOpcode   = 0x3DE // Updated for 6.3

	ActionOpcode      = 0x1E6 // Updated for 6.3
	AoEAction8Opcode  = 0x31B // Updated for 6.3
	AoEAction16Opcode = 0x3BE // Updated for 6.3
	AoEAction24Opcode = 0x2C4 // Updated for 6.3
	AoEAction32Opcode = 0x214 // Updated for 6.3

	ObjectSpawnOpcode = 0xFF  // Updated for 6.3
	PlayerSpawnOpcode = 0x321 // Updated for 6.3
	NPCSpawnOpcode    = 0x166 // Updated for 6.3
	NPCSpawn2Opcode   = 0xC8  // Updated for 6.3

	MovementOpcode = 0x183 // Updated for 6.3
	SetPosOpcode   = 0x181 // Updated for 6.3

	CastingOpcode = 0xB0 // Updated for 6.3

	HateRankingOpcode = 0xB8 // Updated for 6.3
	HateListOpcode    = 0xC0 // Updated for 6.3

	PlayerStatsOpcode = 0xE6 // Updated for 6.3

	EquipChangeOpcode = 0x169 // Updated for 6.3

	EventPlayOpcode    = 0x36B // Updated for 6.3
	EventPlay4Opcode   = 0x3D6 // Updated for 6.3
	EventPlay8Opcode   = 0x161 // Updated for 6.3
	EventPlay16Opcode  = 0x158 // Updated for 6.3
	EventPlay32Opcode  = 0x3CB // Updated for 6.3
	EventPlay64Opcode  = 0x223 // Updated for 6.3
	EventPlay128Opcode = 0x26E // Updated for 6.3
	EventPlay255Opcode = 0x135 // Updated for 6.3

	MountOpcode = 0x219 // Updated for 6.3

	WeatherChangeOpcode = 0x344 // Updated for 6.3

	PrepareZoningOpcode = 0x19D // Updated for 6.3

	GaugeOpcode = 0xF9 // Updated for 6.3

	WaymarkOpcode         = 0x235 // Updated for 6.3
	PerformOpcode         = 0x3CA // Updated for 6.3
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
