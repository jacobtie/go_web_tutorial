package main

import (
	"html/template"
	"log"
	"net/http"
)

// var templates = template.Must(template.ParseFiles("about.html", "notfound.html", "index.html"))

func renderTemplate(w http.ResponseWriter, filename string, data map[string]interface{}) {
	t, _ := template.ParseFiles("base.html", filename)
	t.ExecuteTemplate(w, "base", data)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request), route string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(r.URL.Path[len(route):]) > 1 {
			// templates.ExecuteTemplate(w, "notfound.html", map[string]interface{}{"Route": r.URL.Path[1:]})
			renderTemplate(w, "notfound.html", map[string]interface{}{"Route": r.URL.Path[1:]})
			return
		}
		fn(w, r)
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	// templates.ExecuteTemplate(w, "index.html", nil)
	renderTemplate(w, "index.html", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	// templates.ExecuteTemplate(w, "about.html", nil)
	renderTemplate(w, "about.html", nil)
}

func main() {
	http.HandleFunc("/", makeHandler(mainHandler, "/"))
	http.HandleFunc("/about", makeHandler(aboutHandler, "/about"))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
