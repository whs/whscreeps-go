package task

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/whs/whscreeps/os"
	"github.com/whs/whscreeps/screeps/game"
	"github.com/whs/whscreeps/screeps/object"
	"github.com/whs/whscreeps/whscreeps/store"
)

type Main struct {
}

func (m Main) Name() string {
	return "Main"
}

func (m Main) Execute(ctx context.Context) {
	logger := zerolog.Ctx(ctx)
	root := store.CtxStore(ctx)
	scheduler := os.CtxScheduler(ctx)

	for name, creepData := range root.Creeps {
		if creepData.ProbeLock() {
			continue
		}
		creep, err := object.GetCreep(name)
		if err == object.ErrNoCreep {
			continue
		}
		if creep.Spawning() {
			continue
		}

		logger.Info().Str("name", name).Msg("Found idle creep - assigning to harvest")
		token := creepData.Lock()
		scheduler.Schedule(&Harvest{
			CreepID: name,
			Token:   token,
		}, os.PriorityVeryLow)
	}

	scheduler.ScheduleTransient(BuildHarvester{Spawn: "Spawn1"}, os.PriorityVeryLow)

	currentTick := game.GetTime()
	timeSinceLastHousekeeping := currentTick - root.LastHousekeeping
	if timeSinceLastHousekeeping > 10 {
		scheduler.ScheduleTransient(Housekeeping{}, os.PriorityVeryHigh)
	} else {
		scheduler.ScheduleTransient(Housekeeping{}, os.PriorityBackground)
	}
}
