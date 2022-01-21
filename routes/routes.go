package routes

import (
	"clean-architect/db"
	"clean-architect/features/role"
	"clean-architect/features/user"
	"clean-architect/repositories"
	"clean-architect/repositories/mysql"

	"github.com/labstack/echo/v4"
)

func SetupRoute(e *echo.Echo) {

	// repository
	var (
		gormDB                               = db.GormDB()
		userRepo repositories.UserRepository = mysql.NewUserRepository(gormDB)
		roleRepo repositories.RoleRepository = mysql.NewRoleRepository(gormDB)
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

	e.GET("/users", userHandler.AllUserHandler)
	e.GET("/roles", roleHandler.AllRoleHandler)

}
