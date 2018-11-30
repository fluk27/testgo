package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)
//UserController is struct
type UserController struct {
}
// List function
func (UserController) List(c echo.Context) error {
	return c.String(http.StatusOK, "List user")
}
// Find function
func (UserController)Find(c echo.Context) error {
	return c.String(http.StatusOK, "Find user")
}
