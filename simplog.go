package simplog

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Level is represents logging levels.
type Level int

const (
	Debug Level = iota
	Info
	Warning
	Error
	Fatal
)

const (
	debugTag   = "DEBUG"
	infoTag    = "INFO"
	warningTag = "WARN"
	errorTag   = "ERROR"
	fatalTag   = "FATAL"
)

const (
	defaultLevel = Info
	defaultFlags = log.LstdFlags | log.Lshortfile
)

type Simplog struct {
	name    string
	verbose bool
	level   Level
	loggr   *log.Logger
}

// New creates a new Simplog instance.
func New(name string, verbose bool, out io.Writer) *Simplog {
	// Convention to make name always all uppercase
	name = strings.ToUpper(name)

	// Create list of outputs. Default will of course be the given one.
	outputs := []io.Writer{out}

	// If verbosity is wanted, stdout will be added as output
	if verbose {
		outputs = append(outputs, os.Stdout)
	}

	l := log.New(io.MultiWriter(outputs...), name, defaultFlags)
	return &Simplog{
		name:    name,
		verbose: verbose,
		level:   defaultLevel,
		loggr:   l,
	}
}

// SetLevel sets the level of the logger.
func (s *Simplog) SetLevel(lvl Level) {
	s.level = lvl
}

// SetFlags sets the flags of the logger.
func (s *Simplog) SetFlags(flag int) {
	s.loggr.SetFlags(flag)
}

func (s *Simplog) write(levelTag, msg string) {
	_ = s.loggr.Output(3, levelTag+" "+msg)
}

// Debug logs a debug message in style of fmt.Print
func (s *Simplog) Debug(v ...interface{}) {
	if Debug >= s.level {
		s.write(debugTag, fmt.Sprint(v...))
	}
}

// Debugln logs a debug message in style of fmt.Println
func (s *Simplog) Debugln(v ...interface{}) {
	if Debug >= s.level {
		s.write(debugTag, fmt.Sprintln(v...))
	}
}

// Debugf logs a debug message in style of fmt.Printf
func (s *Simplog) Debugf(format string, v ...interface{}) {
	if Debug >= s.level {
		s.write(debugTag, fmt.Sprintf(format, v...))
	}
}

// Info logs a info message in style of fmt.Print
func (s *Simplog) Info(v ...interface{}) {
	if Info >= s.level {
		s.write(infoTag, fmt.Sprint(v...))
	}
}

// Infoln logs a info message in style of fmt.Println
func (s *Simplog) Infoln(v ...interface{}) {
	if Info >= s.level {
		s.write(infoTag, fmt.Sprintln(v...))
	}
}

// Infof logs a info message in style of fmt.Printf
func (s *Simplog) Infof(format string, v ...interface{}) {
	if Info >= s.level {
		s.write(infoTag, fmt.Sprintf(format, v...))
	}
}

// Warning logs a warning message in style of fmt.Print
func (s *Simplog) Warning(v ...interface{}) {
	if Warning >= s.level {
		s.write(warningTag, fmt.Sprint(v...))
	}
}

// Warningln logs a warning message in style of fmt.Println
func (s *Simplog) Warningln(v ...interface{}) {
	if Warning >= s.level {
		s.write(warningTag, fmt.Sprintln(v...))
	}
}

// Warningf logs a warning message in style of fmt.Printf
func (s *Simplog) Warningf(format string, v ...interface{}) {
	if Warning >= s.level {
		s.write(warningTag, fmt.Sprintf(format, v...))
	}
}

// Error logs a error message in style of fmt.Print
func (s *Simplog) Error(v ...interface{}) {
	if Error >= s.level {
		s.write(errorTag, fmt.Sprint(v...))
	}
}

// Errorln logs a error message in style of fmt.Println
func (s *Simplog) Errorln(v ...interface{}) {
	if Error >= s.level {
		s.write(errorTag, fmt.Sprintln(v...))
	}
}

// Errorf logs a error message in style of fmt.Printf
func (s *Simplog) Errorf(format string, v ...interface{}) {
	if Error >= s.level {
		s.write(errorTag, fmt.Sprintf(format, v...))
	}
}

// Fatal logs a fatal message in style of fmt.Print
func (s *Simplog) Fatal(v ...interface{}) {
	if Fatal >= s.level {
		s.write(fatalTag, fmt.Sprint(v...))
	}
}

// Fatalln logs a fatal message in style of fmt.Println
func (s *Simplog) Fatalln(v ...interface{}) {
	if Fatal >= s.level {
		s.write(fatalTag, fmt.Sprintln(v...))
	}
}

// Fatalf logs a fatal message in style of fmt.Printf
func (s *Simplog) Fatalf(format string, v ...interface{}) {
	if Fatal >= s.level {
		s.write(fatalTag, fmt.Sprintf(format, v...))
	}
}
