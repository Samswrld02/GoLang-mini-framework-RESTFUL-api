package routes

import (
	middleware "http/http/Middleware"
	usercontroller "http/http/controllers/UserController"
	"http/router"
	"net/http"
)

// setup usercontrolelr
var userController *usercontroller.UserController = usercontroller.NewUserController()

// setup for the routes
func HandleRoutes() *router.Router {

	//create router instance
	Router := router.NewRouter()

	//user routes and middleware registration
	Router.GET("/user", userController.Index).Middleware = []func(http.ResponseWriter, *http.Request){middleware.LogMiddleware, middleware.PrintMiddleware}

	return Router
}
