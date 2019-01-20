package datatypes

// EquipChange defines the data array for a equip change block
type EquipChange struct {
	WeaponMain WeaponGear
	WeaponSub  WeaponGear
	U1a        byte
	ClassJob   byte
	U1c        byte
	U1d        byte
	Head       Gear
	Body       Gear
	Hand       Gear
	Leg        Gear
	Foot       Gear
	Ear        Gear
	Neck       Gear
	Wrist      Gear
	Ring1      Gear
	Ring2      Gear
	U2         uint32
}

func (EquipChange) IsBlockData() {}
