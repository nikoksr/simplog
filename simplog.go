package simplog

import (
	"context"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	// Options are the options for the logger.
	Options struct {
		// Debug enables debug logging.
		Debug bool

		// IsServer indicates whether the logger is used by a server or a client.
		IsServer bool

		// DisableStacktrace disables the stacktrace.
		DisableStacktrace bool
	}

	symbols = struct {
		fromZapLevel map[zapcore.Level]string
		mu           sync.Mutex
	}
)

const (
	defaultDebugSymbol   = "🐞"
	defaultInfoSymbol    = "💡"
	defaultWarningSymbol = "⚠️ "
	defaultErrorSymbol   = "🔥"
	defaultFatalSymbol   = "💀"
	defaultPanicSymbol   = "🚨"
	defaultDPanicSymbol  = "🚨"
)

var (
	defaultOptions = &Options{
		Debug:             false,
		IsServer:          false,
		DisableStacktrace: true,
	}

	// activeSymbols is used to store the active symbols. It is used to allow for changing the symbols at runtime. It
	// is used by the visualLevelEncoder to encode the level to a human-readable prefix.
	activeSymbols = symbols{
		fromZapLevel: map[zapcore.Level]string{
			zapcore.DebugLevel:  defaultDebugSymbol,
			zapcore.InfoLevel:   defaultInfoSymbol,
			zapcore.WarnLevel:   defaultWarningSymbol,
			zapcore.ErrorLevel:  defaultErrorSymbol,
			zapcore.FatalLevel:  defaultFatalSymbol,
			zapcore.PanicLevel:  defaultPanicSymbol,
			zapcore.DPanicLevel: defaultDPanicSymbol,
		},
	}
)

// visualLevelEncoder is a zapcore.Encoder that encodes a zapcore.Level to a human-readable prefix.
func visualLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(activeSymbols.fromZapLevel[level])
}

func productionCLIEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		FunctionKey:      zapcore.OmitKey,
		LevelKey:         "L",
		MessageKey:       "M",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      visualLevelEncoder,
		ConsoleSeparator: " ",
	}
}

func newLogger(name string, opts *Options) *zap.SugaredLogger {
	if opts == nil {
		opts = defaultOptions
	}

	var config zap.Config

	// In debug mode, client and server applications use the same configuration.
	if opts.Debug {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()

		// Special case for client production logging. Servers should use the default production config - json encoding
		if !opts.IsServer {
			config.Encoding = "console"
			config.EncoderConfig = productionCLIEncoderConfig()
		}
	}

	// Optionally disable stacktrace.
	if opts.DisableStacktrace {
		config.DisableStacktrace = true
	}

	logger, err := config.Build()
	if err != nil {
		logger = zap.NewNop()
	}

	// Add the name to the logger.
	logger = logger.Named(name)

	return logger.Sugar()
}

func defaultLogger() *zap.SugaredLogger {
	return newLogger("simplog-default", defaultOptions)
}

// NewWithOptions returns a new logger with the given options. If the options are nil, the default logger is returned.
func NewWithOptions(opts *Options) *zap.SugaredLogger {
	return newLogger("simplog", opts)
}

// NewClientLogger returns a new logger that's meant to be used by client-type applications. It uses a human-readable
// format.
func NewClientLogger(debug bool) *zap.SugaredLogger {
	return NewWithOptions(&Options{
		Debug:             debug,
		IsServer:          false,
		DisableStacktrace: true,
	})
}

// NewServerLogger returns a new logger that's meant to be used by servers. It uses structured logging when run in
// production, and a human-readable format when run in development.
func NewServerLogger(debug bool) *zap.SugaredLogger {
	return NewWithOptions(&Options{
		Debug:             debug,
		IsServer:          true,
		DisableStacktrace: true,
	})
}

// As recommended by 'revive' linter.
type contextKey string

var loggerKey contextKey = "simplog"

// WithLogger returns a new context.Context with the given logger.
func WithLogger(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

// FromContext returns the logger from the given context. If the context does not contain a logger, the default logger
// is returned. If the context is nil, the default logger is returned.
func FromContext(ctx context.Context) *zap.SugaredLogger {
	if logger, ok := ctx.Value(loggerKey).(*zap.SugaredLogger); ok {
		return logger
	}

	return defaultLogger()
}

func setSymbol(level zapcore.Level, symbol string) {
	activeSymbols.mu.Lock()
	activeSymbols.fromZapLevel[level] = symbol
	activeSymbols.mu.Unlock()
}

// SetDebugSymbol sets the debug symbol.
func SetDebugSymbol(symbol string) {
	setSymbol(zapcore.DebugLevel, symbol)
}

// SetInfoSymbol sets the info symbol.
func SetInfoSymbol(symbol string) {
	setSymbol(zapcore.InfoLevel, symbol)
}

// SetWarnSymbol sets the warning symbol.
func SetWarnSymbol(symbol string) {
	setSymbol(zapcore.WarnLevel, symbol)
}

// SetErrorSymbol sets the error symbol.
func SetErrorSymbol(symbol string) {
	setSymbol(zapcore.ErrorLevel, symbol)
}

// SetFatalSymbol sets the fatal symbol.
func SetFatalSymbol(symbol string) {
	setSymbol(zapcore.FatalLevel, symbol)
}

// SetPanicSymbol sets the panic symbol.
func SetPanicSymbol(symbol string) {
	setSymbol(zapcore.PanicLevel, symbol)
}

// SetDPanicSymbol sets the dpanic symbol.
func SetDPanicSymbol(symbol string) {
	setSymbol(zapcore.DPanicLevel, symbol)
}
