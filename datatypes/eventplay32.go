package datatypes

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
)

type EventPlay32Data interface {
	isEventPlay32Data()
}

// EventPlay32 defines the data array for a event scene block
// Note: This requires an EventStart block.
type EventPlay32 struct {
	EventPlayHeader
	Data EventPlay32Data
}

func (EventPlay32) IsBlockData() {}

// UnmarshalBytes decodes the provided raw bytes into the EventPlay32 structure
func (d *EventPlay32) UnmarshalBytes(data []byte) error {
	headerLen := binary.Size(d.EventPlayHeader)
	if len(data) != (headerLen + 128) {
		return fmt.Errorf("length mismatch: %d != %d", len(data), headerLen+128)
	}

	err := binary.Read(bytes.NewReader(data[:headerLen]), binary.LittleEndian, &d.EventPlayHeader)
	if err != nil {
		return err
	}

	switch d.EventID {
	case 0xA0001:
		c := CraftState{}
		err := binary.Read(bytes.NewReader(data[headerLen:]), binary.LittleEndian, &c)
		if err != nil {
			return err
		}
		d.Data = c
	default:
		c := GenericEventPlay32Data{}
		err := binary.Read(bytes.NewReader(data[headerLen:]), binary.LittleEndian, &c)
		if err != nil {
			return err
		}
		d.Data = c
	}
	return nil
}

// UnmarshalJSON decodes a JSON EventPlay32 struct with either a crafting
// state data or generic byte data
func (d *EventPlay32) UnmarshalJSON(data []byte) error {
	type genericEventPlay32 struct {
		EventPlayHeader
		Data json.RawMessage
	}
	g := genericEventPlay32{}
	err := json.Unmarshal(data, &g)
	if err != nil {
		return err
	}
	d.EventPlayHeader = g.EventPlayHeader
	c := CraftState{}
	err = json.Unmarshal(g.Data, &c)
	if err == nil {
		d.Data = c
		return nil
	}

	p := GenericEventPlay32Data{}
	err = json.Unmarshal(g.Data, &p)
	if err == nil {
		d.Data = p
		return nil
	}

	return err
}

type GenericEventPlay32Data [32]uint32

func (GenericEventPlay32Data) isEventPlay32Data() {}

type CraftState struct {
	U1                uint32
	U3                uint32
	U4                uint32
	CraftAction       uint32
	U2                uint32
	StepNum           uint32
	Progress          uint32
	ProgressDelta     int32
	Quality           uint32
	QualityDelta      int32
	HQChance          uint32
	Durability        uint32
	DurabilityDelta   int32
	CurrentCondition  uint32
	PreviousCondition uint32
	U6                [17]uint32
}

func (CraftState) isEventPlay32Data() {}
