package main

import (
	"log"
	"net/http"
	"os"

	"github.com/FadyGamilM/side_working/handlers"
)

// main entry point ..
func main() {
	// instantiate the logger instance we will use
	// currently we will use a logger against the cmd
	logger := log.New(os.Stdout, "[Product-API]", log.LstdFlags)

	// create custom serve mux instead of the default-serve-mux
	customServeMux := http.NewServeMux()

	// instance of the home handler and apply dependency injection by injecting the logger instance
	homeHandler := handlers.NewHome(logger)

	// match the specific handler to specific path
	customServeMux.Handle("/home", homeHandler)

	log.Println("Server is up and running on port 9090")

	// listen to a port 9090 and assign the custom serve mux
	http.ListenAndServe("localhost:9090", customServeMux)
}

/**
http.HandleFunc(PATH, HTTP_HANDLER)
//!=> HandleFunc register a method to a path into something called default serve mux
//!=> default serve mux is a server multiplexer so it allows you to have multiple handlers and match handlers with matched patterns
//!=> ListenAndServe constructs a server and register a http handler to it and if we pass nil, it register the default serve mux
//!=> a Handler in http package is just an interface with one method known as ServeHTTP
//!=> so HandleFunc takes my method and convert it into an http handler
*/
