package routes

import (
	"app/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/name/:name", controllers.UserNameParam())
	router.POST("/user", controllers.CreateUser()) //add this
	router.GET("/yourname", controllers.UserNameQuery())
	router.GET("/user/:userId", controllers.GetAUser())
	router.PUT("/user/:userId", controllers.EditUser())
	router.DELETE("/user/:userId", controllers.DeleteUser())
	router.GET("/users", controllers.GetAllUser())

}
