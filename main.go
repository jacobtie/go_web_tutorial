package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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
		program := r.FormValue("source")
		message, err := compileProgram(program)
		renderTemplate(w, "interpreter.html", map[string]interface{}{"Program": program, "Message": message, "Error": err})
	default:
		renderTemplate(w, "interpreter.html", nil)
	}
}

/* Compiles a program using https://golang.org/compile API endpoint */
func compileProgram(program string) (string, bool) {
	// Sets the endpoint and HTTP method
	url := "https://golang.org/compile"
	method := "POST"

	// Sets the request body using the program
	payload := strings.NewReader("---boundary\n" +
		"Content-Disposition: form-data; name=\"body\"\n\n" +
		program + "\n\n---boundary--")
	// Creates a reference to the HTTP client object
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	// Creates request object
	req, err := http.NewRequest(method, url, payload)

	// Prints the error message if something goes wrong
	if err != nil {
		fmt.Println(err)
	}

	// Adds the header for the boundary in the request body
	req.Header.Add("content-type", "multipart/form-data; boundary=-boundary")

	// Does the request, closes the request, and extracts the body
	res, _ := client.Do(req)
	defer res.Body.Close() // Runs after function finishes
	body, _ := ioutil.ReadAll(res.Body)

	// Maps the output of the body to the data map
	var data map[string]string
	json.Unmarshal(body, &data)

	// Creates and sets boolean variable to tell if there was a compiler error from input
	compErr := data["compile_errors"] != ""

	// Returns either the compile error message or string and whether a compile error occurred
	return string(data["compile_errors"]) + string(data["output"]), compErr
}

func main() {
	http.HandleFunc("/", makeHandler(mainHandler, "/"))
	http.HandleFunc("/about", makeHandler(aboutHandler, "/about"))
	http.HandleFunc("/interpret", makeHandler(interpretHandler, "/interpret"))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
