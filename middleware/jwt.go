package middleware

// import (
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// func JWTAuth() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 		// 获取token
// 		token := c.Request.Header.Get("x-token")

// 		// 校验token
// 		if b, _ := strconv.ParseBool(token); !b {
// 			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
// 		}
// 	}
// }
