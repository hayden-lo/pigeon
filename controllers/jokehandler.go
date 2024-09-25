package controllers

import (
	"pigeon/dao"

	"github.com/gin-gonic/gin"
)

// request jokes
type JokeReq struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

func GetJokeByPage(router *gin.Engine) {
	router.POST("/getJokeByPage", func(c *gin.Context) {
		var data JokeReq
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON data"})
			return
		}

		jokes, err := dao.GetJokeByPage(data.Page, data.PageSize)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error fetching jokes"})
			return
		}

		c.JSON(200, jokes)
	})
}

// record user actions
type UserActReq struct {
	JokeId  string `json:"jokeId"`
	ActType string `json:"actType"`
}

func RecordUserAct(router *gin.Engine) {
	router.POST("/RecordUserAct", func(c *gin.Context) {
		var data UserActReq
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON data"})
			return
		}

		err := dao.InsertUserAct(data.JokeId, data.ActType)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error fetching jokes"})
			return
		}

		c.JSON(200, gin.H{"message": "操作成功"})
	})
}
