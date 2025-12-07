package http

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mitsunoir/go-practice/usecase"
)

func NewCreateBookHandler(uc *usecase.CreateBookUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req createBookRequest
		if err := c.Bind(&req); err != nil {
			return err
		}
		slog.Info("CreateBook", "req", req)
		ctx := c.Request().Context()
		b, err := uc.Execute(ctx, &usecase.CreateBookRequest{Title: req.Title, Authors: req.Authors})
		if err != nil {
			// TODO JSON object として返す
			// TODO エラーを domain で定義してちゃんと判定して返す
			return c.JSON(http.StatusBadRequest, "You are wrong")
		}
		return c.JSON(http.StatusOK, &createBookResponse{
			ID:      b.ID,
			Title:   b.Title,
			Authors: b.Authors,
		})
	}
}

type createBookRequest struct {
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
}

type createBookResponse struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
}
