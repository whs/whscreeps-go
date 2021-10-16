package task

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/whs/whscreeps/os"
	"github.com/whs/whscreeps/screeps/object"
	"github.com/whs/whscreeps/screeps/types"
	"github.com/whs/whscreeps/whscreeps/store"
)

type Harvest struct {
	CreepID       string
	HarvestTarget string
	ReturnTarget  string

	Token   store.LockToken
	creep   object.Creep
	logger  zerolog.Logger
	context context.Context
	store   *store.RootStore
}

func (h *Harvest) Name() string {
	return fmt.Sprintf("Harvest(%s)", h.CreepID)
}

func (h *Harvest) Execute(ctx context.Context) {
	h.context = ctx
	h.logger = zerolog.Ctx(ctx).With().
		Str("harvestTarget", h.HarvestTarget).
		Str("returnTarget", h.ReturnTarget).
		Logger()
	h.store = store.CtxStore(ctx)

	if !h.store.Creeps[h.CreepID].IsMyLock(h.Token) {
		h.logger.Warn().Msg("Someone stole my lock!")
		return
	}

	creep, err := object.GetCreep(h.CreepID)
	if err == object.ErrNoCreep {
		h.logger.Debug().Msg("Creep died")
		return
	} else if err != nil {
		h.logger.Err(err).Send()
		return
	}

	h.creep = creep

	if creep.Spawning() {
		h.logger.Trace().Msg("Creep still spawning")
		os.Schedule(ctx, h, os.PriorityVeryLow)
		return
	}

	store := creep.Store()

	if store.GetFreeCapacity() > 0 {
		if h.ReturnTarget != "" {
			h.logger.Trace().Msg("Objective complete!")
			h.release()
			return
		}

		h.harvest()
	} else {
		h.recall()
	}
}

func (h *Harvest) release() {
	h.store.Creeps[h.CreepID].Unlock(h.Token)
}

func (h *Harvest) harvest() {
	var target object.RoomObject
	if h.HarvestTarget == "" {
		sources := h.creep.Room().FindSources(false)

		if len(sources) == 0 {
			log.Info().Msg("Releasing - no harvest target found")
			h.release()
			return
		}

		target = sources[0]
		h.HarvestTarget = target.ID()
	} else {
		target = *object.RoomObjectById(h.HarvestTarget)
	}

	err := h.creep.Harvest(target)
	if err == types.ErrNotInRange {
		h.logger.Trace().Msg("Moving to harvesting target")
		h.creep.MoveToTarget(target, nil)
		os.Schedule(h.context, h, os.PriorityVeryLow)
	} else if err != nil {
		h.logger.Warn().Err(err).Msg("Harvest fail")
		h.HarvestTarget = ""
		os.Schedule(h.context, h, os.PriorityVeryLow)
	} else {
		h.logger.Trace().Msg("Harvesting")
		os.Schedule(h.context, h, os.PriorityLow)
	}
}

func (h *Harvest) recall() {
	var target object.RoomObject
	if h.ReturnTarget == "" {
		targets := h.creep.Room().FindStructures()
		for _, item := range targets {
			structureType := item.Type()
			if structureType != types.StructureExtension && structureType != types.StructureSpawn {
				continue
			}
			if item.Store().GetFreeCapacityOf(types.ResourceEnergy) == 0 {
				continue
			}
			target = item.RoomObject
			h.ReturnTarget = target.ID()
			break
		}
		if h.ReturnTarget == "" {
			// TODO: Move away from the source a bit
			log.Info().Msg("Releasing - no return target found")
			h.release()
			return
		}
	} else {
		target = *object.RoomObjectById(h.ReturnTarget)
	}

	err := h.creep.Transfer(target, types.ResourceEnergy)
	if err == types.ErrNotInRange {
		h.logger.Trace().Msg("Moving to transfer target")
		h.creep.MoveToTarget(target, nil)
		os.Schedule(h.context, h, os.PriorityVeryLow)
	} else if err != nil {
		h.logger.Warn().Err(err).Msg("Fail to transfer")
		h.ReturnTarget = ""
		os.Schedule(h.context, h, os.PriorityVeryLow)
	} else {
		h.logger.Trace().Msg("Transferring")
		os.Schedule(h.context, h, os.PriorityNormal)
	}
}
