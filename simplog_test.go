package simplog

import (
	"bytes"
	"context"
	"net/url"
	"strings"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestNewLogger(t *testing.T) {
	t.Parallel()

	type args struct {
		options *Options
	}
	cases := []struct {
		name string
		args args
	}{
		{
			name: "nil options",
			args: args{
				options: nil,
			},
		},
		{
			name: "empty options",
			args: args{
				options: &Options{},
			},
		},
		{
			name: "with options; client in debug mode",
			args: args{
				options: &Options{
					Debug:    true,
					IsServer: false,
				},
			},
		},
		{
			name: "with options; client in production mode",
			args: args{
				options: &Options{
					Debug:    false,
					IsServer: false,
				},
			},
		},
		{
			name: "with options; server in debug mode",
			args: args{
				options: &Options{
					Debug:    true,
					IsServer: true,
				},
			},
		},
		{
			name: "with options; server in production mode",
			args: args{
				options: &Options{
					Debug:    false,
					IsServer: true,
				},
			},
		},
		{
			name: "with options; client with stacktrace",
			args: args{
				options: &Options{
					IsServer:          false,
					DisableStacktrace: false,
				},
			},
		},
		{
			name: "with options; client without stacktrace",
			args: args{
				options: &Options{
					IsServer:          false,
					DisableStacktrace: true,
				},
			},
		},
	}

	for _, tc := range cases {
		var logger *zap.SugaredLogger

		// Doing this to cover all the different constructors.
		if tc.args.options == nil {
			logger = NewWithOptions(tc.args.options) // Can handle nil options.
			tc.args.options = defaultOptions         // Doing this to avoid nil pointer dereference in the comparisons below.
		} else {
			if !tc.args.options.IsServer {
				logger = NewClientLogger(tc.args.options.Debug)
			} else {
				logger = NewServerLogger(tc.args.options.Debug)
			}
		}

		if logger == nil {
			t.Fatal("Returned a nil logger")
		}

		if tc.args.options.Debug {
			if !logger.Desugar().Core().Enabled(zap.DebugLevel) {
				t.Error("Debug logger not enabled at debug level")
			}
		} else {
			if !logger.Desugar().Core().Enabled(zap.InfoLevel) {
				t.Error("Production logger not enabled at info level")
			}
		}
	}
}

func TestFromContext(t *testing.T) {
	t.Parallel()

	// Positive case; context has a logger.
	l1 := NewWithOptions(nil)
	if l1 == nil {
		t.Fatal("NewClientLogger returned a nil logger")
	}

	ctx := WithLogger(context.Background(), l1)
	if ctx == nil {
		t.Fatal("WithLogger returned a nil context")
	}

	l2 := FromContext(ctx)
	if l2 == nil {
		t.Fatal("FromContext returned a nil logger")
	}

	// Comparing memory addresses should be sufficient to determine equality.
	if l1 != l2 {
		t.Error("FromContext returned a different logger")
	}

	// Context has no logger, returned logger should still not be nil.
	ctx = context.Background()
	l3 := FromContext(ctx)
	if l3 == nil {
		t.Fatal("FromContext returned a nil logger")
	}
}

// This is for testing purposes only. Copied from:
//   - https://github.com/uber-go/zap/blob/v1.22.0/sink.go#L61
type nopCloserSink struct{ zapcore.WriteSyncer }

func (nopCloserSink) Close() error { return nil }

func TestVisualLevelEncoder(t *testing.T) {
	t.Parallel()

	// Custom memory sink to capture output.
	buf := bytes.NewBuffer(nil)
	memFactory := func(u *url.URL) (zap.Sink, error) {
		return nopCloserSink{zapcore.AddSync(buf)}, nil
	}

	err := zap.RegisterSink("mem", memFactory)
	if err != nil {
		t.Fatalf("register new zap sink: %v", zap.Error(err))
	}

	// Create a logger with the custom sink and the visualLevelEncoder that we want to test.
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = visualLevelEncoder
	config.OutputPaths = []string{"mem://"}

	logger, err := config.Build()
	if err != nil {
		t.Fatalf("build logger: %v", err)
	}

	// Log a message at each level and check the buffer for the expected output.
	// Debug
	logger.Debug("test")
	if !strings.Contains(buf.String(), activeSymbols.fromZapLevel[zap.DebugLevel]) {
		t.Fatal("expected debug symbol in log message")
	}
	buf.Reset()

	// Info
	logger.Info("test")
	if !strings.Contains(buf.String(), activeSymbols.fromZapLevel[zap.InfoLevel]) {
		t.Fatal("expected info symbol in log message")
	}
	buf.Reset()

	// Warning
	logger.Warn("test")
	if !strings.Contains(buf.String(), activeSymbols.fromZapLevel[zap.WarnLevel]) {
		t.Fatal("expected warn symbol in log message")
	}
	buf.Reset()

	// Error
	logger.Error("test")
	if !strings.Contains(buf.String(), activeSymbols.fromZapLevel[zap.ErrorLevel]) {
		t.Fatal("expected error symbol in log message")
	}
	buf.Reset()

	// Now repeat the process but with custom symbols.

	// Debug
	debugSymbol := "<!!D!!>"
	SetDebugSymbol(debugSymbol)
	logger.Debug("test")
	if !strings.Contains(buf.String(), activeSymbols.fromZapLevel[zap.DebugLevel]) {
		t.Fatalf("expected debug symbol %q in log message: %s", debugSymbol, buf.String())
	}
	buf.Reset()

	// Info
	infoSymbol := "<!!I!!>"
	SetInfoSymbol(infoSymbol)
	logger.Info("test")
	if !strings.Contains(buf.String(), activeSymbols.fromZapLevel[zap.InfoLevel]) {
		t.Fatalf("expected info symbol %q in log message: %s", infoSymbol, buf.String())
	}
	buf.Reset()

	// Warning
	warnSymbol := "<!!W!!>"
	SetWarnSymbol(warnSymbol)
	logger.Warn("test")
	if !strings.Contains(buf.String(), activeSymbols.fromZapLevel[zap.WarnLevel]) {
		t.Fatalf("expected warn symbol %q in log message: %s", warnSymbol, buf.String())
	}
	buf.Reset()

	// Error
	errorSymbol := "<!!E!!>"
	SetErrorSymbol(errorSymbol)
	logger.Error("test")
	if !strings.Contains(buf.String(), activeSymbols.fromZapLevel[zap.ErrorLevel]) {
		t.Fatalf("expected error symbol %q in log message: %s", errorSymbol, buf.String())
	}
	buf.Reset()
}
