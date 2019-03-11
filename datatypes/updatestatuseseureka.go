package datatypes

// UpdateStatusesEureka defines the data array for an update statuses block inside Eureka
type UpdateStatusesEureka struct {
	ElementalLevel byte
	U1             [3]byte
	ClassJob       byte
	Level1         byte
	Level          byte
	MaxLevel       byte
	CurrentHP      uint32
	MaxHP          uint32
	CurrentMP      uint16
	MaxMP          uint16
	CurrentTP      uint16
	U2             uint16

	Statuses [30]StatusEffect
}

func (UpdateStatusesEureka) IsBlockData() {}
