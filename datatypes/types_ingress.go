package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x30B // Updated for 6.1
	InitZoneOpcode      = 0x1E7 // Updated for 6.1
	ControlOpcode       = 0x2E7 // Updated for 6.1
	ControlSelfOpcode   = 0x28F // Updated for 6.1
	ControlTargetOpcode = 0x399 // Updated for 6.1
	RemoveEntityOpcode  = 0x2FD // Updated for 6.1
	UpdateHPMPTPOpcode  = 0x94  // Updated for 6.1

	ChatZoneOpcode = 0x19C // Updated for 6.1

	UpdateStatusesOpcode       = 0x31D // Updated for 6.1
	UpdateStatusesEurekaOpcode = 0x31E // Updated for 6.1
	UpdateStatusesBossOpcode   = 0x11A // Updated for 6.1

	ActionOpcode      = 0x21B // Updated for 6.1
	AoEAction8Opcode  = 0x30E // Updated for 6.1
	AoEAction16Opcode = 0x153 // Updated for 6.1
	AoEAction24Opcode = 0xE1  // Updated for 6.1
	AoEAction32Opcode = 0x356 // Updated for 6.1

	ObjectSpawnOpcode = 0x11D // Updated for 6.1
	PlayerSpawnOpcode = 0x2BC // Updated for 6.1
	NPCSpawnOpcode    = 0x12F // Updated for 6.1
	NPCSpawn2Opcode   = 0x20A // Updated for 6.1

	MovementOpcode = 0x366 // Updated for 6.1
	SetPosOpcode   = 0x23A // Updated for 6.1

	CastingOpcode = 0x6F // Updated for 6.1

	HateRankingOpcode = 0x89  // Updated for 6.1
	HateListOpcode    = 0x2A2 // Updated for 6.1

	EquipChangeOpcode = 0x2E6 // Updated for 6.1

	EventPlayOpcode    = 0xF3  // Updated for 6.1
	EventPlay4Opcode   = 0x2DB // Updated for 6.1
	EventPlay8Opcode   = 0x17B // Updated for 6.1
	EventPlay16Opcode  = 0x369 // Updated for 6.1
	EventPlay32Opcode  = 0x131 // Updated for 6.1
	EventPlay64Opcode  = 0x256 // Updated for 6.1
	EventPlay128Opcode = 0x337 // Updated for 6.1
	EventPlay255Opcode = 0x17F // Updated for 6.1

	MountOpcode = 0x2E2 // Updated for 6.1

	WeatherChangeOpcode = 0x2F8 // Updated for 6.1

	PrepareZoningOpcode = 0x2C9 // Updated for 6.1

	GaugeOpcode = 0x1BE // Updated for 6.1

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
