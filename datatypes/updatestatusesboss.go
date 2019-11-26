package datatypes

// UpdateStatusesBoss defines the data array for an update statuses block that supports 60 status effects
type UpdateStatusesBoss struct {
	Statuses2 [30]StatusEffect
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

func (UpdateStatusesBoss) IsBlockData() {}
