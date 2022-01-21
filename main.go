package main

import (
	"clean-architect/routes"

	"github.com/labstack/echo/v4"
)

func main() {

	var server = echo.New()

	routes.SetupMiddleware(server)
	routes.SetupRoute(server)

	server.Logger.Fatal(server.Start(":8080"))
}
