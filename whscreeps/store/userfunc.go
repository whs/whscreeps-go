package store

import (
	"github.com/rs/zerolog/log"
	"github.com/whs/whscreeps/screeps/memory"
	"syscall/js"
)

func init() {
	js.Global().Set("resetCreepState", js.FuncOf(resetCreepState))
}

func resetCreepState(this js.Value, args []js.Value) interface{} {
	memSegment := memory.GetSegment(0)
	var mem RootStore
	err := memory.GetJSON(memSegment, &mem)
	if err != nil {
		log.Error().Err(err).Msg("Error parsing memory")
		return nil
	}

	mem.Creeps[args[0].String()] = &CreepData{}

	return "Completed!"
}
