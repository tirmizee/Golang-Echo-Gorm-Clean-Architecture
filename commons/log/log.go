package log

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func InfoWithRequestID(c echo.Context, i ...interface{}) {
	log.Info(getMessages(c, i)...)
}

func DebugWithRequestID(c echo.Context, i ...interface{}) {
	log.Debug(getMessages(c, i)...)
}

func ErrorWithRequestID(c echo.Context, i ...interface{}) {
	log.Error(getMessages(c, i)...)
}

func WarnWithRequestID(c echo.Context, i ...interface{}) {
	log.Warn(getMessages(c, i)...)
}

func getRequestID(c echo.Context) string {
	return "requestID -> " + c.Response().Header().Get(echo.HeaderXRequestID) + ", "
}

func getMessages(c echo.Context, i []interface{}) []interface{} {
	requestID := getRequestID(c)
	messages := make([]interface{}, 0)
	messages = append(messages, requestID)
	messages = append(messages, i...)
	return messages
}

func init() {}
