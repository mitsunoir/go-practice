package http

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mitsunoir/go-practice/usecase"
)

type getBookRequest struct {
	ID string `param:"id"`
}

type getBookResponse struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
}

func NewGetBookHandler(uc *usecase.GetBookUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req getBookRequest
		if err := c.Bind(&req); err != nil {
			return err
		}
		slog.Info("GetBook", "req", req)
		ctx := c.Request().Context()
		b, err := uc.Execute(ctx, req.ID)
		if err != nil {
			// TODO JSON object として返す
			// TODO エラーを domain で定義してちゃんと判定して返す
			return c.JSON(http.StatusBadRequest, "Not found")
		}
		return c.JSON(http.StatusOK, &getBookResponse{
			ID:      b.ID,
			Title:   b.Title,
			Authors: b.Authors,
		})
	}
}
