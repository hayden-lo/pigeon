package controllers

import (
	"pigeon/dao"

	"github.com/gin-gonic/gin"
)

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
