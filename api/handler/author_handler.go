package handler

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	author "github.com/mitsunoir/go-practice/domain/author"
)

type AuthorHandler struct {
	Service *author.Service
}

func NewAuthorHandler(s *author.Service) *AuthorHandler {
	return &AuthorHandler{s}
}

type getAuthorRequest struct {
	ID string `param:"id"`
}

type getAuthorResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *AuthorHandler) GetAuthor(c echo.Context) error {
	var req getAuthorRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	slog.Info("GetAuthor", "req", req)
	ctx := c.Request().Context()
	a, err := h.Service.FindAuthor(ctx, req.ID)
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

type createAuthorRequest struct {
	Name string `json:"name"`
}

type createAuthorResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *AuthorHandler) CreateAuthor(c echo.Context) error {
	var req createAuthorRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	slog.Info("CreateAuthor", "req", req)
	ctx := c.Request().Context()
	a, err := h.Service.CreateAuthor(ctx, &author.CreateAuthorRequest{Name: req.Name})
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
