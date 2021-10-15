package game

import (
	"github.com/teamortix/golang-wasm/wasm"
)

type GCL struct {
	Level         int `wasm:"level"`
	Progress      int `wasm:"progress"`
	ProgressTotal int `wasm:"progressTotal"`
}

func GetGCL() (gcl GCL) {
	err := wasm.FromJSValue(Game.Get("gcl"), &gcl)
	if err != nil {
		panic(err)
	}
	return
}
