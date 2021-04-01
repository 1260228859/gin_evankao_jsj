package main

import (
	// "img_tag/routers"
	"img_tag/routers"
)

func main() {
	// 初始化router
	router := routers.InitRouter()
	router.Run("127.0.0.1:8080")

}
