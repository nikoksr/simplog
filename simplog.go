package simplog

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

// Level represents the logging level.
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
	defaultLevel = Fatal
	defaultFlags = log.LstdFlags | log.Lshortfile
)

var writeLock sync.Mutex

// Simplog represents an active logger object.
type Simplog struct {
	name    string
	verbose bool
	level   Level
	loggr   *log.Logger
}

// New creates a new Simplog instance.
//
// If verbose is true, log messages get printed to stdout
func New(name string, verbose bool, out io.Writer) *Simplog {
	// Convention to make name always all uppercase
	name = strings.ToUpper(name) + " "

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
//
// Quiet   = -1
// Debug   =  0
// Info    =  1
// Warning =  2
// Error   =  3
// Fatal   =  4
func (s *Simplog) SetLevel(lvl Level) {
	s.level = lvl
}

// SetFlags sets the flags of the logger.
// See: https://godoc.org/github.com/timehop/golog/log#pkg-constants
func (s *Simplog) SetFlags(flag int) {
	s.loggr.SetFlags(flag)
}

// Write log message to outputs.
func (s *Simplog) write(level Level, levelTag, msg string) {
	if s.level >= level {
		if !strings.HasSuffix(msg, "\n") {
			msg += "\n"
		}
		writeLock.Lock()
		defer writeLock.Unlock()
		_ = s.loggr.Output(3, levelTag+" "+msg)
	}
}

// Debug logs a debug message in style of fmt.Printf. New line will be automatically appended.
func (s *Simplog) Debug(format string, v ...interface{}) {
	s.write(Debug, debugTag, fmt.Sprintf(format, v...))
}

// Info logs a info message in style of fmt.Printf. New line will be automatically appended.
func (s *Simplog) Info(format string, v ...interface{}) {
	s.write(Info, infoTag, fmt.Sprintf(format, v...))
}

// Warning logs a warning message in style of fmt.Printf. New line will be automatically appended.
func (s *Simplog) Warning(format string, v ...interface{}) {
	s.write(Warning, warningTag, fmt.Sprintf(format, v...))
}

// Error logs a error message in style of fmt.Printf. New line will be automatically appended.
func (s *Simplog) Error(format string, v ...interface{}) {
	s.write(Error, errorTag, fmt.Sprintf(format, v...))
}

// Fatal logs a fatal message in style of fmt.Printf. New line will be automatically appended.
func (s *Simplog) Fatal(format string, v ...interface{}) {
	s.write(Fatal, fatalTag, fmt.Sprintf(format, v...))
}
