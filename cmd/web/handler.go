package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Define a home handler func that writes a byte slice containing "hello from the home page" in the response body
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// because this is a subtree url pattern, any url path that ends with a trailing slash
	// will work for this handler so to prevent this we have to put a check to make sure the url patter matches
	// what we intended and return a 404 not found error for any other that does not match

	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		// app.errorLog.Println(err.Error())
		// http.Error(w, "Error parsing template", http.StatusInternalServerError)

		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		// app.errorLog.Println(err.Error())
		// http.Error(w, "Error executing template", http.StatusInternalServerError)
		app.serverError(w, err)
	}

}

// add a show snippets handler function
func (app *application) showSnippets(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Showing specific snippet %d", id)
}

// create snippet handle func
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("this is the create snippet page"))
}
