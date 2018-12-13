package controllers

import (
	"net/http"

	"github.com/fluk27/testgo/services"
	"github.com/labstack/echo"
)

//UserController is struct
type UserController struct {
}

// List function
func (UserController) ListUserAll(c echo.Context) error {
	userService := &services.UserService{}
	users,err := userService.FindAll()
	if err != nil {
		return c.JSON(http.StatusOK,"I don't have value from database")
	}
	return c.JSON(http.StatusOK, users)
}

// Find function
func (UserController) FindByID(c echo.Context) error {
	userService := &services.UserService{}
	valueByFindID := c.Param("id")
	user, err := userService.FindByID(valueByFindID)
	if err != nil || user == nil {
		return c.JSON(http.StatusNotFound, &map[string]interface{}{
			"code" : "FIND_NOT_FOUND",
			"message" : "find not found ",
		})
	}
	return c.JSON(http.StatusOK, user)
}
