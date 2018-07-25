package routes

import (
	"fmt"

	"github.com/bankzxcv/streamproject/controllers"
	"github.com/gorilla/mux"
)

func Router(r *mux.Router) {
	fmt.Println("Run Routers")
	// Initial Controller


	// Route Handler
	r.HandleFunc("/v1/", controllers.Check)
}
