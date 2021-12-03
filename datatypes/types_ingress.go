package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x196 // Updated for 6.0
	InitZoneOpcode      = 0x2C4 // Updated for 6.0
	ControlOpcode       = 0x17E // Updated for 6.0
	ControlSelfOpcode   = 0x2E6 // Updated for 6.0
	ControlTargetOpcode = 0x168 // Updated for 6.0
	RemoveEntityOpcode  = 0x18A // Updated for 6.0
	UpdateHPMPTPOpcode  = 0x296 // Updated for 6.0

	ChatZoneOpcode = 0x1A6 // Updated for 6.0

	UpdateStatusesOpcode       = 0x2C5 // Updated for 6.0
	UpdateStatusesEurekaOpcode = 0x371 // Updated for 6.0
	UpdateStatusesBossOpcode   = 0x13A // Updated for 6.0

	ActionOpcode      = 0x35A // Updated for 6.0
	AoEAction8Opcode  = 0x1BA // Updated for 6.0
	AoEAction16Opcode = 0x2CE // Updated for 6.0
	AoEAction24Opcode = 0x2ED // Updated for 6.0
	AoEAction32Opcode = 0x23A // Updated for 6.0

	ObjectSpawnOpcode = 0x319 // Updated for 6.0
	PlayerSpawnOpcode = 0x133 // Updated for 6.0
	NPCSpawnOpcode    = 0x32E // Updated for 6.0
	NPCSpawn2Opcode   = 0x2F2 // Updated for 6.0

	MovementOpcode = 0x235 // Updated for 6.0
	SetPosOpcode   = 0x199 // Updated for 6.0

	CastingOpcode = 0x108 // Updated for 6.0

	HateRankingOpcode = 0x15A // Updated for 6.0
	HateListOpcode    = 0x1E2 // Updated for 6.0

	EquipChangeOpcode = 0xFA // Updated for 6.0

	EventPlayOpcode    = 0xA5  // Updated for 6.0
	EventPlay4Opcode   = 0x22E // Updated for 6.0
	EventPlay8Opcode   = 0x18B // Updated for 6.0
	EventPlay16Opcode  = 0x1F4 // Updated for 6.0
	EventPlay32Opcode  = 0x65  // Updated for 6.0
	EventPlay64Opcode  = 0x3A8 // Updated for 6.0
	EventPlay128Opcode = 0x16E // Updated for 6.0
	EventPlay255Opcode = 0x366 // Updated for 6.0

	MountOpcode = 0x28E // Updated for 6.0

	WeatherChangeOpcode = 0x1FD // Updated for 6.0

	PrepareZoningOpcode = 0x90 // Updated for 6.0

	GaugeOpcode = 0x283 // Updated for 6.0

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
