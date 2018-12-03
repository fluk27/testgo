package controllers

import (
	"fmt"
	"net/http"

	"github.com/fluk27/testgo/services"
	"github.com/labstack/echo"
)

//BookController is struct

// List this book.controller
func  List(c echo.Context) error {
	bookService := &services.BookService{}
	return c.JSON(http.StatusOK, bookService.FindAll())
}

// Find this book.controller
func  Find(c echo.Context) error {
	bookService := services.BookService{}
	id := c.Param("id")
	book,err := bookService.FindByID(id)

	fmt.Println(book.Name)

	book.Name = "New name"

	book2 := book
	book2.Name = "Book 2"

	fmt.Println("Book name", book.Name)
	fmt.Println("Book2 name", book2.Name)

	if err != nil {
		return c.JSON(http.StatusOK,err.Error())
	}
	return c.JSON(http.StatusOK,book)
}
