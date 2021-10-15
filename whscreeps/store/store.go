package store

import (
	"context"
	"encoding/gob"
	"github.com/rs/zerolog/log"
	"github.com/whs/whscreeps/screeps/object"
	"github.com/whs/whscreeps/whscreeps/utils"
)

type RootStore struct {
	Creeps           map[string]*CreepData
	LastHousekeeping int
}

func (r *RootStore) GC() {
	r.gcCreeps()
}

func (r *RootStore) Init() {
	if r.Creeps == nil {
		r.Creeps = make(map[string]*CreepData)
	}
}

func (r *RootStore) gcCreeps() {
	// remove missing creep from creepdata
	creeps := object.GetCreepNames()
	for name, _ := range r.Creeps {
		if !utils.StrContains(creeps, name) {
			log.Trace().Str("name", name).Msg("Deleted stale creep key")
			delete(r.Creeps, name)
		}
	}
}

type CreepData struct {
	Mutex
}

func init() {
	gob.Register(RootStore{})
}

type storeContextKey int

var storeContextValue storeContextKey

func CtxStore(ctx context.Context) *RootStore {
	return ctx.Value(storeContextValue).(*RootStore)
}

func WithStore(ctx context.Context, store *RootStore) context.Context {
	return context.WithValue(ctx, storeContextValue, store)
}
