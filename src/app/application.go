package app

import (
	"github.com/dzikrisyafi/kursusvirtual_utils-go/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	router.Use(cors.Default())
	mapUrls()

	logger.Info("start the application...")
	router.Run(":8001")
}
