package utils

import (
	"github.com/whs/whscreeps/screeps/types"
	"syscall/js"
)

func R(resp js.Value) error {
	out := types.Error(resp.Int())
	if out == types.Ok {
		return nil
	}
	return out
}
