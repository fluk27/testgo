package controllers

import (
	"net/http"

	"github.com/fluk27/testgo/services"
	"github.com/labstack/echo"
)

//BookController is struct
type BookController struct {
}

// List this book.controller
func (BookController) List(c echo.Context) error {
	bookService := services.BookService{}
	return c.JSON(http.StatusOK, bookService.FindAll())
}

// Find this book.controller
func (BookController) Find(c echo.Context) error {
	bookService := services.BookService{}
	id := c.Param("id")
	book,err := bookService.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusOK,err.Error())
	}
	return c.JSON(http.StatusOK,book)
}
