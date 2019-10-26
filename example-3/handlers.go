package main

import (
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("<h1>Not Found</h1>"))
		return
	}
	w.Write([]byte("<h1>This is Home</h1>"))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World</h1>"))
}

func LogHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World with Middlewares</h1>"))
}

func PassAuthenHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Congrats!</h1>"))
}
