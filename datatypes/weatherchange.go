package datatypes

// WeatherChange defines the data array for a weather change block
type WeatherChange struct {
	ID    byte
	U1    byte
	U2    byte
	U3    byte
	Delay float32
}

func (WeatherChange) IsBlockData() {}
