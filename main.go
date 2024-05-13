package main

import (
	"golang-restapi-starter/albums"
	"golang-restapi-starter/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize new HTTP router
	router := gin.Default()

	// Add custom middleware to server
	router.Use(middleware.LoggerMiddleware())

	// Create root endpoint
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello from the RestAPI!",
		})
	})

	// Endpoint to retrieve all albums
	router.GET("/albums", albums.GetAlbums)

	// Endpoint to retrieve a single album based on ID
	router.GET("/albums/:id", albums.GetAlbumByID)

	// Endpoint to create new album
	router.POST("/albums", albums.PostAlbums)

	// ********** DEMO ROUTES **********

	// Demo route with URL parameters
	router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, "User ID: "+id)
	})

	// Demo route with query parameters
	router.GET("/search", func(c *gin.Context) {
		query := c.DefaultQuery("q", "default-value")
		c.String(200, "Search query: "+query)
	})

	// ********** ROUTE GROUPS **********

	// Public routes (no authentication required)
	public := router.Group("/public")
	{
		public.GET("/info", func(c *gin.Context) {
			c.String(200, "Public information")
		})
		public.GET("/products", func(c *gin.Context) {
			c.String(200, "Public product list")
		})
	}

	// Private routes (require authentication)
	private := router.Group("/private")
	private.Use(middleware.AuthMiddleware())
	{
		private.GET("/data", func(c *gin.Context) {
			c.String(200, "Private data accessible after authentication")
		})
		private.POST("/create", func(c *gin.Context) {
			c.String(200, "Create a new resource")
		})
	}

	// Start HTTP server
	router.Run(":8000")

	// NOTE
	// By default, Gin will start listening on 8080
	// if the port is not explicitly specified.
}
