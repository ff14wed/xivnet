package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x12A // Updated for 6.11a
	InitZoneOpcode      = 0x86  // Updated for 6.11a
	ControlOpcode       = 0x24B // Updated for 6.11a
	ControlSelfOpcode   = 0x334 // Updated for 6.11a
	ControlTargetOpcode = 0x370 // Updated for 6.11a
	RemoveEntityOpcode  = 0x111 // Updated for 6.11a
	UpdateHPMPTPOpcode  = 0x231 // Updated for 6.11a

	ChatZoneOpcode = 0x9C // Updated for 6.11a

	UpdateStatusesOpcode       = 0x32E // Updated for 6.11a
	UpdateStatusesEurekaOpcode = 0x244 // Updated for 6.11a
	UpdateStatusesBossOpcode   = 0x33A // Updated for 6.11a

	ActionOpcode      = 0xB5  // Updated for 6.11a
	AoEAction8Opcode  = 0x14F // Updated for 6.11a
	AoEAction16Opcode = 0x1B4 // Updated for 6.11a
	AoEAction24Opcode = 0x3A0 // Updated for 6.11a
	AoEAction32Opcode = 0x168 // Updated for 6.11a

	ObjectSpawnOpcode = 0x305 // Updated for 6.11a
	PlayerSpawnOpcode = 0x336 // Updated for 6.11a
	NPCSpawnOpcode    = 0x26D // Updated for 6.11a
	NPCSpawn2Opcode   = 0x207 // Updated for 6.11a

	MovementOpcode = 0x132 // Updated for 6.11a
	SetPosOpcode   = 0x1D9 // Updated for 6.11a

	CastingOpcode = 0x3DF // Updated for 6.11a

	HateRankingOpcode = 0x19D // Updated for 6.11a
	HateListOpcode    = 0x34C // Updated for 6.11a

	EquipChangeOpcode = 0x194 // Updated for 6.11a

	EventPlayOpcode    = 0x85  // Updated for 6.11a
	EventPlay4Opcode   = 0x2F4 // Updated for 6.11a
	EventPlay8Opcode   = 0x176 // Updated for 6.11a
	EventPlay16Opcode  = 0x2D4 // Updated for 6.11a
	EventPlay32Opcode  = 0x2FF // Updated for 6.11a
	EventPlay64Opcode  = 0x289 // Updated for 6.11a
	EventPlay128Opcode = 0x3A5 // Updated for 6.11a
	EventPlay255Opcode = 0xD9  // Updated for 6.11a

	MountOpcode = 0x18C // Updated for 6.11a

	WeatherChangeOpcode = 0x1AD // Updated for 6.11a

	PrepareZoningOpcode = 0x2A0 // Updated for 6.11a

	GaugeOpcode = 0x1C2 // Updated for 6.11a

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
