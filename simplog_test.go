package simplog

import (
	"bytes"
	"io"
	"log"
	"os"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		name    string
		verbose bool
	}

	tests := []struct {
		name string
		args args
		want *Simplog
	}{
		{
			"CORE",
			args{
				"CORE",
				false,
			},
			&Simplog{
				name:    "CORE ",
				verbose: false,
				level:   defaultLevel,
				loggr: log.New(
					io.MultiWriter([]io.Writer{&bytes.Buffer{}}...),
					"CORE ",
					defaultFlags,
				),
			},
		},
		{
			"APP",
			args{
				"app",
				true,
			},
			&Simplog{
				name:    "APP ",
				verbose: true,
				level:   defaultLevel,
				loggr: log.New(
					io.MultiWriter([]io.Writer{&bytes.Buffer{}, os.Stdout}...),
					"APP ",
					defaultFlags,
				),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.name, tt.args.verbose, &bytes.Buffer{})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
