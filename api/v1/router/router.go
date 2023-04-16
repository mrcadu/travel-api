package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"travel-api/internal/config"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()
	err := router.Run("localhost:" + config.GetProperty("SERVER_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	return router
}
