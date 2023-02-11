package handlers

import (
	"log"
	"net/http"

	"github.com/FadyGamilM/product_api/repository"
)

// handler type
type Product struct {
	logger *log.Logger
}

// constructor
func NewProduct(l *log.Logger) *Product {
	return &Product{logger: l}
}

// implement the http.handler
func (handler *Product) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		handler.Get(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (handler *Product) Get(w http.ResponseWriter, r *http.Request) {
	productsData := repository.GetAll()
	err := productsData.ToJson(w)
	if err != nil {
		http.Error(w, "unable to encode the response back", http.StatusInternalServerError)
		handler.logger.Printf("error in product handler => %s", err)
		return
	}
}
