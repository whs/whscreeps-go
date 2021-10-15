package object

import (
	"github.com/teamortix/golang-wasm/wasm"
	"github.com/whs/whscreeps/screeps/types"
	"github.com/whs/whscreeps/screeps/utils"
	"syscall/js"
)

type Creep struct {
	RoomObject
}

func (c *Creep) Spawning() bool {
	return c.Raw.Get("spawning").Bool()
}

func (c *Creep) Room() Room {
	return Room{Raw: c.Raw.Get("room")}
}

func (c *Creep) Store() Store {
	return Store{Raw: c.Raw.Get("store")}
}

func (c *Creep) Say(message string, public bool) error {
	return utils.R(c.Raw.Call("say", message, public))
}

func (c *Creep) Harvest(target js.Wrapper) error {
	return utils.R(c.Raw.Call("harvest", target))
}

func (c *Creep) MoveToCoord(x int, y int, options map[string]interface{}) error {
	return utils.R(c.Raw.Call("moveTo", x, y, wasm.ToJSValue(options)))
}

func (c *Creep) MoveToTarget(target js.Wrapper, options map[string]interface{}) error {
	return utils.R(c.Raw.Call("moveTo", target, wasm.ToJSValue(options)))
}

func (c *Creep) Transfer(target js.Wrapper, resourceType types.Resource) error {
	return utils.R(c.Raw.Call("transfer", target, resourceType))
}

func (c *Creep) TransferAmount(target js.Wrapper, resourceType types.Resource, amount int) error {
	return utils.R(c.Raw.Call("transfer", target, resourceType, amount))
}
