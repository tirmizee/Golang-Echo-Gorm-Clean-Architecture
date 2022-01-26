package middleware

import (
	commons "clean-architect/commons/error"
	"clean-architect/config"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func GlobalErrorHandler(e *echo.Echo) func(error, echo.Context) {
	return func(err error, c echo.Context) {

		if c.Response().Committed {
			return
		}

		if httpError, ok := err.(*commons.CustomError); ok {
			code := strings.ToLower(httpError.Code)
			httpError.Message = config.ERRConfig[code]
			c.Logger().Error(c.JSON(http.StatusOK, httpError))
		} else {
			e.DefaultHTTPErrorHandler(err, c)
		}

	}
}
