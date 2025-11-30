package handler

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	book "github.com/mitsunoir/go-practice/domain/book"
)

type BookHandler struct {
	Service *book.Service
}

func NewBookHandler(s *book.Service) *BookHandler {
	return &BookHandler{s}
}

type getBookRequest struct {
	ID string `param:"id"`
}

type getBookResponse struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func (h *BookHandler) GetBook(c echo.Context) error {
	var req getBookRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	slog.Info("GetBook", "req", req)
	ctx := c.Request().Context()
	b, err := h.Service.FindBook(ctx, req.ID)
	if err != nil {
		// TODO JSON object として返す
		// TODO エラーを domain で定義してちゃんと判定して返す
		return c.JSON(http.StatusBadRequest, "Not found")
	}
	return c.JSON(http.StatusOK, &getBookResponse{
		ID:    b.ID,
		Title: b.Title,
	})
}

type createBookRequest struct {
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
}

type createBookResponse struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func (h *BookHandler) CreateBook(c echo.Context) error {
	var req createBookRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	slog.Info("CreateBook", "req", req)
	b, err := book.NewBook(req.Title, req.Authors)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	b, err = h.Service.CreateBook(ctx, b)
	return c.JSON(http.StatusOK, &createBookResponse{
		ID:    b.ID,
		Title: b.Title,
	})
}
