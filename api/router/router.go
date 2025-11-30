package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	handler "github.com/mitsunoir/go-practice/api/handler"
)

func Setup(bookHandler *handler.BookHandler) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())

	api := e.Group("/api/v1")
	book := api.Group("/books")
	book.GET("/:id", bookHandler.GetBook)
	book.POST("", bookHandler.CreateBook)
	return e
}
