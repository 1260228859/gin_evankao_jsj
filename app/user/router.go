package user

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine) {
	r.GET("/user", GetUser)
}