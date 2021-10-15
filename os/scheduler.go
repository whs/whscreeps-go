package os

import (
	"container/heap"
	"context"
	"encoding/gob"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/whs/whscreeps/screeps/memory"
)

const (
	PriorityRealtime   = 0
	PriorityVeryHigh   = 1000
	PriorityHigh       = 5000
	PriorityNormal     = 10_000
	PriorityLow        = 30_000
	PriorityVeryLow    = 40_000
	PriorityBackground = 50_000
)

func init() {
	gob.Register(storedQueue{})
}

type scheduledTask struct {
	Task     Task
	Priority uint16
	Persist  bool
}

func (s *scheduledTask) MarshalZerologObject(e *zerolog.Event) {
	e.Str("name", s.Task.Name())
	e.Uint16("priority", s.Priority)
	e.Bool("persist", s.Persist)
}

type heapScheduledList []scheduledTask

func (h heapScheduledList) Len() int {
	return len(h)
}

func (h heapScheduledList) Less(i, j int) bool {
	return h[i].Priority < h[j].Priority
}

func (h heapScheduledList) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *heapScheduledList) Push(x interface{}) {
	*h = append(*h, x.(scheduledTask))
}

func (h *heapScheduledList) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

type storedQueue struct {
	Queue     *heapScheduledList
	NextQueue *[]scheduledTask
}

type Scheduler struct {
	queue     heapScheduledList
	nextQueue []scheduledTask
	mem       memory.Memory
}

func GetScheduler(mem memory.Memory) *Scheduler {
	out := &Scheduler{
		mem:       mem,
		nextQueue: make([]scheduledTask, 0),
	}
	out.load()
	return out
}

func (s *Scheduler) load() {
	var stored storedQueue
	err := memory.Get(s.mem, &stored)
	if err != nil {
		log.Warn().Err(err).Msg("fail to load scheduled tasks")
		s.queue = make(heapScheduledList, 0)
		return
	}

	count := 0
	if stored.Queue != nil {
		count += len(*stored.Queue)
	}
	if stored.NextQueue != nil {
		count += len(*stored.NextQueue)
	}

	out := make(heapScheduledList, count)
	i := 0
	if stored.Queue != nil {
		for _, item := range *stored.Queue {
			// TODO: I think we could optimize this to not persist entirely
			if item.Persist {
				out[i] = item
				i += 1
			}
		}
	}
	if stored.NextQueue != nil {
		for _, item := range *stored.NextQueue {
			out[i] = item
			i += 1
		}
	}

	s.queue = out[:i]
	heap.Init(&s.queue)
}

// Save save this scheduler state to the original memory location
func (s *Scheduler) Save() {
	s.CopyTo(s.mem)
}

// CopyTo save this scheduler state to the memory location provided
func (s *Scheduler) CopyTo(mem memory.Memory) {
	err := memory.Set(mem, storedQueue{
		Queue:     &s.queue,
		NextQueue: &s.nextQueue,
	})
	if err != nil {
		panic(err)
	}
}

func (s *Scheduler) Run(ctx context.Context) {
	ch := ctx.Done()
	subCtx := context.WithValue(ctx, schedulerContextValue, s)
	for {
		_, stop := <-ch
		if stop {
			break
		}
		if len(s.queue) == 0 {
			break
		}
		task := heap.Pop(&s.queue).(scheduledTask)
		taskLogger := log.With().Object("task", &task).Logger()
		taskLogger.Trace().Msg("Executing task")
		task.Task.Execute(taskLogger.WithContext(subCtx))
		s.Save()
	}
}

func (s *Scheduler) Schedule(task Task, priority uint16) {
	s.nextQueue = append(s.nextQueue, scheduledTask{
		Task:     task,
		Priority: priority,
		Persist:  true,
	})
}

func (s *Scheduler) ScheduleTransient(task Task, priority uint16) {
	heap.Push(&s.queue, scheduledTask{
		Task:     task,
		Priority: priority,
		Persist:  false,
	})
}

func Schedule(ctx context.Context, task Task, priority uint16) {
	scheduler := CtxScheduler(ctx)
	scheduler.Schedule(task, priority)
}
