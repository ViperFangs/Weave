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
	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, req *http.Request) {
		// Before any call to WriteHeader or Write, declare
		// the trailers you will set during the HTTP
		// response. These three headers are actually sent in
		// the trailer.
		w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Create a new server using the ServeMux as the handler
	server := &http.Server{
		Addr:    ":" + port, // Address to listen on
		Handler: mux,        // Handler to use
	}

	// Start the server
	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(server.ListenAndServe())
}
