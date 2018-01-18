package datatypes

import "encoding/json"

// PackedCoord represents a 16 bit packed representation of a float32
type PackedCoord uint16

// PackedPosition defines the data array for position
type PackedPosition struct {
	X, Z, Y PackedCoord
}

// Float returns the floating point representation of the coord
func (pc PackedCoord) Float() float32 {
	return (float32(pc) * 2000 / 65535) - 1000
}

// SetFloat sets the PackedCoord to the uint16 representation of the floating
// point coord
func (pc *PackedCoord) SetFloat(coord float32) {
	pos := 65535 * (coord + 1000) / 2000
	*pc = PackedCoord(uint16(pos + 0.5))
}

// MarshalJSON marhals the packed coordinate as a floating point value
func (pc PackedCoord) MarshalJSON() ([]byte, error) {
	return json.Marshal(pc.Float())
}

// UnmarshalJSON marhals the floating point value as a packed coordinate
func (pc *PackedCoord) UnmarshalJSON(data []byte) error {
	var v float32
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	pc.SetFloat(v)
	return nil
}
