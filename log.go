package log

import (
	"runtime"
	"strconv"
	"strings"
	"time"
)

const defaultCallerDepth = 4

var (
	errorWriter = newWriter("error")
	infoWriter  = newWriter("info")
)

type Data map[string]interface{}

type log struct {
	error   string
	message string
	callers []string
	time    string
	data    Data
}

func newLog() *log {
	return &log{time: time.Now().Format("2006-01-02 15:04:05")}
}

func (lg *log) AddData(addition Data) *log {
	data := lg.data
	for key, value := range addition {
		data[key] = value
	}
	lg.data = data
	return lg
}

func (lg *log) Error(err error, messages ...string) {
	lg.error = err.Error()
	lg.message = strings.Join(messages, " ")
	lg.callers = make([]string, 0, defaultCallerDepth)

	skip := 3
	if len(lg.data) > 0 {
		skip--
	}
	pc := make([]uintptr, defaultCallerDepth)
	runtime.Callers(skip, pc)
	frames := runtime.CallersFrames(pc)
	for {
		frame, more := frames.Next()
		lg.callers = append(lg.callers, frame.File+":"+strconv.Itoa(frame.Line))
		if !more {
			break
		}
	}

	errorWriter.print(lg)
}

func (lg *log) Info(messages ...string) {
	lg.message = strings.Join(messages, " ")
	infoWriter.print(lg)
}
