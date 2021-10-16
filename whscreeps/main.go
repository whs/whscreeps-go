package whscreeps

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/whs/whscreeps/os"
	"github.com/whs/whscreeps/screeps/game"
	"github.com/whs/whscreeps/screeps/memory"
	"github.com/whs/whscreeps/whscreeps/store"
	"github.com/whs/whscreeps/whscreeps/task"
)

func Loop() {
	defer log.Debug().Msgf("Loop ended. CPU time used: %.4f", game.GetCPUUsed())

	memSegment := memory.GetSegment(0)
	var mem store.RootStore
	err := memory.GetJSON(memSegment, &mem)
	if err != nil {
		log.Error().Err(err).Msg("Error parsing memory")
	}
	mem.Init()
	defer func() {
		err := memory.SetJSON(memSegment, mem)
		if err != nil {
			log.Fatal().Err(err).Msg("Fail to set memory")
		}
	}()

	scheduler := os.GetScheduler(memory.GetSegment(1))

	scheduler.ScheduleTransient(task.Main{}, os.PriorityNormal)

	ctx := store.WithStore(context.Background(), &mem)
	scheduler.Run(ctx)
}
