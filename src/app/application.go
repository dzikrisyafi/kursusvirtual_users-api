package app

import (
	"github.com/dzikrisyafi/kursusvirtual_utils-go/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("start the application...")
	router.Run(":8001")
}
