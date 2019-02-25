package datatypes

// PrepareZoning defines the data array for a prepare zoning block
type PrepareZoning struct {
	U1              uint32
	TerritoryTypeID uint16
	U2              uint16
	U3, U4          uint32
}

func (PrepareZoning) IsBlockData() {}
