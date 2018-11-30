package routes

import (
	"github.com/labstack/echo"
)

func init() {
	e := echo.New()
	userRoutes(e)
	booksRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
