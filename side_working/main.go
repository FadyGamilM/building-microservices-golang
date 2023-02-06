package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("home handler")
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "error while handling the request", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "data => %s", data)
	})

	// listen to a port
	http.ListenAndServe("localhost:3000", nil)
}
