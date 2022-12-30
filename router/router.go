package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"local.package/controller"
	"local.package/requests"
)

type router struct {
	controller controller.Controller
}

type Router interface {
	SetupRouter() *gin.Engine
}

func NewRouter(controller controller.Controller) *router {
	return &router{controller}
}

func (router *router) SetupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/getUserName", func(c *gin.Context) {
		router.controller.GetUserName()
	})

	r.GET("/getUserToken", func(c *gin.Context) {
		router.controller.GetUserName()
	})

	r.GET("/update", func(c *gin.Context) {
		c.FileAttachment("./esp_bin/file.zip", "filename.zip")
	})

	r.POST("/sendMessage", func(c *gin.Context) {
		var sendMessage requests.SendMessageRequest
		if err := c.ShouldBindJSON(&sendMessage); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		httpStatus := router.controller.SendMessage(sendMessage)

		c.JSON(httpStatus, gin.H{"message": sendMessage.Message, "from": sendMessage.From})
		fmt.Printf("Message:%v\n", sendMessage.Message)
	})

	r.POST("/getMessage", func(c *gin.Context) {
		var getMessage requests.GetMessageRequest
		if err := c.ShouldBindJSON(&getMessage); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("Message To:%v\n", getMessage.To)
		httpStatus, messagelist := router.controller.GetMessage(getMessage)

		c.JSON(httpStatus, gin.H{"messages": messagelist})
	})

	return r
}
