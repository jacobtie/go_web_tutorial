package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// Helper function to render template with base layout
func renderTemplate(w http.ResponseWriter, filename string, data map[string]interface{}) {
	t, _ := template.ParseFiles("base.html", filename)
	t.ExecuteTemplate(w, "base", data)
}

// Middleware to forward any unknown routes to 404 page
func makeHandler(fn func(http.ResponseWriter, *http.Request), route string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// log.Println(route)
		// log.Println(r.URL.Path)
		if len(r.URL.Path[len(route):]) > 1 {
			log.Printf("GET %s not found, routing to 404", r.URL.Path[1:])
			renderTemplate(w, "notfound.html", map[string]interface{}{"Route": r.URL.Path[1:]})
			return
		}
		fn(w, r)
	}
}

// Renders home page
func mainHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GET index")
	renderTemplate(w, "index.html", nil)
}

// Renders about page
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GET about")
	renderTemplate(w, "about.html", nil)
}

// Renders and handles interpreter page
func interpretHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		log.Println("GET interpret")
		// Renders normal interpreter page
		renderTemplate(w, "interpreter.html", nil)
	case http.MethodPost:
		log.Println("POST interpret")
		// Parses data from form
		r.ParseForm()
		if r.FormValue("sample") == "" {
			// Extracts source code
			program := r.FormValue("source")
			// Compiles and runs program with Google's endpoint to get output/error
			message, err := compileProgram(program)
			// Replaces line breaks in message with HTML line breaks
			message = strings.Replace(message, "\n", "\n<br />", -1)
			// Renders template with correct data
			renderTemplate(w, "interpreter.html", map[string]interface{}{"Program": program, "Message": template.HTML(message), "Error": err})
		} else {
			program, err := ioutil.ReadFile("samples/" + r.FormValue("sample") + "/main" + ".go")
			if err != nil {
				log.Println(err)
			}
			renderTemplate(w, "interpreter.html", map[string]interface{}{"Program": string(program)})
		}
	default:
		// Renders normal interpreter page in case a different HTTP method is used
		renderTemplate(w, "interpreter.html", nil)
	}
}

// Handles favicon
func iconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "favicon.ico")
}

// Handles info page
func infoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GET info")
	renderTemplate(w, "info.html", nil)
}

// Handles code libary page
func libraryHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GET libary")
	renderTemplate(w, "library.html", nil)
}

// Entry point of the program
func main() {
	// Sets up static directory serve
	fs := http.FileServer(http.Dir("static/"))
	// Sets up routes
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/favicon.ico", iconHandler)
	http.HandleFunc("/", makeHandler(mainHandler, "/"))
	http.HandleFunc("/about", makeHandler(aboutHandler, "/about"))
	http.HandleFunc("/info", makeHandler(infoHandler, "/info"))
	http.HandleFunc("/interpret", makeHandler(interpretHandler, "/interpret"))
	http.HandleFunc("/library", makeHandler(libraryHandler, "/library"))
	// Starts the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
