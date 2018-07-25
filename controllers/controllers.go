package controllers

import (
	"fmt"
	"net/http"
)


func Check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "System is Okay.")
}

func Get(w http.ResponseWriter, r *http.Request) {
}

func Create(w http.ResponseWriter, r *http.Request) {
}

func Update(w http.ResponseWriter, r *http.Request) {
}

func Delete(w http.ResponseWriter, r *http.Request) {
}
