package main

import (
	"context"

	"github.com/nikoksr/simplog"
)

func main() {
	// Create a new logger instance. Disabling the stacktrace may come in handy if you want to reduce the log output or
	// want another library to handle the stacktrace, e.g. cockroachdb/errors.
	logger := simplog.NewWithOptions(&simplog.Options{
		Debug:             false,
		IsServer:          false,
		DisableStacktrace: true,
	})

	// Set custom log level symbols
	simplog.SetDebugSymbol("[DEBUG]")
	simplog.SetInfoSymbol("[INFO]")
	simplog.SetWarnSymbol("[WARN]")
	simplog.SetErrorSymbol("[ERROR]")

	// Bind the logger to a context
	ctx := simplog.WithLogger(context.Background(), logger)

	// Get the logger from the context
	ctxLogger := simplog.FromContext(ctx)

	ctxLogger.Info("You're awesome!")
}
