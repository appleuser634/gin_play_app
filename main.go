package main

import (
	"local.package/Controller"
	"local.package/Model"
    "github.com/gin-gonic/gin"
)


var model = Model.NewModel()
var controller = Controller.NewController(model)
var router = gin.Default()

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
		controller.GetUserToken()
	})


    router.Run(":3000")
}
