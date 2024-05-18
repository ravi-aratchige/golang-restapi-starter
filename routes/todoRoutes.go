package routes

import (
	"net/http"
	"rawdog-web-server/controllers"
)

// Router to manage todo routes
func registerTodoRoutes() {
	// Get all todos
	http.HandleFunc("GET /todos/", controllers.GetAllTodos)

	// Get todo by ID
	http.HandleFunc("GET /todos/{id}/", controllers.GetTodoById)

	// Create todo
	http.HandleFunc("POST /todos/", controllers.CreateTodo)

	// Update todo by ID
	http.HandleFunc("PATCH /todos/{id}/", controllers.UpdateTodoById)

	// Delete todo by ID
	http.HandleFunc("DELETE /todos/{id}/", controllers.DeleteTodoById)
}
