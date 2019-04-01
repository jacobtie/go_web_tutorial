package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

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
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close() // Runs after function finishes

	// Maps the output of the body to the data map
	var data map[string]string
	json.Unmarshal(body, &data)

	// Creates and sets boolean variable to tell if there was a compiler error from input
	compErr := data["compile_errors"] != ""

	// Returns either the compile error message or string and whether a compile error occurred
	return string(data["compile_errors"]) + string(data["output"]), compErr
}
