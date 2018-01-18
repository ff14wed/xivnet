package datatypes

// WeaponGear defines the data struct for a weapon gear
type WeaponGear struct {
	ModelMain1, ModelMain2, ModelMain3, ModelMain4 uint16
	ModelSub1, ModelSub2, ModelSub3, ModelSub4     uint16
}

// Gear defines the data struct for a normal gear
type Gear struct {
	ModelID uint16
	Variant uint8
	Dye     uint8
}
