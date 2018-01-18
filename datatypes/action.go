package datatypes

// Action defines the data array for an ability block
type Action struct {
	TargetID     uint32 // Target of actions
	U1           uint32
	ActionIDName uint32
	U2           uint32
	U3           uint32
	U4           uint32
	UnkID1       uint32
	Direction    uint16 // Quantized direction 0x0000 ~ 0xFFFF, NWSE <=> 0,0x4000,0x8000,0xC000
	ActionID     uint16

	U6a         byte
	U6b         byte
	U7a         byte
	NumAffected byte

	Pad1    uint32
	Effects ActionEffects

	TargetID2 uint32 // Target of effects
	U9        uint32
	U10       uint32
}

// ActionEffects defines a block of 8 action effects
type ActionEffects [8]ActionEffect

// ActionEffect defines the data array for an effect that resulted from the action
type ActionEffect struct {
	Type byte
	P1   byte
	P2   byte
	P3   byte // Usually affects percentage message in battle log
	P4   uint16
	P5   byte // 0 affects target, 128 affects self
	P6   byte
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
