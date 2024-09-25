package main

import (
	"fmt"
	"pigeon/controllers"
	"pigeon/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// // 读取配置
	// config, err := utils.LoadConfig("config.yaml")
	// if err != nil {
	// 	log.Fatalf("读取配置文件失败: %v", err)
	// }

	// 设置路由
	controllers.GetJokeByPage(router)
	controllers.RecordUserAct(router)

	// 启动服务器
	port := fmt.Sprintf(":%d", utils.GlobalConfig.Server.Port)
	fmt.Print(port)
	// if err := router.Run(port); err != nil {
	// 	panic(err)
	// }
}
