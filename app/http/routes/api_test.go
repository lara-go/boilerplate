package routes_test

import (
	"testing"

	"github.com/lara-go/boilerplate/app/support"
	"github.com/stretchr/testify/suite"
)

type APIRoutesSuite struct {
	suite.Suite
	support.TestSuite
}

func TestAPIRoutesSuit(t *testing.T) {
	suite.Run(t, new(APIRoutesSuite))
}

func (s *APIRoutesSuite) TestAPIRoutes() {
	e := s.HTTP(s.T())

	e.GET("/").Expect().Status(200)
	e.GET("/redirect").Expect().Status(200)
	e.GET("/http-error").Expect().Status(404)
	e.GET("/custom-error").Expect().Status(500)
	e.GET("/panic").Expect().Status(500)
	e.GET("/not-found").Expect().Status(404)

	e.GET("/bar/text1").
		Expect().Status(200).
		Body().Contains("text1")

	e.POST("/form").WithFormField("text", "Text input").
		Expect().Status(200).
		Body().Contains("Text input")

	e.POST("/json").WithHeader("Accept", "application/json").WithJSON(map[string]interface{}{"text": "Text value"}).
		Expect().Status(200).
		JSON().Object().ValueEqual("foo", "Text value")

	e.POST("/json").WithHeader("Accept", "application/json").WithJSON(map[string]interface{}{"text": ""}).
		Expect().Status(422)

	e.POST("/").Expect().Status(405)
	e.POST("/bla-bla-ba").Expect().Status(404)
}
