package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x11F // Updated for 5.41
	InitZoneOpcode      = 0x381 // Updated for 5.41
	ControlOpcode       = 0x278 // Updated for 5.41
	ControlSelfOpcode   = 0x3A3 // Updated for 5.41
	ControlTargetOpcode = 0xBA  // Updated for 5.41
	RemoveEntityOpcode  = 0x1E1 // Updated for 5.41
	UpdateHPMPTPOpcode  = 0xED  // Updated for 5.41

	ChatZoneOpcode = 0xCF // Updated for 5.41

	UpdateStatusesOpcode       = 0x14C // Updated for 5.41
	UpdateStatusesEurekaOpcode = 0x15A // Updated for 5.41
	UpdateStatusesBossOpcode   = 0x2DB // Updated for 5.41

	ActionOpcode      = 0x13F // Updated for 5.41
	AoEAction8Opcode  = 0x12A // Updated for 5.41
	AoEAction16Opcode = 0xF4  // Updated for 5.41
	AoEAction24Opcode = 0x382 // Updated for 5.41
	AoEAction32Opcode = 0x217 // Updated for 5.41

	ObjectSpawnOpcode = 0x1EA // Updated for 5.41
	PlayerSpawnOpcode = 0x283 // Updated for 5.41
	NPCSpawnOpcode    = 0x251 // Updated for 5.41
	NPCSpawn2Opcode   = 0x186 // Updated for 5.41

	MovementOpcode = 0x31A // Updated for 5.41
	SetPosOpcode   = 0x159 // Updated for 5.41

	CastingOpcode = 0x244 // Updated for 5.41

	HateRankingOpcode = 0x2FD // Updated for 5.41
	HateListOpcode    = 0x87  // Updated for 5.41

	EquipChangeOpcode = 0xB3 // Updated for 5.41

	EventPlayOpcode    = 0x346 // Updated for 5.41
	EventPlay4Opcode   = 0x274 // Updated for 5.41
	EventPlay8Opcode   = 0x138 // Updated for 5.41
	EventPlay16Opcode  = 0x3AA // Updated for 5.41
	EventPlay32Opcode  = 0x1E7 // Updated for 5.41
	EventPlay64Opcode  = 0x234 // Updated for 5.41
	EventPlay128Opcode = 0xFD  // Updated for 5.41
	EventPlay255Opcode = 0x2E1 // Updated for 5.41

	MountOpcode = 0x148 // Updated for 5.41

	WeatherChangeOpcode = 0x22E // Updated for 5.41

	PrepareZoningOpcode = 0x16A // Updated for 5.41

	GaugeOpcode = 0x7F // Updated for 5.41

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
