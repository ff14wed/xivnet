package datatypes

// ActionHeader defines the header for an action packet. This is common
// for single target or multi-target action packets.
type ActionHeader struct {
	TargetID          uint32 // Target of actions
	U1                uint32
	ActionIDName      uint32 // Action as shown in battle log
	GlobalCounter     uint32
	AnimationLockTime float32
	UnkID1            uint32
	HiddenAnimation   uint16
	Direction         uint16 // Quantized direction 0x0000 ~ 0xFFFF, NWSE <=> 0,0x4000,0x8000,0xC000
	ActionID          uint16 // The animation of the action
	Variation         byte
	EffectDisplayType byte

	U6a         byte
	NumAffected byte
	U6b         uint16
	U7          uint32
	U8          uint16
}

// Action defines the data array for an ability block
type Action struct {
	ActionHeader

	Effects ActionEffects

	U9  uint32
	U10 uint16

	TargetID2   uint32 // Target of effects
	EffectFlags uint32
	U12         uint32
}

func (Action) IsBlockData() {}

// ActionEffects defines a block of 8 action effects
type ActionEffects [8]ActionEffect

// ActionEffect defines the data array for an effect that resulted from the action
type ActionEffect struct {
	Type        byte
	HitSeverity byte
	P3          byte
	Percentage  byte
	// Total Damage = 65535 * Multiplier * (Flags & 0x40) + Damage
	Multiplier byte
	Flags      byte // (Flags & 0xA0) means attacker receives damage instead
	Damage     uint16
}

/*
	POTENTIALLY OUTDATED INFORMATION:
	Type 1 - Dodge/Miss.
	Type 2 - Fully resisted.
  Type 3 - Normal damage, P4 is damage
         - P1 is a bitfield. 0 is normal, 1 is crit hit, 2 is direct hit, 3 is both.
	Type 4 - HP recover message (P1: 1, P4: HP recovered)
	Type 5 - Blocked damage. P3 is blocked percentage. P4 is damage.
	Type 6 - Parried damage. P3 is blocked percentage. P4 is damage.
	Type 7 - Invulnerable
	Type 8 - Has no effect.
	Type 9 - Bane spreading DoTs
	Type 10 - You lose MP (P4: MP lost).
	Type 11 - MP recover message (P4: MP recovered).
	Type 12 - You lose TP (P4: TP lost).
	Type 13 - TP recover message (P4: TP recovered).
	Type 14 - GP recover message (P4: TP recovered)
	Type 15(or 16) - Inflicted status effect
          - P1 is base damage
          - P2 is crit % * 10. For damage debuff, it indicates percent modifier.
          - P4 is status ID
	Type 17-19 - Recovered from the effect of P4 status ID
	Type 21 - Status P4 has no effect.
	Type 25-26 - Enmity increases message
  Type 28 - Used Skill ID, P4 is skill ID
	Type 38 - Mounting a chocobo/mount.
	Type 46 - Suffer bad breath statuses (poison, blind, slow, heavy, damage down)
	Type 48 - Mounting a turret.
	Type 51 - Fully resist status P4
	Type 53 - Sentenced to death message.
	Type 55 - Lost the effect of sheltron. Recovered P4 MP.
	Type 57 - Player teleport.
	Type 58 - ??
	Type 61 - Reflected! You cast back the attack.
*/
