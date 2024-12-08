/*
Реализовать паттерн «адаптер» на любом примере.
*/

package main

import (
	"fmt"
	"io"
	"os"
)

const ()

type ExternalLogger struct {
	w io.Writer
}

func (el *ExternalLogger) Log(level string, msg string) {
	if level != "debug" && level != "warn" && level != "error" {
		level = "info"
	}
	message := fmt.Sprintf("[%s]\t- %s\n", level, msg)
	if level == "error" {
		message = "\033[31m" + message + "\033[0m"
	}
	el.w.Write([]byte(message))
}

func (el *ExternalLogger) Close() {
	el.w = nil
}

type Logger interface {
	Debug(msg ...any)
	Info(msg ...any)
	Warn(msg ...any)
	Error(msg ...any)
	io.Closer
}

type OurLoggerAdapter struct {
	*ExternalLogger
}

func NewOurLoggerAdapter(w io.Writer) *OurLoggerAdapter {
	return &OurLoggerAdapter{
		ExternalLogger: &ExternalLogger{w: w},
	}
}

func (l *OurLoggerAdapter) Debug(msg ...any) {
	l.Log("debug", fmt.Sprintf("%v", msg))
}

func (l *OurLoggerAdapter) Info(msg ...any) {
	l.Log("info", fmt.Sprintf("%v", msg))
}

func (l *OurLoggerAdapter) Warn(msg ...any) {
	l.Log("warn", fmt.Sprintf("%v", msg))
}

func (l *OurLoggerAdapter) Error(msg ...any) {
	l.Log("error", fmt.Sprintf("%v", msg))
}

func (l *OurLoggerAdapter) Close() error {
	l.ExternalLogger.Close()
	return nil
}

func main() {
	logger := NewOurLoggerAdapter(os.Stdout)
	defer logger.Close()
	SomeHandler(logger)
}

func SomeHandler(logger Logger) {
	logger.Debug("Debug message", 42, 1337, true)
	logger.Info("Info message", 1e5, 2.5, "hello")
	logger.Warn("Warn message", 1_000_000_000, "world")
	logger.Error("Error message", 3.14, "pi", '!')
}
