package app

import (
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/controllers/enrolls"
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/controllers/users"
)

func mapUrls() {
	// user end point
	router.POST("/users", users.Create)
	router.GET("/users", users.GetAll)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)

	router.GET("/internal/users/search", users.Search)
	router.POST("/users/login", users.Login)

	// enroll end point
	router.GET("/internal/enrolls/:course_id", enrolls.Get)
	router.POST("/internal/enrolls", enrolls.Create)
}
