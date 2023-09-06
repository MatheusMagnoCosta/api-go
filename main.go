package main

import (
	"github.com/gin-gonic/gin"
	"main/api/routes"
)

func main() {
	app := gin.Default()

	routes.AppRoutes(app)

	app.Run("localhost:3001")
}
