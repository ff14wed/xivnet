package datatypes

// NPCSpawn defines the data array for a new entity block
// Notes:
// GMRank from PlayerSpawn corresponds to MobAggression in
// this struct
type NPCSpawn struct {
	PlayerSpawn
	U31 uint64
}

func (NPCSpawn) IsBlockData() {}
