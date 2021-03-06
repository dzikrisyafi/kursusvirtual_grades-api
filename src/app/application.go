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
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"OPTIONS", "GET", "POST", "PUT", "DELETE"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Content-Length"},
	}))
	mapUrls()

	logger.Info("start the application...")
	router.Run(":8005")
}
