package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	db "github.com/mitsunoir/go-practice/infra/db"
	http "github.com/mitsunoir/go-practice/interface/http"
	"github.com/mitsunoir/go-practice/usecase"
)

func main() {
	bookRepo := db.NewBookRepository()
	authorRepo := db.NewAuthorRepository()

	createBookUseCase := usecase.NewCreateBookUseCase(bookRepo, authorRepo)
	createBookHandler := http.NewCreateBookHandler(createBookUseCase)

	getBookUseCase := usecase.NewGetBookUseCase(bookRepo)
	getBookHandler := http.NewGetBookHandler(getBookUseCase)

	createAuthorUseCase := usecase.NewCreateAuthorUseCase(authorRepo)
	createAuthorHandler := http.NewCreateAuthorHandler(createAuthorUseCase)

	getAuthorUseCase := usecase.NewGetAuthorUseCase(authorRepo)
	getAuthorHandler := http.NewGetAuthorHandler(getAuthorUseCase)

	e := echo.New()
	e.Use(middleware.Logger())

	api := e.Group("/api/v1")

	author := api.Group("/authors")
	author.GET("/:id", getAuthorHandler)
	author.POST("", createAuthorHandler)

	book := api.Group("/books")
	book.GET("/:id", getBookHandler)
	book.POST("", createBookHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
