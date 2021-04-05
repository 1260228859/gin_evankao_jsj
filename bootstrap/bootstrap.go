package bootstrap

import (
	"fmt"
	"img_tag/common/db"
	"img_tag/common/variable"
	_ "img_tag/config"
	"img_tag/model"

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

	variable.Db.AutoMigrate(&model.User{})

	var user = model.User{Phone: "18702279164", Name: "高健涛"}
	user.Create()

	model.User{}.Find(&model.User{})
	// variable.Db.Find(&user, 1)
	// fmt.Println(user, 11111)

}
