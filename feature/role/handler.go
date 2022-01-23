package role

import (
	"clean-architect/commons/log"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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

	id := c.Param("id")

	res, err := h.service.FindByID(c, id)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, "Not found")
		}

		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) AllRoleHandler(c echo.Context) error {

	log.InfoWithID(c, "AllRoleHandler", "dddd")

	res, err := h.service.AllRole(c)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, "Not found")
		}

		return err
	}

	return c.JSON(http.StatusOK, res)
}
