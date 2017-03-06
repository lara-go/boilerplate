package main

import (
	"github.com/lara-go/app/app/conf"
	"github.com/lara-go/app/app/providers"
	"github.com/lara-go/larago/foundation"
)

const (
	name        = "larago"
	version     = "0.0.1"
	description = "LaraGo Application"
)

func main() {
	application := foundation.
		MakeApplication(name, version, description).
		SetConfigLoader(conf.ConfigLoader).
		SetApplicationServiceProvider(&providers.ApplicationServiceProvider{})

	foundation.Handle(application)
}
