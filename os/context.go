package os

import (
	"context"
	"github.com/whs/whscreeps/screeps/game"
	"time"
)

var DefaultCpuLimit = 20

// DeadlineLeeway deducts some amount of time from the deadline for runtime
var DeadlineLeeway = 500 * time.Microsecond

// GetDeadline returns the current tick execution deadline in context.Context
func GetDeadline(ctx context.Context) (context.Context, context.CancelFunc) {
	cpu := game.GetCPU()
	runtime := DefaultCpuLimit
	if cpu.Limit != nil { // somehow the runtime doesn't play well with null value
		runtime = *cpu.TickLimit
	}

	//usedCpu := 0
	usedCpu := game.GetCPUUsed()

	deadline := time.Duration(runtime)*time.Millisecond - time.Duration(usedCpu)*time.Millisecond - DeadlineLeeway
	return context.WithTimeout(ctx, deadline)
}

type schedulerContextKey int

var schedulerContextValue schedulerContextKey

func CtxScheduler(ctx context.Context) *Scheduler {
	return ctx.Value(schedulerContextValue).(*Scheduler)
}
