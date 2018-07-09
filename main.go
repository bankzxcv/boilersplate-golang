package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "System is OKAY.")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Check)

	http.ListenAndServe(":3000", r)
}
