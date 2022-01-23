package route

import (
	"clean-architect/config/db"
	"clean-architect/config/redis"
	"clean-architect/feature/role"
	"clean-architect/feature/user"
	"clean-architect/repository"
	"clean-architect/repository/mysql"
	repositoryredis "clean-architect/repository_redis"

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
		gormDB      = db.NewMysqlDB()
		redisClient = redis.NewRedisClient()

		redisRepo *repositoryredis.RedisRepository = repositoryredis.NewRedisRepository(redisClient)
		userRepo  repository.UserRepository        = mysql.NewUserRepository(gormDB)
		roleRepo  repository.RoleRepository        = mysql.NewRoleRepository(gormDB)
	)

	// service
	var (
		userService = user.NewUserService(redisRepo, userRepo)
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
	g.GET("/users/:id", userHandler.AllUserHandler)
	g.GET("/roles", roleHandler.AllRoleHandler)
	g.GET("/roles/:id", roleHandler.FindByIDHandler)

}
