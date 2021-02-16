package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x1D7 // Updated for 5.45a
	InitZoneOpcode      = 0x233 // Updated for 5.45a
	ControlOpcode       = 0xC2  // Updated for 5.45a
	ControlSelfOpcode   = 0x3D5 // Updated for 5.45a
	ControlTargetOpcode = 0x171 // Updated for 5.45a
	RemoveEntityOpcode  = 0x185 // Updated for 5.45a
	UpdateHPMPTPOpcode  = 0x19B // Updated for 5.45a

	ChatZoneOpcode = 0x1BA // Updated for 5.45a

	UpdateStatusesOpcode       = 0x243 // Updated for 5.45a
	UpdateStatusesEurekaOpcode = 0x2C7 // Updated for 5.45a
	UpdateStatusesBossOpcode   = 0x90  // Updated for 5.45a

	ActionOpcode      = 0x27F // Updated for 5.45a
	AoEAction8Opcode  = 0x9B  // Updated for 5.45a
	AoEAction16Opcode = 0x28C // Updated for 5.45a
	AoEAction24Opcode = 0x2AD // Updated for 5.45a
	AoEAction32Opcode = 0xA7  // Updated for 5.45a

	ObjectSpawnOpcode = 0x336 // Updated for 5.45a
	PlayerSpawnOpcode = 0x1AB // Updated for 5.45a
	NPCSpawnOpcode    = 0x2C9 // Updated for 5.45a
	NPCSpawn2Opcode   = 0x240 // Updated for 5.45a

	MovementOpcode = 0x6B // Updated for 5.45a
	SetPosOpcode   = 0xC6 // Updated for 5.45a

	CastingOpcode = 0x34C // Updated for 5.45a

	HateRankingOpcode = 0x217 // Updated for 5.45a
	HateListOpcode    = 0xE5  // Updated for 5.45a

	EquipChangeOpcode = 0x24A // Updated for 5.45a

	EventPlayOpcode    = 0x276 // Updated for 5.45a
	EventPlay4Opcode   = 0x28D // Updated for 5.45a
	EventPlay8Opcode   = 0x318 // Updated for 5.45a
	EventPlay16Opcode  = 0x20D // Updated for 5.45a
	EventPlay32Opcode  = 0x3B5 // Updated for 5.45a
	EventPlay64Opcode  = 0x39E // Updated for 5.45a
	EventPlay128Opcode = 0x2F4 // Updated for 5.45a
	EventPlay255Opcode = 0x9D  // Updated for 5.45a

	MountOpcode = 0x122 // Updated for 5.45a

	WeatherChangeOpcode = 0x167 // Updated for 5.45a

	PrepareZoningOpcode = 0x1EE // Updated for 5.45a

	GaugeOpcode = 0x18E // Updated for 5.45a

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
