package datatypes

// BeginCrafting defines the data array for a begin crafting message
type BeginCrafting struct {
	U1, U2, U3  uint32
	Ingredients [9]uint32
	U4          [21]uint32
	Amounts     [9]byte
	U5          [3]byte
	U6          [30]uint32
}