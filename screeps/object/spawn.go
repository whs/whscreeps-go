package object

import (
	"github.com/whs/whscreeps/screeps/types"
	"github.com/whs/whscreeps/screeps/utils"
	"github.com/whs/whscreeps/wasm"
	"syscall/js"
)

type Spawn struct {
	Structure
}

func (s *Spawn) Name() string {
	return s.Raw.Get("name").String()
}

func (s *Spawn) Spawning() *Spawning {
	val := s.Raw.Get("spawning")
	if val.IsNull() {
		return nil
	}
	return &Spawning{val}
}

func (s *Spawn) Store() {
	panic("fixme")
}

func (s *Spawn) SpawnCreep(parts []types.BodyPart, name string, options SpawnOption) error {
	return utils.R(s.Raw.Call("spawnCreep", wasm.ToJSValue(parts), name, wasm.ToJSValue(options)))
}

type SpawnOption map[string]interface{}

func (s SpawnOption) EnergyStructures(structures []*Structure) SpawnOption {
	s["energyStructures"] = structures
	return s
}

func (s SpawnOption) DryRun(v bool) SpawnOption {
	s["dryRun"] = v
	return s
}

func (s SpawnOption) Direction(v []types.Direction) SpawnOption {
	s["directions"] = v
	return s
}

func (s *Spawn) RecycleCreep(creep *Creep) error {
	return utils.R(s.Raw.Call("recycleCreep", creep))
}

func (s *Spawn) RenewCreep(creep *Creep) error {
	return utils.R(s.Raw.Call("renewCreep", creep))
}

type Spawning struct {
	Raw js.Value
}

func (s Spawning) JSValue() js.Value {
	return s.Raw
}

// Name of the new creep
func (s Spawning) Name() string {
	return s.Raw.Get("name").String()
}

func (s Spawning) NeedTime() int {
	return s.Raw.Get("needTime").Int()
}

func (s Spawning) RemainingTime() int {
	return s.Raw.Get("remainingTime").Int()
}

func (s Spawning) Spawn() Spawn {
	return Spawn{Structure{RoomObject{Raw: s.Raw.Get("spawn")}}}
}

// Cancel spawning immediately. Energy spent on spawning is not returned.
func (s Spawning) Cancel() error {
	return utils.R(s.Raw.Call("cancel"))
}
