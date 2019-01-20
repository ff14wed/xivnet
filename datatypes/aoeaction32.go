package datatypes

// AoEAction32 defines the data array for an AoE action block
type AoEAction32 struct {
	ActionHeader

	EffectsList [32]ActionEffects

	U9  uint32
	U10 uint16

	Targets [32]uint64

	Position PackedPosition
	U11      uint16
	U12      uint32
}

func (AoEAction32) IsBlockData() {}
