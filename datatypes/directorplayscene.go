package datatypes

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"

	"github.com/ff14wed/xivnet/v3/internal/bytestring"
)

type PlaySceneData interface {
	isPlaySceneData()
}

// DirectorPlayScene defines the data array for a event scene block
// Note: This requires an EventStart block.
// The event play header might change as we figure out more information about
// it.
type DirectorPlayScene struct {
	EventPlay
	Data PlaySceneData
}

func (DirectorPlayScene) IsBlockData() {}

// UnmarshalBytes decodes the provided raw bytes into the DirectorPlayScene structure
func (d *DirectorPlayScene) UnmarshalBytes(data []byte) error {
	headerLen := binary.Size(d.EventPlay)
	if len(data) != (headerLen + 120) {
		return fmt.Errorf("length mismatch: %d != %d", len(data), headerLen+120)
	}

	err := binary.Read(bytes.NewReader(data[:headerLen]), binary.LittleEndian, &d.EventPlay)
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
		c := GenericPlaySceneData{}
		err := binary.Read(bytes.NewReader(data[headerLen:]), binary.LittleEndian, &c)
		if err != nil {
			return err
		}
		d.Data = c
	}
	return nil
}

// UnmarshalJSON decodes a JSON DirectorPlayScene struct with either a crafting
// state data or generic byte data
func (d *DirectorPlayScene) UnmarshalJSON(data []byte) error {
	type genericPlayScene struct {
		EventPlay
		Data json.RawMessage
	}
	g := genericPlayScene{}
	err := json.Unmarshal(data, &g)
	if err != nil {
		return err
	}
	d.EventPlay = g.EventPlay
	c := CraftState{}
	err = json.Unmarshal(g.Data, &c)
	if err == nil {
		d.Data = c
		return nil
	}

	p := GenericPlaySceneData{}
	err = json.Unmarshal(g.Data, &p)
	if err == nil {
		d.Data = p
		return nil
	}

	return err
}

type GenericPlaySceneData [120]byte

func (GenericPlaySceneData) isPlaySceneData() {}

// MarshalJSON returns the marshaled version of the GenericPlaySceneData
func (p GenericPlaySceneData) MarshalJSON() ([]byte, error) {
	return bytestring.BytesToByteString(p[:])
}

// UnmarshalJSON returns the unmarshaled version of the GenericPlaySceneData
func (p *GenericPlaySceneData) UnmarshalJSON(data []byte) error {
	pb, err := bytestring.ByteStringToBytes(data)
	if err != nil {
		return err
	}
	if len(pb) != len(p) {
		return fmt.Errorf("length mismatch: %d != %d", len(pb), len(p))
	}
	copy(p[:], pb)
	return nil
}

type CraftState struct {
	U1                uint32
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

func (CraftState) isPlaySceneData() {}
