package object

import (
	"github.com/whs/whscreeps/screeps/types"
	"syscall/js"
)

type Room struct {
	Raw js.Value
}

func (r Room) JSValue() js.Value {
	return r.Raw
}

func (r Room) find(what types.FindType) []RoomObject {
	val := r.Raw.Call("find", what)

	length := val.Length()
	out := make([]RoomObject, val.Length())
	for i := 0; i < length; i++ {
		out[i] = RoomObject{Raw: val.Index(i)}
	}

	return out
}

func (r Room) FindSources(active bool) []RoomObject {
	what := types.FindSources
	if active {
		what = types.FindSourcesActive
	}
	return r.find(what)
}

func (r Room) FindStructures() []Structure {
	return asStructures(r.find(types.FindStructures))
}

func asStructures(val []RoomObject) []Structure {
	out := make([]Structure, len(val))
	for i, v := range val {
		out[i] = Structure{v}
	}
	return out
}
