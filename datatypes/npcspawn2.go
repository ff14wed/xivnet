package datatypes

// NPCSpawn2 defines the data array for a new entity block
// Notes:
// GMRank from PlayerSpawn corresponds to MobAggression in
// this struct
// This packet type is encountered in the wild when spawning Alliance Raid
// bosses
type NPCSpawn2 struct {
	PlayerSpawn
	U31 [47]uint64
}

func (NPCSpawn2) IsBlockData() {}
