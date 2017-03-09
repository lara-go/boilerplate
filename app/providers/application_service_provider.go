package providers

import (
	"github.com/lara-go/boilerplate/app/commands"
	"github.com/lara-go/larago"
	"github.com/lara-go/larago/cache"
	"github.com/lara-go/larago/database"
	"github.com/lara-go/larago/foundation"
	"github.com/lara-go/larago/http"
	"github.com/lara-go/larago/validation"
)

// ApplicationServiceProvider struct.
type ApplicationServiceProvider struct{}

// Register service.
func (p *ApplicationServiceProvider) Register(application *larago.Application) {
	application.Register(
		// Register common services.
		&foundation.ServiceProvider{},
		&validation.ServiceProvider{},
		&database.ServiceProvider{},
		&cache.ServiceProvider{},
		&http.ServiceProvider{},

		&HTTPServiceProvider{},
	)

	application.Commands(
		&commands.HelloWorld{},
	)
}
