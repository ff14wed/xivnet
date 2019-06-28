package datatypes

// EgressCraftEvent defines the data array for an crafting event
// Name subject to change
type EgressCraftEvent struct {
	EventID, U2, U3 uint32
	Ingredients     [9]uint32
	U4              [21]uint32
	Amounts         [9]byte
	U5              [3]byte
	U6              [30]uint32
}

func (EgressCraftEvent) IsBlockData() {}
