package datatypes

// UpdateStatuses defines the data array for an update statuses block
type UpdateStatuses struct {
	ClassJob  byte
	Level1    byte
	Level     uint16
	CurrentHP uint32
	MaxHP     uint32
	CurrentMP uint16
	MaxMP     uint16
	CurrentTP uint16
	U2        uint16 // 0?

	Statuses [30]StatusEffect

	U3 uint32
}

func (UpdateStatuses) IsBlockData() {}
