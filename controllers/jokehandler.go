package controllers

import (
	"fmt"
	"net/http"
	"pigeon/dao"
	"pigeon/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// request jokes
type JokeReq struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

func GetJokeByPage(router *gin.Engine) {
	router.POST("/getJokeByPage", func(c *gin.Context) {
		deviceId := c.GetHeader("deviceId")
		if deviceId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"code": 103, "message": "Invalid Header"})
			return
		}

		var data JokeReq
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 101, "message": "Invalid JSON data"})
			return
		}

		jokes, err := dao.GetJokeByPage(deviceId, data.Page, data.PageSize)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": 201, "message": "Error fetching jokes"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "Success", "data": jokes})
	})
}

// record user actions
type UserActReq struct {
	JokeId  string `json:"jokeId"`
	ActType string `json:"actType"`
}

func RecordUserAct(router *gin.Engine) {
	router.POST("/recordUserAct", func(c *gin.Context) {
		deviceId := c.GetHeader("deviceId")
		if deviceId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"code": 103, "message": "Invalid Header"})
			return
		}

		var data UserActReq
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 101, "message": "Invalid JSON data"})
			return
		}

		if !(data.ActType == "like" || data.ActType == "show") {
			c.JSON(http.StatusBadRequest, gin.H{"code": 102, "message": "Wrong actType value"})
			return
		}

		err := dao.InsertUserAct(deviceId, data.JokeId, data.ActType)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 301, "message": "Error inserting record"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "Success"})
	})
}

func GetUserShowHistory(router *gin.Engine) {
	router.POST("/getUserShowHistory", func(c *gin.Context) {
		deviceId := c.GetHeader("deviceId")
		if deviceId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"code": 103, "message": "Invalid Header"})
			return
		}

		jokes, err := dao.GetUserShowHistory(deviceId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 201, "message": "Error fetching user show history"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "Success", "data": jokes})
	})
}

func GetFreeJokes(router *gin.Engine) {
	router.GET("/getFreeJokes", func(c *gin.Context) {
		start, startErr := strconv.Atoi(c.DefaultQuery("start", "0"))
		end, endErr := strconv.Atoi(c.DefaultQuery("end", "9"))

		if startErr != nil || endErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 102, "message": "Start or end value type is wrong"})
			return
		}

		freeJokeResp, err := dao.GetFreeJokes(start, end)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 201, "message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, freeJokeResp)
	})
}

func UpdateFreeJokes(router *gin.Engine) {
	router.GET("/updateFreeJokes", func(c *gin.Context) {

		err := dao.UpsertFreeJokes()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 301, "message": "Error upserting record"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "Success"})
	})
}

func Test(router *gin.Engine) {
	router.GET("/test", func(c *gin.Context) {
		fmt.Println(utils.GetNowDate())
	})
}
