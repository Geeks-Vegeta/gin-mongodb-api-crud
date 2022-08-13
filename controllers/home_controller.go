package controllers

import (
	"app/configs"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NameController() gin.HandlerFunc {
	return func(c *gin.Context) {
		envname := configs.GetNameFromEnv()
		x := fmt.Sprintf("welcome to my website my name is %s and this is initial route", envname)
		c.JSON(http.StatusOK, gin.H{
			"message": x,
		})
	}
}
