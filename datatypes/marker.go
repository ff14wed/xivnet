package datatypes

// Marker defines the data array for a place marker block
type Marker struct {
	Type byte
	U1   byte
	U2   uint16
	U3   uint32
	X    float32
	Z    float32
	Y    float32
	U4   uint32
}

func (Marker) IsBlockData() {}
