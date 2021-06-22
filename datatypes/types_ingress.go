package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x3CF // Updated for 5.57
	InitZoneOpcode      = 0x1E5 // Updated for 5.57
	ControlOpcode       = 0x164 // Updated for 5.57
	ControlSelfOpcode   = 0x356 // Updated for 5.57
	ControlTargetOpcode = 0x281 // Updated for 5.57
	RemoveEntityOpcode  = 0x235 // Updated for 5.57
	UpdateHPMPTPOpcode  = 0x3BE // Updated for 5.57

	ChatZoneOpcode = 0x372 // Updated for 5.57

	UpdateStatusesOpcode       = 0x192 // Updated for 5.57
	UpdateStatusesEurekaOpcode = 0x227 // Updated for 5.57
	UpdateStatusesBossOpcode   = 0x2F6 // Updated for 5.57

	ActionOpcode      = 0x8F  // Updated for 5.57
	AoEAction8Opcode  = 0x247 // Updated for 5.57
	AoEAction16Opcode = 0x2C1 // Updated for 5.57
	AoEAction24Opcode = 0x295 // Updated for 5.57
	AoEAction32Opcode = 0x34C // Updated for 5.57

	ObjectSpawnOpcode = 0x254 // Updated for 5.57
	PlayerSpawnOpcode = 0x18D // Updated for 5.57
	NPCSpawnOpcode    = 0x1CE // Updated for 5.57
	NPCSpawn2Opcode   = 0x179 // Updated for 5.57

	MovementOpcode = 0x233 // Updated for 5.57
	SetPosOpcode   = 0x1A3 // Updated for 5.57

	CastingOpcode = 0x2CC // Updated for 5.57

	HateRankingOpcode = 0x1E1 // Updated for 5.57
	HateListOpcode    = 0x184 // Updated for 5.57

	EquipChangeOpcode = 0x1DF // Updated for 5.57

	EventPlayOpcode    = 0x36A // Updated for 5.57
	EventPlay4Opcode   = 0x397 // Updated for 5.57
	EventPlay8Opcode   = 0xA6  // Updated for 5.57
	EventPlay16Opcode  = 0x23D // Updated for 5.57
	EventPlay32Opcode  = 0x204 // Updated for 5.57
	EventPlay64Opcode  = 0xEF  // Updated for 5.57
	EventPlay128Opcode = 0x29C // Updated for 5.57
	EventPlay255Opcode = 0x21B // Updated for 5.57

	MountOpcode = 0x9C // Updated for 5.57

	WeatherChangeOpcode = 0x157 // Updated for 5.57

	PrepareZoningOpcode = 0x250 // Updated for 5.57

	GaugeOpcode = 0x377 // Updated for 5.57

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
