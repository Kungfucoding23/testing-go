package controllers

import (
	"net/http"

	"github.com/Kungfucoding23/testeable_code/services"
	"github.com/gin-gonic/gin"
)

//Ping ...
func Ping(c *gin.Context) {
	c.String(http.StatusOK, services.HandlePing())
}
