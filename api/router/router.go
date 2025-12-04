package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	handler "github.com/mitsunoir/go-practice/api/handler"
)

func Setup(authorHandler *handler.AuthorHandler, bookHandler *handler.BookHandler) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())

	api := e.Group("/api/v1")

	author := api.Group("/authors")
	author.GET("/:id", authorHandler.GetAuthor)
	author.POST("", authorHandler.CreateAuthor)

	book := api.Group("/books")
	book.GET("/:id", bookHandler.GetBook)
	book.POST("", bookHandler.CreateBook)
	return e
}
