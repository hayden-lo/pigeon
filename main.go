package main

import (
	"pigeon/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 设置路由
	controllers.GetJokeByPage(router)
	controllers.RecordUserAct(router)

	// 启动服务器
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
