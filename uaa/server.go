package main

import (
	"log"
	"net/http"
)

func (c AppContext) newServer(options map[string]string) {
	http.HandleFunc("/oauth2/token", c.TokenHandler)
	http.HandleFunc("/oauth2/token/validate", c.ValidateHandler)
	http.HandleFunc("/oauth2/public_key", c.PublicKeyHandler)

	log.Println("server is running at", options["port"])
	log.Fatal(http.ListenAndServe(options["port"], nil))
}
