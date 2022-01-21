package routes

import (
	"clean-architect/db"
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

	_ = userRepo
	_ = roleRepo

}
