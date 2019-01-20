package datatypes

// AoEAction16 defines the data array for an AoE action block
type AoEAction16 struct {
	ActionHeader

	EffectsList [16]ActionEffects

	U9  uint32
	U10 uint16

	Targets [16]uint64

	Position PackedPosition
	U11      uint16
	U12      uint32
}

func (AoEAction16) IsBlockData() {}
