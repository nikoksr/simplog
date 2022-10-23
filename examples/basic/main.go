// main implements a basic simplog example.
package main

import "github.com/nikoksr/simplog"

func main() {
	var debug bool

	// Create a client logger; typically used in CLI applications
	clientLogger := simplog.NewClientLogger(debug)

	clientLogger.Info("You're awesome!")
	clientLogger.Warn("Coffee is almost empty!")
	clientLogger.Error("Unable to operate, caffein levels too low.")

	// Create a server logger; typically used in.. well, server applications
	serverLogger := simplog.NewServerLogger(debug)

	serverLogger.Info("You're awesome!")
	serverLogger.Warn("Coffee is almost empty!")
	serverLogger.Error("Unable to operate, caffein levels too low.")
}
