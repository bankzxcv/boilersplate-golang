package main

import (
	"fmt"
	"net/http"
	"github.com/bankzxcv/streamproject/routes"
	"github.com/gorilla/mux"
)

func init() {
	fmt.Println("Initial Data")
}

func main() {
	r := mux.NewRouter()
	routes.Router(r)
	http.ListenAndServe(":3000", r)
}
