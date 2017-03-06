package support

import (
	"github.com/lara-go/app/app/conf"
	"github.com/lara-go/app/app/providers"
	"github.com/lara-go/larago"
	"github.com/lara-go/larago/foundation"
	"github.com/lara-go/larago/support/testsuite"
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
	application := foundation.
		MakeApplication("", "", "").
		SetConfigLoader(conf.ConfigLoader).
		SetApplicationServiceProvider(&providers.ApplicationServiceProvider{})

	s.BootstrapApplication(application)

	return application
}
