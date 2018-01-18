package datatypes

import (
	"bytes"
	"encoding/json"
)

// EntityName is the storage structure for a UTF-8 entity name
type EntityName [32]byte

// StringToEntityName converts a string to an entity name data
func StringToEntityName(s string) EntityName {
	var e EntityName
	b := []byte(s)
	copy(e[:], b)
	return e
}

// String returns the string represenation of the entity name data
func (e EntityName) String() string {
	end := bytes.IndexByte(e[:], 0)
	return string(e[:end])
}

// MarshalJSON writes the string representation of the entity name to JSON
func (e EntityName) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

// UnmarshalJSON converts the string representation of the entity name to
// the data structure
func (e *EntityName) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	copy(e[:], []byte(s))
	return nil
}

// FCTag is the storage structure for a UTF-8 FC tag
type FCTag [6]byte

// StringToFCTag converts a string to an entity name data
func StringToFCTag(s string) FCTag {
	var e FCTag
	b := []byte(s)
	copy(e[:], b)
	return e
}

// String returns the string represenation of the entity name data
func (e FCTag) String() string {
	end := bytes.IndexByte(e[:], 0)
	return string(e[:end])
}

// MarshalJSON writes the string representation of the entity name to JSON
func (e FCTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

// UnmarshalJSON converts the string representation of the entity name to
// the data structure
func (e *FCTag) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	copy(e[:], []byte(s))
	return nil
}
