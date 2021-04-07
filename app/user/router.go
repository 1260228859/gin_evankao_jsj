package user

import (
	"img_tag/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.Group("user").Use(middleware)
	{
		r.GET("/", GetUser)
	}
}
