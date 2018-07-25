package controllers

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"
)

func Check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "System is Okay.")
}

type Items struct {
	Sesson *mgo.Session
}

func NewItemsController(m *mgo.Session) *Items {
	return &Items{m}
}

func (uc Items) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "null")
}

func (uc Items) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "null")
}

func (uc Items) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "null")
}

func (uc Items) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "null")
}
