package user

import (
	"img_tag/service"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	user := service.User().Find()
	c.JSON(200, user)
}