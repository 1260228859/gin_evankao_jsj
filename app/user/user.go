package user

import (
	"fmt"
	"img_tag/service"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	a := c.Param("id")
	a1 := c.Request.Form
	fmt.Println(a, a1, 111)
	user := service.User().Find()
	c.JSON(200, user)
}
