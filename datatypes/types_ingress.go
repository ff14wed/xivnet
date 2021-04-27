package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x334 // Updated for 5.5a
	InitZoneOpcode      = 0x85  // Updated for 5.5a
	ControlOpcode       = 0x2DB // Updated for 5.5a
	ControlSelfOpcode   = 0x23E // Updated for 5.5a
	ControlTargetOpcode = 0x368 // Updated for 5.5a
	RemoveEntityOpcode  = 0x322 // Updated for 5.5a
	UpdateHPMPTPOpcode  = 0x301 // Updated for 5.5a

	ChatZoneOpcode = 0x38F // Updated for 5.5a

	UpdateStatusesOpcode       = 0x19A // Updated for 5.5a
	UpdateStatusesEurekaOpcode = 0x101 // Updated for 5.5a
	UpdateStatusesBossOpcode   = 0x250 // Updated for 5.5a

	ActionOpcode      = 0x3E7 // Updated for 5.5a
	AoEAction8Opcode  = 0x20A // Updated for 5.5a
	AoEAction16Opcode = 0x1F9 // Updated for 5.5a
	AoEAction24Opcode = 0x1EB // Updated for 5.5a
	AoEAction32Opcode = 0x13E // Updated for 5.5a

	ObjectSpawnOpcode = 0x191 // Updated for 5.5a
	PlayerSpawnOpcode = 0x36F // Updated for 5.5a
	NPCSpawnOpcode    = 0x203 // Updated for 5.5a
	NPCSpawn2Opcode   = 0x1B3 // Updated for 5.5a

	MovementOpcode = 0x26D // Updated for 5.5a
	SetPosOpcode   = 0x7B  // Updated for 5.5a

	CastingOpcode = 0x2AD // Updated for 5.5a

	HateRankingOpcode = 0x1C0 // Updated for 5.5a
	HateListOpcode    = 0x78  // Updated for 5.5a

	EquipChangeOpcode = 0x28F // Updated for 5.5a

	EventPlayOpcode    = 0x32B // Updated for 5.5a
	EventPlay4Opcode   = 0xEE  // Updated for 5.5a
	EventPlay8Opcode   = 0xE4  // Updated for 5.5a
	EventPlay16Opcode  = 0x1CD // Updated for 5.5a
	EventPlay32Opcode  = 0x3B0 // Updated for 5.5a
	EventPlay64Opcode  = 0x65  // Updated for 5.5a
	EventPlay128Opcode = 0x137 // Updated for 5.5a
	EventPlay255Opcode = 0x21A // Updated for 5.5a

	MountOpcode = 0x177 // Updated for 5.5a

	WeatherChangeOpcode = 0x2C0 // Updated for 5.5a

	PrepareZoningOpcode = 0x251 // Updated for 5.5a

	GaugeOpcode = 0x18D // Updated for 5.5a

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
