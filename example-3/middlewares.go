package main

import (
	"log"
	"net/http"
)

func Authen(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("custom-header") == "foobar" {
			h(w, r)
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func AddHeader(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("custom-header", "foobar")
		h(w, r)
	}
}

func Logger(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Header)
		defer log.Println("---------------------")
		h(w, r)
	}
}
