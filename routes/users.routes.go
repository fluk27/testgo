package routes

import (
	"github.com/fluk27/testgo/controllers"
	"github.com/labstack/echo"
)

func userRoutes(route *echo.Echo) {
	user := &controllers.UserController{}
	route.GET("/users", user.ListUserAll)
	route.GET("/users/:id", user.FindByID)
	route.POST("/users",user.Add)

}
