package main

import (
	"html/template"
	"log"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, filename string, data map[string]interface{}) {
	t, _ := template.ParseFiles("base.html", filename)
	t.ExecuteTemplate(w, "base", data)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request), route string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(r.URL.Path[len(route):]) > 1 {
			renderTemplate(w, "notfound.html", map[string]interface{}{"Route": r.URL.Path[1:]})
			return
		}
		fn(w, r)
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html", nil)
}

func interpretHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		renderTemplate(w, "interpreter.html", nil)
	case http.MethodPost:
		r.ParseForm()
		message := r.FormValue("source")
		err := true
		renderTemplate(w, "interpreter.html", map[string]interface{}{"Message": message, "Error": err})
	default:
		renderTemplate(w, "interpreter.html", nil)
	}
}

func main() {
	http.HandleFunc("/", makeHandler(mainHandler, "/"))
	http.HandleFunc("/about", makeHandler(aboutHandler, "/about"))
	http.HandleFunc("/interpret", makeHandler(interpretHandler, "/interpret"))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
