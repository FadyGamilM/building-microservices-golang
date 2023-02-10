package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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

	// create a server to tune some paramters
	server := &http.Server{
		Addr:         "localhost:9090",
		Handler:      customServeMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		// listen and serve using the created server instance
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	// setup the gracefull shutdown for reliability service
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)
	// the program will block untill a ctrl+c or any interrupt signal is happen so we can read it from the channel
	ReceivedSignal := <-signalChannel
	logger.Printf("Main Program exits via a graceful shutdon due to a signal -> ` %s ` has been received", ReceivedSignal)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)

	// listen to a port 9090 and assign the custom serve mux
	// http.ListenAndServe("localhost:9090", customServeMux)
}

/**
http.HandleFunc(PATH, HTTP_HANDLER)
//!=> HandleFunc register a method to a path into something called default serve mux
//!=> default serve mux is a server multiplexer so it allows you to have multiple handlers and match handlers with matched patterns
//!=> ListenAndServe constructs a server and register a http handler to it and if we pass nil, it register the default serve mux
//!=> a Handler in http package is just an interface with one method known as ServeHTTP
//!=> so HandleFunc takes my method and convert it into an http handler
*/
