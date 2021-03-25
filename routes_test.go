package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(HomeHandler))

	res, err := http.Get(srv.URL)
	if err != nil {
		log.Printf("err %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Error("fucked up")
	}
}
