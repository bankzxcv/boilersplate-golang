package routes

import (
	"fmt"

	"github.com/bankzxcv/streamproject/controllers"
	"github.com/gorilla/mux"
)

func Router(r *mux.Router) {
	fmt.Println("Run Routers")
	// Initial Controller

	uc := controllers.NewItemsController(nil)
	// Route Handler
	r.HandleFunc("/v1/", controllers.Check)
	r.HandleFunc("/v1/item1", uc.Get).Methods("GET")
	r.HandleFunc("/v1/item2", uc.Create).Methods("GET")
	r.HandleFunc("/v1/item3", uc.Update).Methods("GET")
	r.HandleFunc("/v1/item4", uc.Delete).Methods("GET")
}
