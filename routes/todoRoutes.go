package routes

import (
	"fmt"
	"golang-restapi-starter/config"
	"golang-restapi-starter/models"

	"github.com/gin-gonic/gin"
)

var db = config.GetDbConnection()

func todoRoutes(superRoute *gin.RouterGroup) {
	subrouter := superRoute.Group("/todos")
	{
		// Route to create a new Todo
		subrouter.POST("/", func(c *gin.Context) {
			var todo models.Todo
			if err := c.ShouldBindJSON(&todo); err != nil {
				c.JSON(400, gin.H{"error": "Invalid JSON data"})
				return
			}

			// Save the Todo to the database
			db.Create(&todo)

			c.JSON(201, todo)
		})

		// Route to get all Todos
		subrouter.GET("/", func(c *gin.Context) {
			var todos []models.Todo

			// Retrieve all Todos from the database
			db.Find(&todos)

			c.JSON(200, todos)
		})

		// Route to get a specific Todo by ID
		subrouter.GET("/:id", func(c *gin.Context) {
			var todo models.Todo
			todoID := c.Param("id")

			// Retrieve the Todo from the database
			result := db.First(&todo, todoID)
			if result.Error != nil {
				c.JSON(404, gin.H{"error": "Todo not found"})
				return
			}

			c.JSON(200, todo)
		})

		// Route to update a Todo by ID
		subrouter.PUT("/:id", func(c *gin.Context) {
			var todo models.Todo
			todoID := c.Param("id")

			// Retrieve the Todo from the database
			result := db.First(&todo, todoID)
			if result.Error != nil {
				c.JSON(404, gin.H{"error": "Todo not found"})
				return
			}

			var updatedTodo models.Todo
			if err := c.ShouldBindJSON(&updatedTodo); err != nil {
				c.JSON(400, gin.H{"error": "Invalid JSON data"})
				return
			}

			// Update the Todo in the database
			todo.Title = updatedTodo.Title
			todo.Description = updatedTodo.Description
			db.Save(&todo)

			c.JSON(200, todo)
		})

		// Route to delete a Todo by ID
		subrouter.DELETE("/:id", func(c *gin.Context) {
			var todo models.Todo
			todoID := c.Param("id")

			// Retrieve the Todo from the database
			result := db.First(&todo, todoID)
			if result.Error != nil {
				c.JSON(404, gin.H{"error": "Todo not found"})
				return
			}

			// Delete the Todo from the database
			db.Delete(&todo)

			c.JSON(200, gin.H{"message": fmt.Sprintf("Todo with ID %s deleted", todoID)})
		})
	}
}
