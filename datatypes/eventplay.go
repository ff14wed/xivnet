package datatypes

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type EventPlayHeader struct {
	ActorID    uint64
	EventID    uint32 // 0xA0001 for crafting
	Scene      uint16
	Pad1       uint16
	Flags      uint32
	P1         uint32
	ParamCount byte
	Pad2       [3]byte
	P2         uint32
}

const CraftingEventID = 0xA0001

// EventPlay defines the data array for an event play block
// Note: This requires an EventStart block.
type EventPlay struct {
	EventPlayHeader
	Data [2]uint32
}

func (EventPlay) IsBlockData() {}
func (e *EventPlay) Header() *EventPlayHeader {
	return &e.EventPlayHeader
}
func (e *EventPlay) EventData() []uint32 {
	return e.Data[:]
}

// EventPlay4 defines the data array for an event play block
// Note: This requires an EventStart block.
type EventPlay4 struct {
	EventPlayHeader
	Data [4]uint32
}

func (EventPlay4) IsBlockData() {}
func (e *EventPlay4) Header() *EventPlayHeader {
	return &e.EventPlayHeader
}
func (e *EventPlay4) EventData() []uint32 {
	return e.Data[:]
}

// EventPlay8 defines the data array for an event play block
type EventPlay8 struct {
	EventPlayHeader
	Data [8]uint32
}

func (EventPlay8) IsBlockData() {}
func (e *EventPlay8) Header() *EventPlayHeader {
	return &e.EventPlayHeader
}
func (e *EventPlay8) EventData() []uint32 {
	return e.Data[:]
}

// EventPlay16 defines the data array for an event play block
type EventPlay16 struct {
	EventPlayHeader
	Data [16]uint32
}

func (EventPlay16) IsBlockData() {}
func (e *EventPlay16) Header() *EventPlayHeader {
	return &e.EventPlayHeader
}
func (e *EventPlay16) EventData() []uint32 {
	return e.Data[:]
}

// EventPlay32 defines the data array for an event play block
type EventPlay32 struct {
	EventPlayHeader
	Data [32]uint32
}

func (EventPlay32) IsBlockData() {}
func (e *EventPlay32) Header() *EventPlayHeader {
	return &e.EventPlayHeader
}
func (e *EventPlay32) EventData() []uint32 {
	return e.Data[:]
}

// EventPlay64 defines the data array for an event play block
type EventPlay64 struct {
	EventPlayHeader
	Data [64]uint32
}

func (EventPlay64) IsBlockData() {}
func (e *EventPlay64) Header() *EventPlayHeader {
	return &e.EventPlayHeader
}
func (e *EventPlay64) EventData() []uint32 {
	return e.Data[:]
}

// EventPlay128 defines the data array for an event play block
type EventPlay128 struct {
	EventPlayHeader
	Data [128]uint32
}

func (EventPlay128) IsBlockData() {}
func (e *EventPlay128) Header() *EventPlayHeader {
	return &e.EventPlayHeader
}
func (e *EventPlay128) EventData() []uint32 {
	return e.Data[:]
}

// EventPlay255 defines the data array for an event play block
type EventPlay255 struct {
	EventPlayHeader
	Data [255]uint32
}

func (EventPlay255) IsBlockData() {}
func (e *EventPlay255) Header() *EventPlayHeader {
	return &e.EventPlayHeader
}
func (e *EventPlay255) EventData() []uint32 {
	return e.Data[:]
}

// IEventPlay defines the interface for the event play packets
type IEventPlay interface {
	Header() *EventPlayHeader
	EventData() []uint32
}

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
	U5                uint32
	Flags             uint32
	U6                [15]uint32
}

func MarshalCraftState(cs *CraftState) ([]uint32, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, cs)
	if err != nil {
		return nil, fmt.Errorf("MarshalCraftState: writing data to buffer: %s", err)
	}

	byteData := buf.Bytes()
	var data []uint32
	for i := 0; i < len(byteData); i += 4 {
		data = append(data, binary.LittleEndian.Uint32(byteData[i:i+4]))
	}
	return data, nil
}

func UnmarshalCraftState(ep IEventPlay) (*CraftState, error) {
	if ep.Header().EventID != CraftingEventID {
		return nil, fmt.Errorf("UnmarshalCraftState: input data is not a craft state")
	}
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, ep.EventData())
	if err != nil {
		return nil, fmt.Errorf("UnmarshalCraftState: writing data to buffer: %s", err)
	}
	craftState := CraftState{}
	craftStateLen := binary.Size(craftState)
	if buf.Len() < craftStateLen {
		return nil, fmt.Errorf("UnmarshalCraftState: not enough data in buffer: %d < %d", buf.Len(), craftStateLen)
	}
	err = binary.Read(buf, binary.LittleEndian, &craftState)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalCraftState: reading data from buffer: %s", err)
	}
	return &craftState, nil
}
