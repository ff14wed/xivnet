package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x11C // Updated for 6.01
	InitZoneOpcode      = 0x309 // Updated for 6.01
	ControlOpcode       = 0x1FB // Updated for 6.01
	ControlSelfOpcode   = 0x2C2 // Updated for 6.01
	ControlTargetOpcode = 0x143 // Updated for 6.01
	RemoveEntityOpcode  = 0x1C6 // Updated for 6.01
	UpdateHPMPTPOpcode  = 0x2B1 // Updated for 6.01

	ChatZoneOpcode = 0x341 // Updated for 6.01

	UpdateStatusesOpcode       = 0x2D2 // Updated for 6.01
	UpdateStatusesEurekaOpcode = 0x34B // Updated for 6.01
	UpdateStatusesBossOpcode   = 0x2F0 // Updated for 6.01

	ActionOpcode      = 0x1D0 // Updated for 6.01
	AoEAction8Opcode  = 0x2B0 // Updated for 6.01
	AoEAction16Opcode = 0x2BE // Updated for 6.01
	AoEAction24Opcode = 0x288 // Updated for 6.01
	AoEAction32Opcode = 0x238 // Updated for 6.01

	ObjectSpawnOpcode = 0x269 // Updated for 6.01
	PlayerSpawnOpcode = 0xF7  // Updated for 6.01
	NPCSpawnOpcode    = 0x39E // Updated for 6.01
	NPCSpawn2Opcode   = 0x2C9 // Updated for 6.01

	MovementOpcode = 0x189 // Updated for 6.01
	SetPosOpcode   = 0x11F // Updated for 6.01

	CastingOpcode = 0x2F2 // Updated for 6.01

	HateRankingOpcode = 0x25B // Updated for 6.01
	HateListOpcode    = 0x123 // Updated for 6.01

	EquipChangeOpcode = 0x1A4 // Updated for 6.01

	EventPlayOpcode    = 0x396 // Updated for 6.01
	EventPlay4Opcode   = 0xA6  // Updated for 6.01
	EventPlay8Opcode   = 0x1AE // Updated for 6.01
	EventPlay16Opcode  = 0x175 // Updated for 6.01
	EventPlay32Opcode  = 0x24A // Updated for 6.01
	EventPlay64Opcode  = 0x3AB // Updated for 6.01
	EventPlay128Opcode = 0x34E // Updated for 6.01
	EventPlay255Opcode = 0x109 // Updated for 6.01

	MountOpcode = 0xB7 // Updated for 6.01

	WeatherChangeOpcode = 0x148 // Updated for 6.01

	PrepareZoningOpcode = 0x2AC // Updated for 6.01

	GaugeOpcode = 0x208 // Updated for 6.01

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
