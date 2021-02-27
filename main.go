package main

import (
	"github.com/Kungfucoding23/testeable_code/controllers"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main() {
	router.GET("/ping", controllers.Ping)
	router.Run(":8080")
}
