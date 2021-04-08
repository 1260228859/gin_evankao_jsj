package user

import (
	"fmt"
	"img_tag/pkg/parse"
	"img_tag/pkg/consts"
	"img_tag/service"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	a := c.Param("id")
	a1 := c.Request.Form
	fmt.Println(a, a1, 111)
	user := service.User().Find()
	parse.Res(c).Res200(user)

}


func GetUserError(c *gin.Context) {
	parse.Res(c).Res400(consts.JWT_AUTH_ERROR, "", nil)

}