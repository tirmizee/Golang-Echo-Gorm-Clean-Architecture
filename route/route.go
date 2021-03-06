package route

import (
	"clean-architect/config/db"
	"clean-architect/feature/role"
	"clean-architect/feature/user"
	_middleware "clean-architect/middleware"
	"clean-architect/repository"
	"clean-architect/repository/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func SetupGlobalErrorHandler(e *echo.Echo) {
	e.HTTPErrorHandler = _middleware.GlobalErrorHandler(e)
}

func SetupLogger(e *echo.Echo) {
	e.Logger.SetLevel(log.DEBUG)
}

func SetupMiddleware(e *echo.Echo) {
	// global or root middleware
	e.Use(middleware.RequestID())

}

func SetupRoute(e *echo.Echo) {

	// repository
	var (
		gormDB = db.NewMysqlDB()

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
	g := e.Group("/admin")
	g.GET("/users", userHandler.AllUserHandler)
	g.GET("/users/:id", userHandler.FindByIDHandler)
	g.POST("/users/create", userHandler.CreateUserHandler)
	g.GET("/roles", roleHandler.AllRoleHandler)
	g.GET("/roles/:id", roleHandler.FindByIDHandler)

}
