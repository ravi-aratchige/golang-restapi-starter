package main

import (
	"fmt"
	"log"
	"net/http"
	"rawdog-web-server/routes"
)

func main() {
	// Define port
	port := ":8080"

	// Define root route and handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello from the webserver!")
		fmt.Fprint(w, "Hello from the webserver!")
	})

	// Bind routers to webserver
	routes.RegisterAllRoutes()

	// Listen on 8080
	log.Printf("Starting the server on %v", port)
	http.ListenAndServe(port, nil)
}
