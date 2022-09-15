package main

import (
	"fmt"
	"net/http"
	"local.package/Controller"
	"local.package/Model"
    "github.com/gin-gonic/gin"
)


var model = Model.NewModel()
var controller = Controller.NewController(model)
var router = gin.Default()


type JsonRequest struct {
	Message string `json:"message"`
}

func main() {

    router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    
    router.GET("/getUserName", func(c *gin.Context) {
		controller.GetUserName()
	})
    
	router.GET("/getUserToken", func(c *gin.Context) {
		controller.GetUserName()
	})

	
	router.POST("/sendMessage", func(c *gin.Context) {
		var json JsonRequest
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		controller.GetUserName()
		c.JSON(http.StatusOK, gin.H{"message": json.Message})

		fmt.Printf("Message:%v\n",json.Message)	
	})


    router.Run(":3000")
}
