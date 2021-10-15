package whscreeps

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/whs/whscreeps/os"
	"github.com/whs/whscreeps/screeps/memory"
	"github.com/whs/whscreeps/whscreeps/store"
	"github.com/whs/whscreeps/whscreeps/task"
)

func Loop() {
	defer log.Debug().Msg("Loop ended")

	ctx, done := os.GetDeadline(context.Background())
	defer done()

	memSegment := memory.GetSegment(0)
	var mem store.RootStore
	err := memory.Get(memSegment, &mem)
	if err != nil {
		log.Error().Err(err).Msg("Error parsing memory")
	}
	mem.Init()
	defer func() {
		err := memory.Set(memSegment, mem)
		if err != nil {
			log.Fatal().Err(err).Msg("Fail to set memory")
		}
	}()

	scheduler := os.GetScheduler(memory.GetSegment(1))

	scheduler.ScheduleTransient(task.Main{}, os.PriorityNormal)

	ctx = store.WithStore(ctx, &mem)
	scheduler.Run(ctx)
}
