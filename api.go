package log

import (
	"runtime"
	"strconv"
)

func WithData(data Data) *log {
	lg := newLog()
	lg.data = data
	return lg
}

func Error(err error, messages ...string) {
	lg := newLog()
	lg.Error(err, messages...)
}

func Info(messages ...string) {
	lg := newLog()
	lg.Info(messages...)
}

func Callers(depth int) []string {
	callers := make([]string, 0, depth)

	pc := make([]uintptr, depth)
	runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc)
	for {
		frame, more := frames.Next()
		callers = append(callers, frame.File+":"+strconv.Itoa(frame.Line))
		if !more {
			break
		}
	}

	return callers
}
