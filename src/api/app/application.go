package app

import (
	"MocksExamples/src/api/handlers"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	router.Group("/test")
	router.POST("/operation", func(c *gin.Context) {
		handlers.GetOperation(c)
	})
	router.Run()
}