package support

import (
	"github.com/lara-go/boilerplate/app/conf"
	"github.com/lara-go/boilerplate/app/providers"
	"github.com/lara-go/larago"
	"github.com/lara-go/larago/cache"
	"github.com/lara-go/larago/database"
	"github.com/lara-go/larago/foundation"
	"github.com/lara-go/larago/http"
	"github.com/lara-go/larago/support/testsuite"
	"github.com/lara-go/larago/validation"
)

// TestSuite default test suite.
type TestSuite struct {
	testsuite.LaragoSuite
}

// SetupTest make application before every test.
func (s *TestSuite) SetupTest() {
	s.Application = s.MakeApplication()
}

// TearDownSuite .
func (s *TestSuite) TearDownSuite() {
	s.Application = nil
}

// MakeApplication factory.
func (s *TestSuite) MakeApplication() *larago.Application {
	application := foundation.MakeApplication("", "", "")

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

	s.BootstrapApplication(application)

	return application
}
