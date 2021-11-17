package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter The new router function creates the router and
// returns it to us. We can now use this function
// to instantiate and test the router outside of the main function
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	return r
}

func main() {
	// The router is now formed by calling the `newRouter` constructor function
	// that we defined above. The rest of the code stays the same
	r := NewRouter()
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}

func handler(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "Hello World!")
	if err != nil {
		return
	}
}
