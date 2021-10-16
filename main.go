package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/whs/whscreeps/whscreeps"
	"os"
	"syscall/js"
)

func loop() {
	whscreeps.Loop()
}

func init() {
	log.Logger = zerolog.New(os.Stderr)
}

func main() {
	js.Global().Set("loop", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		loop()
		return nil
	}))
	log.Info().Msg("Go init completed")
}
