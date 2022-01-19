package role

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) AllRoleHandler(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}
