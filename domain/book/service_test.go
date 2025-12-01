package book_test

import (
	"context"
	"testing"

	"github.com/mitsunoir/go-practice/domain/book"
	mock "github.com/mitsunoir/go-practice/mock"
	gomock "go.uber.org/mock/gomock"
)

// 本当は mock じゃなくて fake が欲しい
func TestBookService_Create(t *testing.T) {
	t.Run("Error: Title must not be empty", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := mock.NewMockRepository(ctrl)
		m.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(nil, nil).
			Times(0)

		svc := book.NewService(m)
		ctx := context.Background()
		b, err := svc.CreateBook(ctx, &book.CreateBookRequest{Title: "", Authors: []string{}})
		if err == nil || b != nil {
			t.Fatal("Book title validation doesn't work!!")
		}
	})

	t.Run("Success: Normal case", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := mock.NewMockRepository(ctrl)
		m.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(nil, nil)

		svc := book.NewService(m)
		ctx := context.Background()
		_, err := svc.CreateBook(ctx, &book.CreateBookRequest{Title: "Testing Book", Authors: []string{}})
		if err != nil {
			t.Fatal("Failed to create book!!", err)
		}
	})
}
