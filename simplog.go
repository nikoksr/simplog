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

func (s *Simplog) write(level Level, levelTag, msg string) {
	if s.level >= level {
		writeLock.Lock()
		defer writeLock.Unlock()
		_ = s.loggr.Output(3, levelTag+" "+msg)
	}
}

// Debug logs a debug message in style of fmt.Print
func (s *Simplog) Debug(v ...interface{}) {
	s.write(Debug, debugTag, fmt.Sprint(v...))
}

// Debugln logs a debug message in style of fmt.Println
func (s *Simplog) Debugln(v ...interface{}) {
	s.write(Debug, debugTag, fmt.Sprintln(v...))
}

// Debugf logs a debug message in style of fmt.Printf
func (s *Simplog) Debugf(format string, v ...interface{}) {
	s.write(Debug, debugTag, fmt.Sprintf(format, v...))
}

// Info logs a info message in style of fmt.Print
func (s *Simplog) Info(v ...interface{}) {
	s.write(Info, infoTag, fmt.Sprint(v...))
}

// Infoln logs a info message in style of fmt.Println
func (s *Simplog) Infoln(v ...interface{}) {
	s.write(Info, infoTag, fmt.Sprintln(v...))
}

// Infof logs a info message in style of fmt.Printf
func (s *Simplog) Infof(format string, v ...interface{}) {
	s.write(Info, infoTag, fmt.Sprintf(format, v...))
}

// Warning logs a warning message in style of fmt.Print
func (s *Simplog) Warning(v ...interface{}) {
	s.write(Warning, warningTag, fmt.Sprint(v...))
}

// Warningln logs a warning message in style of fmt.Println
func (s *Simplog) Warningln(v ...interface{}) {
	s.write(Warning, warningTag, fmt.Sprintln(v...))
}

// Warningf logs a warning message in style of fmt.Printf
func (s *Simplog) Warningf(format string, v ...interface{}) {
	s.write(Warning, warningTag, fmt.Sprintf(format, v...))
}

// Error logs a error message in style of fmt.Print
func (s *Simplog) Error(v ...interface{}) {
	s.write(Error, errorTag, fmt.Sprint(v...))
}

// Errorln logs a error message in style of fmt.Println
func (s *Simplog) Errorln(v ...interface{}) {
	s.write(Error, errorTag, fmt.Sprintln(v...))
}

// Errorf logs a error message in style of fmt.Printf
func (s *Simplog) Errorf(format string, v ...interface{}) {
	s.write(Error, errorTag, fmt.Sprintf(format, v...))
}

// Fatal logs a fatal message in style of fmt.Print
func (s *Simplog) Fatal(v ...interface{}) {
	s.write(Fatal, fatalTag, fmt.Sprint(v...))
}

// Fatalln logs a fatal message in style of fmt.Println
func (s *Simplog) Fatalln(v ...interface{}) {
	s.write(Fatal, fatalTag, fmt.Sprintln(v...))
}

// Fatalf logs a fatal message in style of fmt.Printf
func (s *Simplog) Fatalf(format string, v ...interface{}) {
	s.write(Fatal, fatalTag, fmt.Sprintf(format, v...))
}
