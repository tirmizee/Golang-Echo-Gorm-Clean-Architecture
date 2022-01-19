package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func setupDB() *gorm.DB {
	return nil

}

func setupRoute(db *gorm.DB) *echo.Echo {

	e := echo.New()

	return e
}

func main() {

	var (
		db     = setupDB()
		server = setupRoute(db)
	)

	// server.GET("/users", HandlerPOST)

	server.Logger.Fatal(server.Start(":1323"))
}
