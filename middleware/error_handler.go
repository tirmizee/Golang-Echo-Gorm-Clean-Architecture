package middleware

import (
	commons "clean-architect/commons/error"
	"clean-architect/config"
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GlobalErrorHandler(e *echo.Echo) func(error, echo.Context) {
	return func(err error, c echo.Context) {

		if c.Response().Committed {
			return
		}

		if httpError, ok := err.(*commons.CustomError); ok {
			rejectErrMessage(c, httpError)
			return
		}

		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			httpError := commons.NewCustomError("ERR002")
			rejectErrMessage(c, httpError)
		case errors.Is(err, gorm.ErrInvalidData):
			httpError := commons.NewCustomError("ERR003")
			rejectErrMessage(c, httpError)
		case errors.Is(err, gorm.ErrInvalidValueOfLength):
			httpError := commons.NewCustomError("ERR004")
			rejectErrMessage(c, httpError)
		default:
			httpError := commons.NewCustomError("ERR999")
			rejectErrMessage(c, httpError)

		}

	}
}

func rejectErrMessage(c echo.Context, err *commons.CustomError) {
	code := strings.ToLower(err.Code)
	err.Message = config.ERRConfig[code]
	c.Logger().Error(c.JSON(http.StatusOK, err))
}
