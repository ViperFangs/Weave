package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Create a new server using the ServeMux as the handler
	server := &http.Server{
			Addr:    "localhost:8080", // Address to listen on
			Handler: mux,              // Handler to use
	}
	
	// Start the server
	fmt.Println("Server started on port 8080")
	server.ListenAndServe()
}