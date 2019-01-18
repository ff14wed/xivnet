package datatypes

// WeaponGear defines the data struct for a weapon gear
type WeaponGear struct {
	Model1, Model2, Model3, Model4 uint16
}

// Gear defines the data struct for a normal gear
type Gear struct {
	ModelID uint16
	Variant uint8
	Dye     uint8
}
