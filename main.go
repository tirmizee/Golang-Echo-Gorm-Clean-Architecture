package main

import (
	"clean-architect/route"

	"github.com/labstack/echo/v4"
)

func main() {

	var server = echo.New()

	route.SetupMiddleware(server)
	route.SetupRoute(server)

	server.Logger.Fatal(server.Start(":8080"))
}
