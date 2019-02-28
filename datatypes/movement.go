package datatypes

// Movement defines the data array for a movement block
type Movement struct {
	Direction       uint8 // Quantized direction 0x00~0xFF. NWSE <=> 0,0x40,0x80,0xC0
	HeadRotation    uint8
	AnimationType   uint8
	AnimationState  uint8
	AnimationSpeed  uint8
	UnknownRotation uint8
	Position        PackedPosition
	U3              uint32
}

func (Movement) IsBlockData() {}
