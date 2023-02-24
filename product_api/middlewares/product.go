package middlewares

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/FadyGamilM/product_api/data"
)

type KeyProduct struct{}

type extendedProduct data.Product

func RequestValidator(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// define a new placeholder to format the json request body into it
		prod := new(data.Product)

		// get the json req body and decode it into our data.Product entity type
		decodingErr := json.NewDecoder(r.Body).Decode(prod)

		// handle the error of decoding if exists
		if decodingErr != nil {
			http.Error(w, "error while decoding the body of the request", http.StatusBadRequest)
			return
		}

		// if there is no error, update the context of the req to inject the decoded req.Body into it
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)

		// update the request
		r = r.WithContext(ctx)

		// call the next http handler
		nextHandler.ServeHTTP(w, r)
	})
}
