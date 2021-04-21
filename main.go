package main

import (
	_ "img_tag/bootstrap"
	"img_tag/routers"

	"github.com/spf13/viper"
)




func main() {
	// 初始化router
	router := routers.InitRouter()
	router.Run(viper.GetString("server.port"))

}
