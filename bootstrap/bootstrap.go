package bootstrap

import (
	_ "img_tag/config"

	"github.com/gin-gonic/gin"
)

func init() {
	// 启动准备事件

	// 开启debug模式 TODO: 获取环境变量 控制debug开启模式
	gin.SetMode(gin.DebugMode)
}
