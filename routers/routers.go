package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func handle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "helloworld",
	})
	c.
}


func InitRouter() *gin.Engine {
	var router *gin.Engine
	router = gin.Default()

	router.GET("/", handle)

	return router
}
