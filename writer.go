package log

import (
	"fmt"
	"os"
	"sync"
	"time"
)

const defaultOutputDir = "logs"

type writer struct {
	m      *sync.Mutex
	level  string
	date   string
	output *os.File
}

func newWriter(level string) *writer {
	_, err := os.Stat(defaultOutputDir)
	if err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
		err = os.Mkdir(defaultOutputDir, 0755)
		if err != nil {
			panic(err)
		}
	}

	wtr := &writer{
		m:     &sync.Mutex{},
		level: level,
	}

	wtr.setOutput(time.Now().Format("2006-01-02"))

	return wtr
}

func (wtr *writer) setOutput(date string) {
	wtr.m.Lock()
	defer wtr.m.Unlock()

	// 拿到锁后再判断一次，避免重复操作
	if date == wtr.date {
		return
	}

	filename := defaultOutputDir + "/" + wtr.level + "_" + date + ".log"
	out, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("fail to create new log file:", err)
		return
	}

	_ = wtr.output.Close()

	wtr.date = date
	wtr.output = out
}

func (wtr *writer) print(lg *log) {
	date := lg.time[:10] // e.g. 2021-11-13
	if date != wtr.date {
		wtr.setOutput(date)
	}

	wtr.m.Lock()
	_, _ = wtr.output.Write(lg.toJSON())
	wtr.m.Unlock()
}
