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

// StringToFCTag converts a string to FC tag data
func StringToFCTag(s string) FCTag {
	var e FCTag
	b := []byte(s)
	copy(e[:], b)
	return e
}

// String returns the string represenation of the FC tag data
func (e FCTag) String() string {
	end := bytes.IndexByte(e[:], 0)
	return string(e[:end])
}

// MarshalJSON writes the string representation of the FC tag to JSON
func (e FCTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

// UnmarshalJSON converts the string representation of the FC tag to
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

// FCName is the storage structure for a UTF-8 FC name
type FCName [46]byte

// StringToFCName converts a string to FC name data
func StringToFCName(s string) FCName {
	var e FCName
	b := []byte(s)
	copy(e[:], b)
	return e
}

// String returns the string represenation of the FC name data
func (e FCName) String() string {
	end := bytes.IndexByte(e[:], 0)
	return string(e[:end])
}

// MarshalJSON writes the string representation of the FC name to JSON
func (e FCName) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

// UnmarshalJSON converts the string representation of the FC name to
// the data structure
func (e *FCName) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	copy(e[:], []byte(s))
	return nil
}

type ChatMessage [1024]byte

// StringToChatMessage converts a string to an chat message data
func StringToChatMessage(s string) ChatMessage {
	var e ChatMessage
	b := []byte(s)
	copy(e[:], b)
	return e
}

// String returns the string represenation of the entity name data
func (e ChatMessage) String() string {
	end := bytes.IndexByte(e[:], 0)
	return string(e[:end])
}

// MarshalJSON writes the string representation of the chat message to JSON
func (e ChatMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

// UnmarshalJSON converts the string representation of the chat message to
// the data structure
func (e *ChatMessage) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	copy(e[:], []byte(s))
	return nil
}
