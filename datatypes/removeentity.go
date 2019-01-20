package datatypes

// RemoveEntity defines the data array for a remove entity block
type RemoveEntity struct {
	Index byte
	U1    byte
	U2    uint16
	ID    uint32
}

func (RemoveEntity) IsBlockData() {}
