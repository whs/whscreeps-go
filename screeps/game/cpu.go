package game

import (
	"github.com/teamortix/golang-wasm/wasm"
)

type CPU struct {
	Limit        *int            `wasm:"limit"`
	TickLimit    *int            `wasm:"tickLimit"`
	Bucket       *int            `wasm:"bucket"`
	ShardLimits  *map[string]int `wasm:"shardLimits"`
	Unlocked     *bool           `wasm:"unlocked"`
	UnlockedTime *int            `wasm:"unlockedTime"`
}

func GetCPU() (cpu CPU) {
	err := wasm.FromJSValue(Game.Get("cpu"), &cpu)
	if err != nil {
		panic(err)
	}
	return
}

func GetCPUUsed() float64 {
	return Game.Get("cpu").Call("getUsed").Float()
}
