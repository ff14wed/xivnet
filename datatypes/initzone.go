package datatypes

// InitZone defines the data array for a map change block
type InitZone struct {
	U1a             uint16
	TerritoryTypeID uint16
	U1b             uint16
	U2              uint16

	U3, U4             uint32
	WeatherID, Bitmask byte
	U5b                uint16

	U6 [12]uint32

	X, Y, Z float32

	U7 [4]uint32
}

func (InitZone) IsBlockData() {}
