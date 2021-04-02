package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Option func(*gin.Engine)

var options = []Option{}

func handle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": viper.GetString("evankao.val"),
	})
}

// 路由注册:
func Include(opt ...Option) {
	options = append(options, opt...)
}

func addRouter() {
	Include(
	// 添加各模块蓝图 type Option
	// router1
	// router2
	)
}

func InitRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/", handle)

	addRouter()
	for _, opt := range options {
		opt(router)
	}

	return router
}
