package http

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mitsunoir/go-practice/usecase"
)

type getAuthorRequest struct {
	ID string `param:"id"`
}

type getAuthorResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewGetAuthorHandler(uc *usecase.GetAuthorUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req getAuthorRequest
		if err := c.Bind(&req); err != nil {
			return err
		}
		slog.Info("GetAuthor", "req", req)
		ctx := c.Request().Context()
		a, err := uc.Execute(ctx, req.ID)
		if err != nil {
			// TODO JSON object として返す
			// TODO エラーを domain で定義してちゃんと判定して返す
			return c.JSON(http.StatusBadRequest, "Not found")
		}
		return c.JSON(http.StatusOK, &getAuthorResponse{
			ID:   a.ID,
			Name: a.Name,
		})
	}
}
