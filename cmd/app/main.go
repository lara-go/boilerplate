package main

import "github.com/lara-go/boilerplate/app"

const (
	name        = "larago"
	version     = "0.0.1"
	description = "LaraGo Application"
)

func main() {
	// Make application instance.
	application := app.Bootstrap(name, version, description)

	// Handle request by application.
	application.Run()
}
