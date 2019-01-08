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

	U6, U7, U8, U9, U10, U11, U12, U13, U14 uint32

	X, Z, Y float32
	U15     uint32
}
