package datatypes

// PlayerStats defines the data array for player stats
type PlayerStats struct {
	Strength            uint32
	Dexterity           uint32
	Vitality            uint32
	Intelligence        uint32
	Mind                uint32
	Piety               uint32
	HP                  uint32
	MP                  uint32
	TP                  uint32
	GP                  uint32
	CP                  uint32
	Delay               uint32
	Tenacity            uint32
	AttackPower         uint32
	Defense             uint32
	DirectHitRate       uint32
	Evasion             uint32
	MagicDefense        uint32
	CriticalHit         uint32
	AttackMagicPotency  uint32
	HealingMagicPotency uint32
	ElementalBonus      uint32
	Determination       uint32
	SkillSpeed          uint32
	SpellSpeed          uint32
	Haste               uint32
	Craftsmanship       uint32
	Control             uint32
	Gathering           uint32
	Perception          uint32

	U1 [26]uint32
}

func (PlayerStats) IsBlockData() {}
