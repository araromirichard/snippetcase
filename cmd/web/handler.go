package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Define a home handler func that writes a byte slice containing "hello from the home page" in the response body
func home(w http.ResponseWriter, r *http.Request) {
	// because this is a subtree url pattern, any url path that ends with a trailing slash
	// will work for this handler so to prevent this we have to put a check to make sure the url patter matches
	// what we intended and return a 404 not found error for any other that does not match

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ts, err := template.ParseFiles("./ui/html/home.page.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}

}

// add a show snippets handler function
func showSnippets(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Showing specific snippet %d", id)
}

// create snippet handle func
func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("this is the create snippet page"))
}
