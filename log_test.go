package log

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	Error(errors.New("test error"))

	WithData(Data{"name": "cat"}).Error(errors.New("test error"))
	logCtx := WithData(Data{"name": "dog"})
	logCtx.Error(errors.New("test error"))

	WithData(Data{"name": "cat"}).AddData(Data{"age": 12}).Error(errors.New("test error"))
	logCtx1 := WithData(Data{"name": "dog"})
	logCtx1.AddData(Data{"age": 12})
	logCtx1.Error(errors.New("test error"))
}

func TestInfo(t *testing.T) {
	Info("test info")

	WithData(Data{"name": "cat"}).Info("test info")
	logCtx := WithData(Data{"name": "dog"})
	logCtx.Info("test info")

	WithData(Data{"name": "cat"}).AddData(Data{"age": 12}).Info("test info")
	logCtx1 := WithData(Data{"name": "dog"})
	logCtx1.AddData(Data{"age": 12})
	logCtx1.Info("test info")
}

func TestCallers(t *testing.T) {
	fmt.Println(Callers(3))
	WithData(Data{"callers": Callers(1)}).Info("test")
}

func BenchmarkError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Error(errors.New("test error"))
	}
}

func BenchmarkInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Info("test info")
	}
}

func BenchmarkCallers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Callers(1)
	}
}

func BenchmarkInfoWithCallers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WithData(Data{"caller": Callers(1)}).Info("test info")
	}
}

func BenchmarkJSON(b *testing.B) {
	lg := &log{
		error:   "test error",
		message: "test info",
		callers: []string{"/Users/renjiangdu/Projects/micro-logger/log_test.go:53"},
		time:    "2022-04-19 23:31:01",
		data:    Data{"caller": "/Users/renjiangdu/Projects/micro-logger/log_test.go:53"},
	}
	for i := 0; i < b.N; i++ {
		lg.toJSON()
	}
}

func BenchmarkLibJSON(b *testing.B) {
	lg := struct {
		Error   string   `json:"error"`
		Message string   `json:"message"`
		Callers []string `json:"callers"`
		Time    string   `json:"time"`
		Data    Data     `json:"data"`
	}{
		Error:   "test error",
		Message: "test info",
		Callers: []string{"/Users/renjiangdu/Projects/micro-logger/log_test.go:53"},
		Time:    "2022-04-19 23:31:01",
		Data:    Data{"caller": "/Users/renjiangdu/Projects/micro-logger/log_test.go:53"},
	}
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(lg)
	}
}
