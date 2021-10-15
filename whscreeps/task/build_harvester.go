package task

import (
	"context"
	"encoding/gob"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/whs/whscreeps/screeps/game"
	"github.com/whs/whscreeps/screeps/object"
	"github.com/whs/whscreeps/screeps/types"
)

func init() {
	gob.Register(BuildHarvester{})
}

type BuildHarvester struct {
	Spawn string
}

func (b BuildHarvester) Name() string {
	return fmt.Sprintf("BuildHarvester(%s)", b.Spawn)
}

func (b BuildHarvester) Execute(ctx context.Context) {
	logger := zerolog.Ctx(ctx)
	spawn := object.GetSpawn("Spawn1")
	if spawn.Spawning() != nil {
		return
	}
	name := fmt.Sprintf("Harvester%d", game.GetTime())
	err := spawn.SpawnCreep([]types.BodyPart{types.Work, types.Carry, types.Move}, name, nil)
	if err != nil {
		return
	}
	logger.Debug().Str("name", name).Msg("Spawning new harvester")
}
