package datatypes

// MapChange defines the data array for a map change block
type MapChange struct {
	U1              uint32
	TerritoryTypeID uint16
	U2              uint16
	U3, U4          uint32
}

func (MapChange) IsBlockData() {}
