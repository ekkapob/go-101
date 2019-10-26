//+build !test

package main

import (
	"log"
	"net/http"
)

func newServer(options map[string]string) {
	http.HandleFunc("/home", HomeHandler)

	http.HandleFunc("/log", Logger(LogHandler))

	http.HandleFunc("/header", AddHeader(Logger(LogHandler)))

	http.HandleFunc("/authen/pass", AddHeader(Authen(Logger(PassAuthenHandler))))
	http.HandleFunc("/authen/fail", Authen(AddHeader(Logger(PassAuthenHandler))))

	http.HandleFunc("/", RootHandler)

	log.Println("server is running at", options["port"])
	log.Fatal(http.ListenAndServe(options["port"], nil))

}
