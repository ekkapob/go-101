package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello World</h1>"))
	})

	http.HandleFunc("/home", HomeHandler)

	http.HandleFunc("/log", Logger(LogHandler))

	http.HandleFunc("/header", AddHeader(Logger(LogHandler)))

	http.HandleFunc("/authen/pass", AddHeader(Authen(Logger(PassAuthenHandler))))
	http.HandleFunc("/authen/fail", Authen(AddHeader(Logger(PassAuthenHandler))))

	log.Println("server is running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
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

func Authen(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, v := range r.Header["custom-header"] {
			if v == "foobar" {
				h(w, r)
				return
			}
		}
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}

func AddHeader(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Header["custom-header"] = append(r.Header["custom-header"], "foobar")
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
