package usecase_test

import (
	"context"
	"testing"

	mock_author "github.com/mitsunoir/go-practice/mock/author"
	"github.com/mitsunoir/go-practice/usecase"
	"go.uber.org/mock/gomock"
)

func TestCreateAuthorUseCase_(t *testing.T) {
	t.Run("Success: Normal case", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		ma := mock_author.NewMockRepository(ctrl)
		ma.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(nil, nil).
			Times(1)

		uc := usecase.NewCreateAuthorUseCase(ma)
		ctx := context.Background()
		_, err := uc.Execute(ctx, &usecase.CreateAuthorRequest{Name: "Jane Doe"})
		if err != nil {
			t.Fatal("Failed to create author!!", err)
		}
	})
}
