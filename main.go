package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"local.package/controller"
	"local.package/model"
	"local.package/requests"
)

var mod = model.NewModel()
var con = controller.NewController(mod)
var router = gin.Default()

func main() {

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/getUserName", func(c *gin.Context) {
		con.GetUserName()
	})

	router.GET("/getUserToken", func(c *gin.Context) {
		con.GetUserName()
	})

	router.GET("/update", func(c *gin.Context) {
		c.FileAttachment("./esp_bin/file.zip", "filename.zip")
	})

	router.POST("/sendMessage", func(c *gin.Context) {
		var sendMessage requests.SendMessageRequest
		if err := c.ShouldBindJSON(&sendMessage); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		httpStatus := con.SendMessage(sendMessage)

		c.JSON(httpStatus, gin.H{"message": sendMessage.Message, "from": sendMessage.From})
		fmt.Printf("Message:%v\n", sendMessage.Message)
	})

	router.POST("/getMessage", func(c *gin.Context) {
		var getMessage requests.GetMessageRequest
		if err := c.ShouldBindJSON(&getMessage); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("Message To:%v\n", getMessage.To)
		httpStatus, messagelist := con.GetMessage(getMessage)

		c.JSON(httpStatus, messagelist)
	})

	router.Run(":3000")
}
