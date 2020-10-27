package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x35E // Updated for 5.35
	InitZoneOpcode      = 0x303 // Updated for 5.35
	ControlOpcode       = 0x2DC // Updated for 5.35
	ControlSelfOpcode   = 0x32C // Updated for 5.35
	ControlTargetOpcode = 0x369 // Updated for 5.35
	RemoveEntityOpcode  = 0xBC  // Updated for 5.35
	UpdateHPMPTPOpcode  = 0x153 // Updated for 5.35

	ChatZoneOpcode = 0x39F // Updated for 5.35

	UpdateStatusesOpcode       = 0x3A8 // Updated for 5.35
	UpdateStatusesEurekaOpcode = 0x2D3 // Updated for 5.35
	UpdateStatusesBossOpcode   = 0x28C // Updated for 5.35

	ActionOpcode      = 0x3A9 // Updated for 5.35
	AoEAction8Opcode  = 0x2B3 // Updated for 5.35
	AoEAction16Opcode = 0x3D7 // Updated for 5.35
	AoEAction24Opcode = 0x1AB // Updated for 5.35
	AoEAction32Opcode = 0x258 // Updated for 5.35

	PlayerSpawnOpcode = 0x38E // Updated for 5.35
	NPCSpawnOpcode    = 0x1DA // Updated for 5.35
	NPCSpawn2Opcode   = 0x346 // Updated for 5.35

	MovementOpcode = 0x2C5 // Updated for 5.35
	SetPosOpcode   = 0x1D4 // Updated for 5.35

	CastingOpcode = 0x37B // Updated for 5.35

	HateRankingOpcode = 0xA9  // Updated for 5.35
	HateListOpcode    = 0x3AE // Updated for 5.35

	EquipChangeOpcode = 0x1A2 // Updated for 5.35

	EventPlayOpcode   = 0x39A // Updated for 5.35
	EventPlay4Opcode  = 0x382 // Updated for 5.35
	EventPlay32Opcode = 0x115 // Updated for 5.35

	MountOpcode = 0x3D4 // Updated for 5.35

	WeatherChangeOpcode = 0x1AA // Updated for 5.35

	PrepareZoningOpcode = 0x160 // Updated for 5.35

	GaugeOpcode = 0xEB // Updated for 5.35

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
