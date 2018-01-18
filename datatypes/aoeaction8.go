package datatypes

// AoEAction8 defines the data array for an AoE action block
type AoEAction8 struct {
	TargetID     uint32
	U1           uint32
	ActionIDName uint32
	U2           uint32
	U3           uint32
	U4           uint32
	UnkID1       uint32
	Direction    uint16 // Quantized direction 0x0000 ~ 0xFFFF, NWSE <=> 0,0x4000,0x8000,0xC000
	ActionID     uint16

	U6a         uint8
	U6b         uint8
	U7a         byte
	NumAffected byte

	Pad1        uint32
	EffectsList [8]ActionEffects

	Targets [8]uint64

	Position PackedPosition
	U9       uint16
	U10      uint32
}
