package main

import (
	"github.com/lara-go/boilerplate/app/conf"
	"github.com/lara-go/boilerplate/app/providers"
	"github.com/lara-go/larago/cache"
	"github.com/lara-go/larago/database"
	"github.com/lara-go/larago/foundation"
	"github.com/lara-go/larago/http"
	"github.com/lara-go/larago/validation"
)

const (
	name        = "larago"
	version     = "0.0.1"
	description = "LaraGo Application"
)

func main() {
	application := foundation.
		MakeApplication(name, version, description).
		SetConfigLoader(conf.ConfigLoader)

	application.Register(
		// Register common service providers.
		&foundation.ServiceProvider{},
		&validation.ServiceProvider{},
		&database.ServiceProvider{},
		&cache.ServiceProvider{},
		&http.ServiceProvider{},

		// Default application service provider.
		&providers.ApplicationServiceProvider{},
	)

	foundation.Handle(application)
}
