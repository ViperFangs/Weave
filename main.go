package main

import (
	"log"
	"net/http"
)

func main() {
	const filepathRoot = "."
	const port = "8080"

	// Create a new ServeMux
	mux := http.NewServeMux()
	// Register the http FileServer with the ServeMux
	mux.Handle("/", http.FileServer(http.Dir(filepathRoot)))

	// Create a new server using the ServeMux as the handler
	server := &http.Server{
		Addr:    ":" + port, // Address to listen on
		Handler: mux,        // Handler to use
	}

	// Start the server
	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(server.ListenAndServe())
}
