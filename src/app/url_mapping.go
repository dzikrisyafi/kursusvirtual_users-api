package app

import (
	"github.com/dzikrisyafi/kursusvirtual_middleware/middleware"
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/controllers/cohort"
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/controllers/departments"
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
		usersGroup.GET("/:user_id", users.Get)
		usersGroup.GET("/", users.GetAll)
		usersGroup.PUT("/:user_id", users.Update)
		usersGroup.PATCH("/:user_id", users.Update)
		usersGroup.DELETE("/users/:user_id", users.Delete)
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

	// departmentsGroup
	departmentsGroup := router.Group("/departments")
	departmentsGroup.Use(middleware.Auth())
	{
		departmentsGroup.POST("/", departments.Create)
		departmentsGroup.GET("/:department_id", departments.Get)
		departmentsGroup.GET("/", departments.GetAll)
		departmentsGroup.PUT("/:department_id", departments.Update)
		departmentsGroup.DELETE("/:department_id", departments.Delete)
	}

	// internal group end point
	internalGroup := router.Group("/internal")
	internalGroup.Use(middleware.Auth())
	{
		internalGroup.POST("/enrolls", enrolls.Create)
		internalGroup.GET("/enrolls/:course_id", enrolls.Get)
		internalGroup.PUT("/enrolls/:user_id/:course_id", enrolls.Update)
		internalGroup.DELETE("/enrolls/:enroll_id", enrolls.Delete)

		internalGroup.POST("/cohorts", cohort.Create)
		internalGroup.GET("/cohorts/:cohort_id", cohort.Get)
		internalGroup.GET("/cohorts", cohort.GetAll)
		internalGroup.PUT("/cohorts/:cohort_id", cohort.Update)
		internalGroup.DELETE("/cohorts/:cohort_id", cohort.Delete)
	}

}
