package routes

import (
	"github.com/lara-go/boilerplate/app/http/controllers"
	"github.com/lara-go/boilerplate/app/http/middleware"
	"github.com/lara-go/boilerplate/app/http/requests"
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

	router.Group("/group1", func() {
		router.GET("/").Action(func() string {
			return "Group index"
		})

		router.GET("/path").Action(func() string {
			return "Group path"
		})

		router.Group("/group2", func() {
			router.GET("/path").Action(func() string {
				return "Nested group path"
			})
		})
	}, &middleware.Simple{})

	http.Facade().GET("/facade").Action(func() string {
		return "Facade call"
	})
}
