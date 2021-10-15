package object

import (
	"github.com/whs/whscreeps/screeps/types"
	"syscall/js"
)

type Store struct {
	Raw js.Value
}

func (s Store) JSValue() js.Value {
	return s.Raw
}

func (s Store) GetFreeCapacity() int {
	return s.Raw.Call("getFreeCapacity").Int()
}

func (s Store) GetFreeCapacityOf(resource types.Resource) int {
	return s.Raw.Call("getFreeCapacity", resource).Int()
}
