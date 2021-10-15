package memory

import (
	"syscall/js"
)

var RawMemory Memory = &rawMemory{raw: js.Global().Get("RawMemory")}

type rawMemory struct {
	raw js.Value
}

func (r *rawMemory) Get() string {
	return r.raw.Call("get").String()
}

func (r *rawMemory) Set(val string) {
	r.raw.Call("set", val)
}
