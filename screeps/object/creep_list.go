package object

import (
	"errors"
	"github.com/whs/whscreeps/screeps/game"
	"github.com/whs/whscreeps/screeps/utils"
)

var ErrNoCreep = errors.New("no creep by that name")
var Creeps = game.Game.Get("creeps")

func GetCreep(name string) (Creep, error) {
	creep := Creeps.Get(name)
	if creep.IsUndefined() {
		return Creep{}, ErrNoCreep
	}
	return Creep{RoomObject: RoomObject{Raw: creep}}, nil
}

// GetCreepNames return list of all creeps names
func GetCreepNames() []string {
	return utils.ObjectKeys(Creeps)
}

// GetCreeps get all creeps. This function is very expensive.
func GetCreeps() map[string]Creep {
	names := GetCreepNames()

	out := make(map[string]Creep, len(names))
	for _, name := range names {
		out[name], _ = GetCreep(name)
	}

	return out
}
