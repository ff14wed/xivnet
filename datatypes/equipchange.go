package datatypes

// EquipChange defines the data array for a equip change block
type EquipChange struct {
	Weapon WeaponGear
	U1     uint32
	Head   Gear
	Body   Gear
	Hand   Gear
	Leg    Gear
	Foot   Gear
	Ear    Gear
	Neck   Gear
	Wrist  Gear
	Ring1  Gear
	Ring2  Gear
	U2     uint32
}
