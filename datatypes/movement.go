package datatypes

// Movement defines the data array for a movement block
type Movement struct {
	HeadRotation    uint8 // Doesn't seem to do anything though? Retained to not break anything.
	Direction       uint8 // Quantized direction 0x00~0xFF. NWSE <=> 0,0x40,0x80,0xC0
	AnimationType   uint8
	AnimationState  uint8
	AnimationSpeed  uint8
	UnknownRotation uint8
	Position        PackedPosition
	U3              uint32
}

func (Movement) IsBlockData() {}
