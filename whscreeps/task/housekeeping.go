package task

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/whs/whscreeps/screeps/game"
	"github.com/whs/whscreeps/screeps/object"
	"github.com/whs/whscreeps/whscreeps/store"
)

type Housekeeping struct {
}

func (h Housekeeping) Name() string {
	return "Housekeeping"
}

func (h Housekeeping) Execute(ctx context.Context) {
	logger := zerolog.Ctx(ctx)
	root := store.CtxStore(ctx)
	root.GC()

	creeps := object.GetCreepNames()
	for _, name := range creeps {
		_, ok := root.Creeps[name]
		if ok {
			continue
		}
		root.Creeps[name] = &store.CreepData{}
		logger.Trace().Msgf("Created store entry for '%s'", name)
	}

	root.LastHousekeeping = game.GetTime()
}
