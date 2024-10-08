package controllers

import (
	"pigeon/dao"

	"github.com/gin-gonic/gin"
)

// request jokes
type JokeReq struct {
	DeviceId string `json:"deviceId"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
}

func GetJokeByPage(router *gin.Engine) {
	router.POST("/getJokeByPage", func(c *gin.Context) {
		var data JokeReq
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{"code": 101, "message": "Invalid JSON data"})
			return
		}

		jokes, err := dao.GetJokeByPage(data.DeviceId, data.Page, data.PageSize)
		if err != nil {
			c.JSON(500, gin.H{"error": 201, "message": "Error fetching jokes"})
			return
		}

		c.JSON(200, gin.H{"code": 200, "message": "Success", "data": jokes})
	})
}

// record user actions
type UserActReq struct {
	DeviceId string `json:"deviceId"`
	JokeId   string `json:"jokeId"`
	ActType  string `json:"actType"`
}

func RecordUserAct(router *gin.Engine) {
	router.POST("/recordUserAct", func(c *gin.Context) {
		var data UserActReq
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{"code": 101, "message": "Invalid JSON data"})
			return
		}

		if !(data.ActType == "like" || data.ActType == "show") {
			c.JSON(400, gin.H{"code": 102, "message": "Wrong actType value"})
			return
		}

		err := dao.InsertUserAct(data.DeviceId, data.JokeId, data.ActType)
		if err != nil {
			c.JSON(500, gin.H{"code": 301, "message": "Error inserting record"})
			return
		}

		c.JSON(200, gin.H{"code": 200, "message": "Success"})
	})
}

// get user show history
type UserHistoryReq struct {
	DeviceId string `json:"deviceId"`
	UserId   string `json:"userId"`
}

func GetUserShowHistory(router *gin.Engine) {
	router.POST("/getUserShowHistory", func(c *gin.Context) {
		var data UserHistoryReq
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{"code": 101, "message": "Invalid JSON data"})
			return
		}

		jokes, err := dao.GetUserShowHistory(data.DeviceId)
		if err != nil {
			c.JSON(500, gin.H{"code": 201, "message": "Error fetching user show history"})
			return
		}

		c.JSON(200, gin.H{"code": 200, "message": "Success", "data": jokes})
	})
}
