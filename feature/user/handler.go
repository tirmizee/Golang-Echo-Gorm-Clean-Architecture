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

	res, err := h.service.AllUser(c)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) CreateUserHandler(c echo.Context) error {

	var req CreateUserReq

	if err := c.Bind(&req); err != nil {
		return err
	}

	res, err := h.service.CreateUser(c, &req)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) FindByIDHandler(c echo.Context) error {

	id := c.Param("id")

	res, err := h.service.FindByID(c, id)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
