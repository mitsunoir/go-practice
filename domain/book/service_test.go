package book_test

import (
	"context"
	"errors"
	"testing"

	"github.com/mitsunoir/go-practice/domain/book"
	mock_author "github.com/mitsunoir/go-practice/mock/author"
	mock_book "github.com/mitsunoir/go-practice/mock/book"
	gomock "go.uber.org/mock/gomock"
)

// 本当は mock じゃなくて fake が欲しい
func TestBookService_Create(t *testing.T) {
	t.Run("Error: Title must not be empty", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mb := mock_book.NewMockRepository(ctrl)
		mb.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(nil, nil).
			Times(0)

		ma := mock_author.NewMockRepository(ctrl)
		ma.
			EXPECT().FindByID(gomock.Any(), gomock.Any()).
			Return(nil, nil).
			Times(0)

		mv := book.NewValidator(ma)

		svc := book.NewService(mb, mv)
		ctx := context.Background()
		b, err := svc.CreateBook(ctx, &book.CreateBookRequest{Title: "", Authors: []string{}})
		if err == nil || b != nil {
			t.Fatal("Book title validation doesn't work!!")
		}
	})

	t.Run("Error: Authors must exist", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mb := mock_book.NewMockRepository(ctrl)
		mb.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(nil, nil).
			Times(0)

		ma := mock_author.NewMockRepository(ctrl)
		ma.
			EXPECT().FindByID(gomock.Any(), gomock.Any()).
			Return(nil, errors.New("Author id-of-jane-doe doesn't exist.")).
			Times(1)

		mv := book.NewValidator(ma)

		svc := book.NewService(mb, mv)
		ctx := context.Background()
		b, err := svc.CreateBook(ctx, &book.CreateBookRequest{Title: "Testing Book", Authors: []string{"id-of-jane-doe"}})
		if err == nil || b != nil {
			t.Fatal("Author existence validation doesn't work!!")
		}
	})

	t.Run("Success: Normal case", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mb := mock_book.NewMockRepository(ctrl)
		mb.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(nil, nil)

		ma := mock_author.NewMockRepository(ctrl)
		ma.
			EXPECT().FindByID(gomock.Any(), gomock.Any()).
			Return(nil, nil).
			Times(1)

		mv := book.NewValidator(ma)

		svc := book.NewService(mb, mv)
		ctx := context.Background()
		_, err := svc.CreateBook(ctx, &book.CreateBookRequest{Title: "Testing Book", Authors: []string{"existing-author-id"}})
		if err != nil {
			t.Fatal("Failed to create book!!", err)
		}
	})
}
