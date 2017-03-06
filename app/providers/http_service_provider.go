package providers

import (
	"github.com/lara-go/app/app/http/routes"
	"github.com/lara-go/larago"
	larago_middleware "github.com/lara-go/larago/foundation/http/middleware"
	"github.com/lara-go/larago/http"
)

// HTTPServiceProvider struct.
type HTTPServiceProvider struct{}

// Boot provider
func (p *HTTPServiceProvider) Boot(router *http.Router) {
	router.Middleware(
		&larago_middleware.LogRequests{},
	)

	routes.APIRoutes(router)
}

// Register service.
func (p *HTTPServiceProvider) Register(application *larago.Application) {

}
