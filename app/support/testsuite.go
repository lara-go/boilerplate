package support

import (
	"github.com/lara-go/boilerplate/app"
	"github.com/lara-go/larago"
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
	application := app.Bootstrap("", "", "")

	s.BootstrapApplication(application)

	return application
}
