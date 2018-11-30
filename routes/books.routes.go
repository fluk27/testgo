package routes

import (
	"github.com/fluk27/testgo/controllers"
	"github.com/labstack/echo"
)

func booksRoutes(route *echo.Echo) {
	book := &controllers.BookController{}
	route.GET("/books",book.List)
	route.GET("/books/:id",book.Find)
}
