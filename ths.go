package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	// Define command line flags
	dir := flag.String("dir", "", "Hosting directory (required)")
	port := flag.String("port", "8080", "Server port")
	https := flag.Bool("https", false, "Enable HTTPS")
	certFile := flag.String("cert", "", "Certificate file path")
	keyFile := flag.String("key", "", "Key file path")
	corsEnabled := flag.Bool("cors", false, "Enable CORS")
	help := flag.Bool("help", false, "Show help")

	// Parse command line flags
	flag.Parse()

	// Display help if requested or if directory is not provided
	if *help || *dir == "" {
		printHelp()
		return
	}

	// Handler to serve Angular static files
	staticHandler := http.FileServer(http.Dir(*dir))

	// CORS configuration
	var handler http.Handler = staticHandler
	corsStatus := "DISABLED"
	if *corsEnabled {
		handler = cors.Default().Handler(staticHandler)
		corsStatus = "ENABLED"
	}

	http.Handle("/", handler)

	// Start the server on the specified port
	if *https {
		log.Println("Server started at https://localhost:" + *port)
		log.Println("CORS is " + corsStatus)
		log.Println("Hosting files from: " + *dir)
		log.Println("Certificate File: " + *certFile)
		log.Println("Private Key File: " + *keyFile)
		log.Fatal(http.ListenAndServeTLS(":"+*port, *certFile, *keyFile, nil))
	} else {
		log.Println("Server started at http://localhost:" + *port)
		log.Println("Hosting files from: " + *dir)
		log.Fatal(http.ListenAndServe(":"+*port, nil))
	}
}

func printHelp() {
	fmt.Println("Usage: ths [options]")
	fmt.Println()
	fmt.Println("Options:")
	flag.PrintDefaults()
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  ths --dir=/path/to/your/angular-app --port=8000")
	fmt.Println("    - Starts the server on port 8000, hosting the Angular application in /path/to/your/angular-app.")
}
