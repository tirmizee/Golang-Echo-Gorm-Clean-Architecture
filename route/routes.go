package routes

import (
	"clean-architect/db"
	"clean-architect/feature/role"
	"clean-architect/feature/user"
	"clean-architect/repository"
	"clean-architect/repository/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupMiddleware(e *echo.Echo) {
	// global or root middleware
	e.Use(middleware.RequestID())

}

func SetupRoute(e *echo.Echo) {

	// repository
	var (
		gormDB                             = db.NewGormDB()
		userRepo repository.UserRepository = mysql.NewUserRepository(gormDB)
		roleRepo repository.RoleRepository = mysql.NewRoleRepository(gormDB)
	)

	// service
	var (
		userService = user.NewUserService(userRepo)
		roleService = role.NewRoleService(roleRepo)
	)

	// handler or controller
	var (
		userHandler = user.NewHandler(userService)
		roleHandler = role.NewHandler(roleService)
	)

	// routes
	e.GET("/users", userHandler.AllUserHandler)
	e.GET("/users/:id", userHandler.AllUserHandler)
	e.GET("/roles", roleHandler.AllRoleHandler)
	e.GET("/roles/:id", userHandler.AllUserHandler)

}
