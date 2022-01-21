package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service UserService
}

func NewHandler(s UserService) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) AllUserHandler(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}
