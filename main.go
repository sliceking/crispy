package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"time"
)

var t *template.Template

func main() {
	r := Router{ mux.NewRouter()}
	t = template.Must(template.ParseFiles("views/layout/layout.gohtml"))
	r.mux.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	r.GET("/", HomeHandler)

	srv := &http.Server{
		Handler:      r.mux,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func HomeHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	err := t.Execute(w, "yay")
	if err != nil {
		log.Printf("somethign went wrong: %s", err)
	}
}

// Router is a wrapper for gorilla mux that allows me to hook up routes in a nicer way
type Router struct {
	mux *mux.Router
}

// HandlerFunc is a type to save keystrokes
type HandlerFunc func(http.ResponseWriter, *http.Request)

// GET hooks up a get request
func (r *Router) GET(route string, handler HandlerFunc) {
	r.mux.HandleFunc(route, handler).Methods("GET")
}

// POST hooks up a post request
func (r *Router) POST(route string, handler HandlerFunc) {
	r.mux.HandleFunc(route, handler).Methods("POST")
}

// PUT hooks up a put request
func (r *Router) PUT(route string, handler HandlerFunc) {
	r.mux.HandleFunc(route, handler).Methods("PUT")
}

// DELETE hooks up a delete request
func (r *Router) DELETE(route string, handler HandlerFunc) {
	r.mux.HandleFunc(route, handler).Methods("DELETE")
}
