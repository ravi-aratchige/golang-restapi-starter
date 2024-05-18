package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"rawdog-web-server/config"
	"rawdog-web-server/models"
)

// Connect to database
var db = config.ConnectToDatabase()

// ---------------------------------------------------------------------------------------

// Controller to retrieve all todos from the database
var GetAllTodos = func(w http.ResponseWriter, r *http.Request) {
	// Log controller invocation to console
	log.Println("The getAllTodos controller has been fired!")

	// Retrieve all todos from the database
	var todos []models.Todo
	result := db.Find(&todos)

	// Check for errors in database query
	if result.Error != nil {
		// Send 500 internal server error response to the client
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Set response header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Create response
	response := config.MessageAndDataResponse{
		Message: "Todos retrieved successfully",
		Data:    todos,
	}

	// Encode response to JSON and write to HTTP response
	json.NewEncoder(w).Encode(response)
}

// ---------------------------------------------------------------------------------------

// Controller to retrieve todo from the database based on ID
var GetTodoById = func(w http.ResponseWriter, r *http.Request) {
	// Log controller invocation to console
	log.Println("The getTodoById controller has been fired!")

	// Get the ID from the URL parameters
	id := r.PathValue("id")

	// Retrieve the todo from the database
	var todo models.Todo
	result := db.First(&todo, id)

	// Check for errors in database query
	if result.Error != nil {
		// Send 404 not found response to the client
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	// Set response header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Create response
	response := config.MessageAndDataResponse{
		Message: "Todo retrieved successfully",
		Data:    todo,
	}

	// Encode response to JSON and write to HTTP response
	json.NewEncoder(w).Encode(response)
}

// ---------------------------------------------------------------------------------------

// Controller to create new todo in the database
var CreateTodo = func(w http.ResponseWriter, r *http.Request) {
	// Log controller invocation to console
	log.Println("The createTodo controller has been fired!")

	// Decode the request body into a Todo struct
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		// Send 400 bad request response to the client
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save the new todo to the database
	result := db.Create(&todo)

	// Check for errors in database query
	if result.Error != nil {
		// Send 500 internal server error response to the client
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Set response header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Create response
	response := config.MessageAndDataResponse{
		Message: "Todo created successfully",
		Data:    todo,
	}

	// Encode response to JSON and write to HTTP response
	json.NewEncoder(w).Encode(response)
}

// ---------------------------------------------------------------------------------------

// Controller to update todo in the database based on ID
var UpdateTodoById = func(w http.ResponseWriter, r *http.Request) {
	// Log controller invocation to console
	log.Println("The updateTodoById controller has been fired!")

	// Get the ID from the URL parameters
	id := r.PathValue("id")

	// Retrieve the existing todo from the database
	var todo models.Todo
	result := db.First(&todo, id)

	// Check for errors in database query
	if result.Error != nil {
		// Send 404 not found response to the client
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	// Decode the request body into a Todo struct
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		// Send 400 bad request response to the client
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save the updated todo to the database
	result = db.Save(&todo)

	// Check for errors in database query
	if result.Error != nil {
		// Send 500 internal server error response to the client
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Set response header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Create response
	response := config.MessageAndDataResponse{
		Message: "Todo updated successfully",
		Data:    todo,
	}

	// Encode response to JSON and write to HTTP response
	json.NewEncoder(w).Encode(response)
}

// ---------------------------------------------------------------------------------------

// Controller to delete todo in the database based on ID
var DeleteTodoById = func(w http.ResponseWriter, r *http.Request) {
	// Log controller invocation to console
	log.Println("The deleteTodoById controller has been fired!")

	// Get the ID from the URL parameters
	id := r.PathValue("id")

	// Delete the todo from the database
	result := db.Delete(&models.Todo{}, id)

	// Check for errors in database query
	if result.Error != nil {
		// Send 500 internal server error response to the client
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Set response header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Create response
	response := config.MessageResponse{
		Message: "Todo deleted successfully",
	}

	// Encode response to JSON and write to HTTP response
	json.NewEncoder(w).Encode(response)
}
