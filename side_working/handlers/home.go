package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// a type which represent our handler and it will implement the handler interface
type Home struct {
	logger *log.Logger
}

// a parametarized constructor that takes an instance of any type implements the log.Logger interface and return a new ref of Home type
func NewHome(l *log.Logger) *Home {
	return &Home{
		logger: l,
	}
}

func (handler *Home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// utilize the injected logger
	(*handler).logger.Println("home handler")
	// read the request body
	data, err := ioutil.ReadAll(r.Body)
	// check for errors
	if err != nil {
		// bad request because its something with the request body
		http.Error(w, "error while handling the request", http.StatusBadRequest)
		return
	}
	// return the response
	fmt.Fprintf(w, "data => %s", data)
}
