package providers

import (
	"github.com/lara-go/boilerplate/app/commands"
	"github.com/lara-go/larago"
)

// ApplicationServiceProvider struct.
type ApplicationServiceProvider struct{}

// Register service.
func (p *ApplicationServiceProvider) Register(application *larago.Application) {
	application.Register(
		&HTTPServiceProvider{},
		&MigrationsServiceProvider{},
	)

	application.Commands(
		&commands.HelloWorld{},
	)
}
