package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/whs/whscreeps/whscreeps"
	"os"
	"runtime"
	"syscall/js"
)

func main() {
	log.Logger = zerolog.New(os.Stderr)

	js.Global().Set("loop", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		whscreeps.Loop()
		return nil
	}))

	log.Info().Msgf("%s started", runtime.Version())

	select {}
}
