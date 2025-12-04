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
	bookService := book.NewService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	authorRepo := db.NewAuthorRepository()
	authorService := author.NewService(authorRepo)
	authorHandler := handler.NewAuthorHandler(authorService)

	e := router.Setup(authorHandler, bookHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
