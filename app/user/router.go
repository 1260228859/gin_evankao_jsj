package user

import (
	// "img_tag/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	userRouter := r.Group("user").Use()
	{
		userRouter.GET("/:id", GetUser)
	}
}
