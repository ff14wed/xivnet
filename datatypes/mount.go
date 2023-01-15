package datatypes

// Mount defines the data array for a getting on mount block
type Mount struct {
	ID    uint16
	Color byte
	U1    byte
	U2    byte
	U3    uint32 // Highly likely to be a Gear
	U4    uint32 // Highly likely to be a Gear
	U5    uint32 // Highly likely to be a Gear
}

func (Mount) IsBlockData() {}
