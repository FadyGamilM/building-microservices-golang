package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/FadyGamilM/product_api/data"
	"github.com/FadyGamilM/product_api/repository"
	"github.com/gorilla/mux"
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

func (handler *Product) Put(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	prodID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "error while extracting the id of the product", http.StatusBadRequest)
	}
	// create a new placeholder for the given body of the request
	newProduct := new(data.Product)

	err = newProduct.FromJson(r.Body)
	if err != nil {
		http.Error(w, "error while decoding the given json object", http.StatusBadRequest)
	}

	err = repository.UpdateProduct(prodID, newProduct)
	if err != nil {
		http.Error(w, "error while updating the product", http.StatusInternalServerError)
	}
}
