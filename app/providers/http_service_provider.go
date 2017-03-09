package providers

import (
	"github.com/lara-go/boilerplate/app/http/routes"
	"github.com/lara-go/larago"
	larago_middleware "github.com/lara-go/larago/foundation/http/middleware"
	"github.com/lara-go/larago/http"
)

// HTTPServiceProvider struct.
type HTTPServiceProvider struct{}

// Boot provider.
func (p *HTTPServiceProvider) Boot(router *http.Router) {
	// Add ability to automatically validate and inject requests to the route action handler.
	router.SetArgsInjectors(
		&http.RequestsInjector{},
	)

	// Set global middleware.
	router.Middleware(
		&larago_middleware.LogRequests{},
	)

	// Include API routes.
	routes.APIRoutes(router)
}

// Register service.
func (p *HTTPServiceProvider) Register(application *larago.Application) {

}
