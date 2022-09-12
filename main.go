package main

import (
	"local.package/Model"
    
	_ "github.com/mattn/go-sqlite3"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()


    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    
    r.GET("/getUserName", func(c *gin.Context) {
		Model.GetUserName()
	})



    r.Run(":3000")
}
