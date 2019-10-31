package colorlog

import (
	"fmt"
	"time"

	color "github.com/morgulbrut/color256"
)

// Level (int) for Typechecking
type Level int

// Loglevels
const (
	TRACE Level = iota // Only makes sense for loglevel
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
	OFF
)

var logLevel Level
var timed bool
var timeFormat = time.RFC3339

// SetLogLevel sets the loglevel
func SetLogLevel(l Level) {
	logLevel = l
}

// SetTimeFormat sets the time format
// Check the time documentation on how to format
func SetTimeFormat(f string) {
	timeFormat = f
}

// SetTimed sets if a timestamp is showed
func SetTimed(b bool) {
	timed = b
}

// log prints a log message without a timestamp
func log(l Level, format string, a ...interface{}) {
	if l >= logLevel {
		switch l {
		case TRACE:
			fmt.Println(color.Magenta(color.Bold("Trace   : ")) + fmt.Sprintf(format, a...))
		case DEBUG:
			fmt.Println(color.Cyan(color.Bold("Debug   : ")) + fmt.Sprintf(format, a...))
		case INFO:
			fmt.Println(color.Green(color.Bold("Info    : ")) + fmt.Sprintf(format, a...))
		case WARNING:
			fmt.Println(color.Yellow(color.Bold("Warning : ")) + fmt.Sprintf(format, a...))
		case ERROR:
			fmt.Println(color.Orange(color.Bold("Error   : ")) + fmt.Sprintf(format, a...))
		case FATAL:
			fmt.Println(color.HiRed(color.Bold("Fatal   :")) + " " + fmt.Sprintf(format, a...))
		}
	}
}

// timeLog prints a log message with a timestamp
func timeLog(l Level, format string, a ...interface{}) {
	t := time.Now().Format(timeFormat)
	if l >= logLevel {
		switch l {
		case TRACE:
			fmt.Println(color.Magenta(color.Bold("Trace   : %s : ", t)) + fmt.Sprintf(format, a...))
		case DEBUG:
			fmt.Println(color.Cyan(color.Bold("Debug   : %s : ", t)) + fmt.Sprintf(format, a...))
		case INFO:
			fmt.Println(color.Green(color.Bold("Info    : %s : ", t)) + fmt.Sprintf(format, a...))
		case WARNING:
			fmt.Println(color.Yellow(color.Bold("Warning : %s : ", t)) + fmt.Sprintf(format, a...))
		case ERROR:
			fmt.Println(color.Orange(color.Bold("Error   : %s : ", t)) + fmt.Sprintf(format, a...))
		case FATAL:
			fmt.Println(color.HiRed(color.Bold("Fatal   : %s :", t)) + " " + fmt.Sprintf(format, a...))
		}
	}
}

// Trace shows a log message on level TRACE
func Trace(format string, a ...interface{}) {
	if timed {
		timeLog(TRACE, format, a...)
	} else {
		log(TRACE, format, a...)
	}
}

// Debug shows a log message on level DEBUG
func Debug(format string, a ...interface{}) {
	if timed {
		timeLog(DEBUG, format, a...)
	} else {
		log(DEBUG, format, a...)
	}
}

// Info shows a log message on level INFO
func Info(format string, a ...interface{}) {
	if timed {
		timeLog(INFO, format, a...)
	} else {
		log(INFO, format, a...)
	}
}

// Warning shows a log message on level WARNING
func Warning(format string, a ...interface{}) {
	if timed {
		timeLog(WARNING, format, a...)
	} else {
		log(WARNING, format, a...)
	}
}

// Error shows a log message on level ERROR
func Error(format string, a ...interface{}) {
	if timed {
		timeLog(ERROR, format, a...)
	} else {
		log(ERROR, format, a...)
	}
}

// Fatal shows a log message on level FATAL
func Fatal(format string, a ...interface{}) {
	if timed {
		timeLog(FATAL, format, a...)
	} else {
		log(FATAL, format, a...)
	}
}
