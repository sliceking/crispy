package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	r := Router{ mux.NewRouter()}
	r.GET("/", HomeHandler)

	srv := &http.Server{
		Handler:      r.mux,
		Addr:         "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func HomeHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "you did gt")
}

type Router struct {
	mux *mux.Router
}

type HandlerFunc func(http.ResponseWriter, *http.Request)

func (r *Router) GET(route string, handler HandlerFunc) {
	r.mux.HandleFunc(route, handler).Methods("GET")
}
