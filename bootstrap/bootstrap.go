package bootstrap

import (
	"fmt"
	"img_tag/common/conn"
	"img_tag/common/variable"
	_ "img_tag/config"
	"img_tag/model/dao"
	"img_tag/model/schema"

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
	}

	variable.Db.AutoMigrate(&schema.User{})

	dao.CreateUserDaoFactory().Find()
	// variable.Db.Find(&user, 1)
	// fmt.Println(user, 11111)

}
