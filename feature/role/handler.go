package role

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service RoleService
}

func NewHandler(s RoleService) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) AllRoleHandler(c echo.Context) error {

	res, err := h.service.AllRole()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
