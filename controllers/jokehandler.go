package controllers

import (
	"log"
	"pigeon/dao"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetJokeByPage(router *gin.Engine) {
	router.POST("/jokes", func(c *gin.Context) {
		page, err := strconv.Atoi(c.PostForm("page"))
		if err != nil {
			c.JSON(400, gin.H{"error": "page parameter is required"})
			return
		}

		pageSize, err := strconv.Atoi(c.DefaultPostForm("pageSize", "5"))
		if err != nil {
			log.Fatalf("Invalid pageSize value: %v", err)
			return
		}

		jokes, err := dao.GetJokeByPage(page, pageSize)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error fetching jokes"})
			return
		}

		c.JSON(200, jokes)
	})
}
