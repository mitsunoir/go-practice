package http

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mitsunoir/go-practice/usecase"
)

type createAuthorRequest struct {
	Name string `json:"name"`
}

type createAuthorResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewCreateAuthorHandler(uc *usecase.CreateAuthorUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {

		var req createAuthorRequest
		if err := c.Bind(&req); err != nil {
			return err
		}
		slog.Info("CreateAuthor", "req", req)
		ctx := c.Request().Context()
		a, err := uc.Execute(ctx, &usecase.CreateAuthorRequest{Name: req.Name})
		if err != nil {
			// TODO JSON object として返す
			// TODO エラーを domain で定義してちゃんと判定して返す
			return c.JSON(http.StatusBadRequest, "You are wrong")
		}
		return c.JSON(http.StatusOK, &createAuthorResponse{
			ID:   a.ID,
			Name: a.Name,
		})
	}
}
