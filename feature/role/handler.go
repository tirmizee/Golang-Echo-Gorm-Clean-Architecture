package role

import (
	"clean-architect/commons/log"
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

func (h *Handler) FindByIDHandler(c echo.Context) error {
	log.InfoWithID(c, "FindByIDHandler")

	id := c.Param("id")

	res, err := h.service.FindByID(c, id)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) AllRoleHandler(c echo.Context) error {
	log.InfoWithID(c, "AllRoleHandler")

	res, err := h.service.AllRole(c)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
