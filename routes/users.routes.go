package routes

import (
	"github.com/fluk27/testgo/controllers"
	"github.com/labstack/echo"
)

func userRoutes(route *echo.Echo) {
	user := &controllers.UserController{}
	route.GET("/users", user.List)
}
