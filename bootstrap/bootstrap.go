package bootstrap

import (
	"fmt"
	"img_tag/common/db"
	"img_tag/common/variable"
	_ "img_tag/config"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func init() {
	// 启动准备事件

	// 开启debug模式 TODO: 获取环境变量 控制debug开启模式
	gin.SetMode(gin.DebugMode)

	// 连接mysql
	if dbMysql, err := db.InitMysqlClient(); err != nil {

	} else {
		variable.Db = dbMysql
	}

}
