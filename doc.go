/*
Package simplog provides a simple logging library for Go. It is factually not a logger itself, but rather helps you by
providing simple and opinionated ways quickly set up uber-go/zap as a logger.

Usage:

	package main

	import "github.com/nikoksr/simplog"

	func main() {
	  // Using the manual configuration; alternatively you can use NewClientLogger() or NewServerLogger().
	  logger := simplog.NewWithOptions(&simplog.Options{
	    Debug:             false,
	    IsServer:          true,
	  })

	  // At this point, you're using a zap.SugaredLogger and can use it as you would normally do.
	  logger.Info("You're awesome!")
	  logger.Warn("Coffee is almost empty!")
	  logger.Error("Unable to operate, caffein levels too low.")
	}
*/
package simplog
