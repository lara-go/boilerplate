package routes

import (
	"github.com/lara-go/app/app/http/controllers"
	"github.com/lara-go/app/app/http/middleware"
	"github.com/lara-go/app/app/http/requests"
	"github.com/lara-go/larago/http"
)

// APIRoutes adds routes for API.
func APIRoutes(router *http.Router) {
	controller := &controllers.APIController{}

	// Simple request with custom middleware.
	router.GET("/").Action(controller.Index).Middleware(&middleware.Simple{})

	// Redirect.
	router.GET("/redirect").Action(controller.Redirect)
	// Aliased route with param
	router.GET("/foo/:bar").Action(controller.FooBar).As("foobar")

	// Forms.
	router.GET("/bar/:text").Action(controller.BarBaz)
	router.POST("/form").Action(controller.Form)
	router.POST("/json").Action(controller.JSON).Validate(&requests.PostNews{})

	// Errors.
	router.GET("/http-error").Action(controller.HTTPError)
	router.GET("/custom-error").Action(controller.CustomError)
	router.GET("/panic").Action(controller.Panic)
}
