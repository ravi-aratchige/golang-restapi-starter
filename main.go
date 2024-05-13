package main

import (
	"golang-restapi-starter/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize HTTP server
	app := gin.Default()

	// Initialize base router
	baseRouter := app.Group("/api/v1")

	// Register all routes with base router
	routes.RegisterRoutes(baseRouter)

	// Start server
	app.Run(":8080")
}
