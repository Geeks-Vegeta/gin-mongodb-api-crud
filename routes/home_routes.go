package routes

import (
	"app/controllers"

	"github.com/gin-gonic/gin"
)

func HomeRoutes(router *gin.Engine) {
	router.GET("/", controllers.NameController())

}
