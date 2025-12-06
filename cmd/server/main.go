package main

import (
	"github.com/mitsunoir/go-practice/api/handler"
	router "github.com/mitsunoir/go-practice/api/router"
	author "github.com/mitsunoir/go-practice/domain/author"
	book "github.com/mitsunoir/go-practice/domain/book"
	db "github.com/mitsunoir/go-practice/infra/db"
)

func main() {
	bookRepo := db.NewBookRepository()
	authorRepo := db.NewAuthorRepository()
	bookValidator := book.NewValidator(authorRepo)
	bookService := book.NewService(bookRepo, bookValidator)
	bookHandler := handler.NewBookHandler(bookService)

	authorService := author.NewService(authorRepo)
	authorHandler := handler.NewAuthorHandler(authorService)

	e := router.Setup(authorHandler, bookHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
