package user

import (
	// "img_tag/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	userRouter := r.Group("user").Use()
	{
		userRouter.GET("/error", GetUserError)
		userRouter.GET("/id", GetUser)	
	}
}
