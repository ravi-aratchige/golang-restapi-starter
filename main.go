package main

import (
	"golang-restapi-starter/albums"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize new HTTP server
	server := gin.Default()

	// Create root endpoint
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello from the RestAPI!",
		})
	})

	// Endpoint to retrieve all albums
	server.GET("/albums", albums.GetAlbums)

	// Endpoint to retrieve a single album based on ID
	server.GET("/albums/:id", albums.GetAlbumByID)

	// Endpoint to create new album
	server.POST("/albums", albums.PostAlbums)

	// Start HTTP server
	server.Run(":8000")

	// NOTE
	// By default, Gin will start listening on 8080
	// if the port is not explicitly specified.
}
