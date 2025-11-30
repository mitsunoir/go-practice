package main

import (
	"github.com/mitsunoir/go-practice/api/handler"
	router "github.com/mitsunoir/go-practice/api/router"
	book "github.com/mitsunoir/go-practice/domain/book"
	db "github.com/mitsunoir/go-practice/infra/db"
)

func main() {
	bookRepo := db.NewBookRepository()
	bookService := book.NewService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)
	e := router.Setup(bookHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
