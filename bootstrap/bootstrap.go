package bootstrap

import (
	"fmt"
	"img_tag/common/conn"
	"img_tag/common/variable"
	_ "img_tag/config"
	"img_tag/model/migrate"

	"github.com/gin-gonic/gin"
)

func init() {
	// 启动准备事件

	// 开启debug模式 TODO: 获取环境变量 控制debug开启模式
	gin.SetMode(gin.DebugMode)

	// 连接mysql
	if dbMysql, err := db.InitMysqlClient(); err != nil {
		fmt.Println(err)
	} else {
		variable.Db = dbMysql
		// 设置数据库
	}

	// 数据库自动迁移脚本
	migrate.InitMysqlAutoMigrate()

}
