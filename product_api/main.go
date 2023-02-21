package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/FadyGamilM/product_api/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// instantiate my serve mux
	serveMux := mux.NewRouter()

	// get the handler method and inject the logger
	logger := log.New(os.Stdout, "[Product_API] => ", log.LstdFlags)
	productHandler := handlers.NewProduct(logger)

	GET_Router := serveMux.Methods(http.MethodGet).Subrouter()
	GET_Router.HandleFunc("/api/products", productHandler.Get)

	PUT_Router := serveMux.Methods(http.MethodPut).Subrouter()
	PUT_Router.HandleFunc("/api/products/{id}", productHandler.Put)

	// construct a server
	server := http.Server{
		Addr:    "localhost:9090",
		Handler: serveMux,
	}

	fmt.Println("Server is running on port 9090")

	server.ListenAndServe()
}
