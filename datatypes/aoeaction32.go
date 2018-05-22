package datatypes

// AoEAction32 defines the data array for an AoE action block
type AoEAction32 struct {
	TargetID     uint32
	U1           uint32
	ActionIDName uint32
	U2           uint32
	U3           uint32
	UnkID1       uint32
	U4           uint16
	Direction    uint16 // Quantized direction 0x0000 ~ 0xFFFF, NWSE <=> 0,0x4000,0x8000,0xC000
	ActionID     uint16
	U5           uint16

	U6a         byte
	NumAffected byte
	U6b         uint16
	U7          uint32
	U8          uint16

	EffectsList [32]ActionEffects

	U9  uint32
	U10 uint16

	Targets [32]uint64

	Position PackedPosition
	U11      uint16
	U12      uint32
}
