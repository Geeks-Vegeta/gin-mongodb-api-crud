package main

import (
	"app/models"
	"app/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	models.Connection()

	routes.UserRoute(router)
	routes.HomeRoutes(router)

	router.Run("localhost:5000")

}
