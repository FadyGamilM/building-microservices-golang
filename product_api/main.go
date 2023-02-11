package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/FadyGamilM/product_api/handlers"
)

func main() {
	// instantiate my serve mux
	serveMux := http.NewServeMux()

	// get the handler method and inject the logger
	logger := log.New(os.Stdout, "[Product_API] => ", log.LstdFlags)
	productHandler := handlers.NewProduct(logger)

	// register the handler to the pattern
	serveMux.Handle("/products", productHandler)

	// construct a server
	server := http.Server{
		Addr:    "localhost:9090",
		Handler: serveMux,
	}

	fmt.Println("Server is running on port 9090")

	server.ListenAndServe()
}
