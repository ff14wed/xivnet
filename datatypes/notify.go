package datatypes

// Notify142 defines the data array for a notify block
type Notify142 struct {
	Type uint16
	Pad1 uint16
	P1   uint32
	P2   uint32
	P3   uint32
	P4   uint32
	Pad2 uint32
}

/* Types
Type:6 => Death
  P2 is the the killer
Type:15 => Cancel Ability, P3: Ability ID
Type 17 => HoT/DoT tick
Type:21 => Remove Buff
Type:34 => Target icon over player.

Below are unconfirmed now that Stormblood arrived.
Type:0, P1:1 => Unsheath
Type:0, P1:0 => Sheath
Type:22, P1: Status slot, P2: Status ID, P3: Extra
Type:23, P2:3  => DoT
Type:23, P2:4  => HoT
Type:503 => Aggro???
  P2 is target ID
Type:39 => Remove Entity ???
*/

func (Notify142) IsBlockData() {}

// Notify143 defines the data array for a notify3 block
type Notify143 struct {
	Type uint16
	Pad1 uint16
	P1   uint32
	P2   uint32
	P3   uint32
	P4   uint32
	P5   uint32
	P6   uint32
	Pad2 uint32
}

func (Notify143) IsBlockData() {}

// Notify144 defines the data array for a notify4 block
type Notify144 struct {
	Type     uint16
	Pad1     uint16
	P1       uint32
	P2       uint32
	P3       uint32
	P4       uint32
	Pad2     uint32
	TargetID uint32 // 0xE0000000 means deselect target
	U1       uint32
}

func (Notify144) IsBlockData() {}
