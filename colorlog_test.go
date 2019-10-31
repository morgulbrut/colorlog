package colorlog

import (
	"testing"
)

func TestLog(t *testing.T) {
	loremIpsum()
	SetTimed(true)
	loremIpsum()
	SetTimeFormat("2006-01-02 15:04:05")
	SetLogLevel(WARNING)
	loremIpsum()
	SetTimeFormat("15:04:05")
	loremIpsum()
	SetTimed(false)
	SetLogLevel(OFF)
	loremIpsum()
}

func loremIpsum() {
	Trace("Trace")
	Debug("Debug")
	Info("Info")
	Warning("Warning")
	Error("Error")
	Fatal("Fatal")
}
