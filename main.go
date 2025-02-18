package main

import (
	"fmt"
	"log"
	"pigeon/controllers"
	"pigeon/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 设置路由
	controllers.GetJokeByPage(router)
	controllers.RecordUserAct(router)
	controllers.GetFreeJoke(router)

	// 启动服务器
	port := fmt.Sprintf(":%d", utils.GlobalConfig.Server.Port)
	if err := router.Run(port); err != nil {
		log.Print(err)
		panic(err)
	}
}
