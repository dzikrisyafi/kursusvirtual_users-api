package app

import (
	"github.com/dzikrisyafi/kursusvirtual_middleware/middleware"
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/controllers/enrolls"
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/controllers/roles"
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/controllers/users"
)

func mapUrls() {
	router.POST("/users/login", users.Login)

	// users group end point
	usersGroup := router.Group("/users")
	usersGroup.Use(middleware.Auth())
	{
		usersGroup.POST("/", users.Create)
		usersGroup.GET("/", users.GetAll)
		usersGroup.GET("/:user_id", users.Get)
		usersGroup.PUT("/:user_id", users.Update)
		usersGroup.PATCH("/:user_id", users.Update)
		usersGroup.DELETE("/:user_id", users.Delete)
	}

	// internal group end point
	internalGroup := router.Group("/internal")
	internalGroup.Use(middleware.Auth())
	{
		internalGroup.GET("/internal/users/search", users.Search)
		internalGroup.GET("/internal/enrolls/:course_id", enrolls.Get)
		internalGroup.POST("/internal/enrolls", enrolls.Create)
	}

	// roles group end point
	rolesGroup := router.Group("/roles")
	rolesGroup.Use(middleware.Auth())
	{
		rolesGroup.GET("/:role_id", roles.Get)
		rolesGroup.GET("/", roles.GetAll)
		rolesGroup.POST("/", roles.Create)
		rolesGroup.PUT("/:role_id", roles.Update)
		rolesGroup.DELETE("/:role_id", roles.Delete)
	}
}
