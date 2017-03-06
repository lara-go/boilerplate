package routes

import (
	"github.com/lara-go/app/app/http/controllers"
	"github.com/lara-go/app/app/http/middleware"
	"github.com/lara-go/larago/http"
)

// APIRoutes adds routes for API.
func APIRoutes(router *http.Router) {
	controller := &controllers.APIController{}

	router.GET("/").Action(controller.Index).Middleware(&middleware.Simple{})
	router.GET("/redirect").Action(controller.Redirect)
	router.GET("/foo/:bar").Action(controller.FooBar).As("foobar")
	router.GET("/bar/:baz").Action(controller.BarBaz)
	router.POST("/form").Action(controller.Form)
	router.POST("/json").Action(controller.JSON)
	router.GET("/http-error").Action(controller.HTTPError)
	router.GET("/custom-error").Action(controller.CustomError)
	router.GET("/panic").Action(controller.Panic)
}
