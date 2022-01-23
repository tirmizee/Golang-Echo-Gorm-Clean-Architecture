package role

import (
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

	res, err := h.service.FindByID(id)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, "Not found")
		}

		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) AllRoleHandler(c echo.Context) error {

	res, err := h.service.AllRole()
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, "Not found")
		}

		return err
	}

	return c.JSON(http.StatusOK, res)
}
