package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	// Get command line arguments
	dir := flag.String("dir", ".", "Hosting directory")
	port := flag.String("port", "8080", "Port")
	flag.Parse()

	// Handler to serve Angular static files
	staticHandler := http.FileServer(http.Dir(*dir))

	// CORS configuration
	corsHandler := cors.Default().Handler(staticHandler)

	// Root path
	http.Handle("/", corsHandler)

	// Start the server on the specified port
	log.Println("Server started at http://localhost:" + *port)
	log.Println("Reading files from: " + *dir)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
