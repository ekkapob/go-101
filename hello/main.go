package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/slow", slowHandler)
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	html := `
		<form action="/greet" method="post">
			<label>Type your name<input type="text" name="name"></input></label>
			<button type="submit">OK</button>
		</form>
	`
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	html := `
		<h1>{{GREETING}}</h1>
		<br/>
		<a href="/">home</a>
	`
	name := r.PostFormValue("name")
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(strings.ReplaceAll(html, "{{GREETING}}", greet(name))))
}

func slowHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 2)
	w.Write([]byte("slow things"))
}

func greet(name string) string {
	return fmt.Sprint("Hello ", name)
}
