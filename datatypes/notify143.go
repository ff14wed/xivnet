package datatypes

// Notify143 defines the data array for a notify3 block
type Notify143 struct {
	Type uint16
	Pad1 uint16
	P1   uint32
	P2   uint32
	P3   uint32
	P4   uint32
	P5   uint32
	P6   uint32
	Pad2 uint32
}

func (Notify143) IsBlockData() {}
