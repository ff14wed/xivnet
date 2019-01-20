package datatypes

// SetPos defines the data array for a movement block
type SetPos struct {
	Direction   uint16
	WaitForLoad byte
	U1          byte
	U2          uint32
	X           float32
	Z           float32
	Y           float32
	U3          uint32
}

func (SetPos) IsBlockData() {}
