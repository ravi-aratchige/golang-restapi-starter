package routes

// Registers all routers found in the routes directory.
// This allows binding the routers to an HTTP web server.
func RegisterAllRoutes() {
	// Register all todo routes
	registerTodoRoutes()
}
