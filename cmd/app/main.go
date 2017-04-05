package main

import (
	"github.com/lara-go/boilerplate/app/conf"
	"github.com/lara-go/boilerplate/app/providers"
	"github.com/lara-go/larago"
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
	// Make application instance.
	application := foundation.MakeApplication(name, version, description)

	// Import application config.
	application.SetConfig(func() larago.Config {
		return conf.NewConfig()
	})

	// Register application services.
	application.Register(
		// Common service providers.
		&cache.ServiceProvider{},
		&database.ServiceProvider{},
		&foundation.ServiceProvider{},
		&http.ServiceProvider{},
		&validation.ServiceProvider{},

		// Default application service provider.
		&providers.ApplicationServiceProvider{},
	)

	// Handle request by application.
	application.Run()
}
