package datatypes

// Movement defines the data array for a movement block
type Movement struct {
	Direction uint8 // Quantized direction 0x00~0xFF. NWSE <=> 0,0x40,0x80,0xC0
	U1        uint8
	U2        uint32
	Position  PackedPosition
	U3        uint32
}

func (Movement) IsBlockData() {}
