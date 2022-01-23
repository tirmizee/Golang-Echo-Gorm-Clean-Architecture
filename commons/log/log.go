package log

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func InfoWithRequestID(c echo.Context, i ...interface{}) {
	messages := getMessages(c, i)
	log.Info(messages...)
}

func DebugWithRequestID(c echo.Context, i ...interface{}) {
	messages := getMessages(c, i)
	log.Debug(messages...)
}

func ErrorWithRequestID(c echo.Context, i ...interface{}) {
	messages := getMessages(c, i)
	log.Error(messages...)
}

func WarnWithRequestID(c echo.Context, i ...interface{}) {
	messages := getMessages(c, i)
	log.Warn(messages...)
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

func init() {

}
