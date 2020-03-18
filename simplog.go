package simplog

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

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

func (s *Simplog) SetLevel(lvl Level) {
	s.level = lvl
}

func (s *Simplog) SetFlags(flag int) {
	s.loggr.SetFlags(flag)
}

func (s *Simplog) write(levelTag, msg string) {
	_ = s.loggr.Output(3, levelTag+" "+msg)
}

func (s *Simplog) Debug(v ...interface{}) {
	if Debug >= s.level {
		s.write(debugTag, fmt.Sprint(v...))
	}
}
func (s *Simplog) Debugln(v ...interface{}) {
	if Debug >= s.level {
		s.write(debugTag, fmt.Sprintln(v...))
	}
}
func (s *Simplog) Debugf(format string, v ...interface{}) {
	if Debug >= s.level {
		s.write(debugTag, fmt.Sprintf(format, v...))
	}
}

func (s *Simplog) Info(v ...interface{}) {
	if Info >= s.level {
		s.write(infoTag, fmt.Sprint(v...))
	}
}
func (s *Simplog) Infoln(v ...interface{}) {
	if Info >= s.level {
		s.write(infoTag, fmt.Sprintln(v...))
	}
}
func (s *Simplog) Infof(format string, v ...interface{}) {
	if Info >= s.level {
		s.write(infoTag, fmt.Sprintf(format, v...))
	}
}

func (s *Simplog) Warning(v ...interface{}) {
	if Warning >= s.level {
		s.write(warningTag, fmt.Sprint(v...))
	}
}
func (s *Simplog) Warningln(v ...interface{}) {
	if Warning >= s.level {
		s.write(warningTag, fmt.Sprintln(v...))
	}
}
func (s *Simplog) Warningf(format string, v ...interface{}) {
	if Warning >= s.level {
		s.write(warningTag, fmt.Sprintf(format, v...))
	}
}

func (s *Simplog) Error(v ...interface{}) {
	if Error >= s.level {
		s.write(errorTag, fmt.Sprint(v...))
	}
}
func (s *Simplog) Errorln(v ...interface{}) {
	if Error >= s.level {
		s.write(errorTag, fmt.Sprintln(v...))
	}
}
func (s *Simplog) Errorf(format string, v ...interface{}) {
	if Error >= s.level {
		s.write(errorTag, fmt.Sprintf(format, v...))
	}
}

func (s *Simplog) Fatal(v ...interface{}) {
	if Fatal >= s.level {
		s.write(fatalTag, fmt.Sprint(v...))
	}
}
func (s *Simplog) Fatalln(v ...interface{}) {
	if Fatal >= s.level {
		s.write(fatalTag, fmt.Sprintln(v...))
	}
}
func (s *Simplog) Fatalf(format string, v ...interface{}) {
	if Fatal >= s.level {
		s.write(fatalTag, fmt.Sprintf(format, v...))
	}
}
