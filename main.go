package main

import (
	"clean-architect/route"

	"github.com/labstack/echo/v4"
)

func Hello() string {
	return "hello world"
}

func main() {

	var server = echo.New()

	route.SetupLogger(server)
	route.SetupGlobalErrorHandler(server)
	route.SetupMiddleware(server)
	route.SetupRoute(server)

	server.Logger.Fatal(server.Start(":8080"))
}
