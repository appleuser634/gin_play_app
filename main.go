package main

import (
	"fmt"
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
		controller.GetUserName()
	})

	
	router.POST("/sendMessage", func(c *gin.Context) {
		fmt.Printf("Params Message:%v\n",c.PostForm("message"))
		fmt.Printf("Params From:%v\n",c.PostForm("from"))
		fmt.Printf("Params To:%v\n",c.PostForm("to"))
		fmt.Printf("Params Token:%v\n",c.PostForm("token"))
		controller.GetUserName()
	})


    router.Run(":3000")
}
