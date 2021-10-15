package object

import (
	"github.com/whs/whscreeps/screeps/game"
	"github.com/whs/whscreeps/screeps/utils"
)

var Spawns = game.Game.Get("spawns")

func GetSpawn(name string) Spawn {
	return Spawn{Structure: Structure{RoomObject: RoomObject{Raw: Spawns.Get(name)}}}
}

// GetSpawnNames return list of all spawns
func GetSpawnNames() []string {
	return utils.ObjectKeys(Spawns)
}

// GetSpawns get all spawns. This function is very expensive.
func GetSpawns() map[string]Spawn {
	names := GetCreepNames()

	out := make(map[string]Spawn, len(names))
	for _, name := range names {
		out[name] = GetSpawn(name)
	}

	return out
}
