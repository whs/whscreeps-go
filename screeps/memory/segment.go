package memory

import (
	"strconv"
	"syscall/js"
)

var segments = js.Global().Get("RawMemory").Get("segments")

type memorySegment struct {
	index string
}

func GetSegment(segment int) Memory {
	return &memorySegment{index: strconv.FormatInt(int64(segment), 10)}
}

func (m *memorySegment) Get() string {
	return segments.Get(m.index).String()
}

func (m *memorySegment) Set(val string) {
	segments.Set(m.index, val)
}
