package main

import (
	"fmt"
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

	w.Write([]byte("Hello from the Home page"))
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

// use the http.NewServerMux() function to initiallize a new server
// register the home func as the handler for the "/" URL pattern
func main() {

	server := http.NewServeMux()

	server.HandleFunc("/", home)
	server.HandleFunc("/snippet", showSnippets)
	server.HandleFunc("/snippet/create", createSnippet)

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil.

	log.Println("Starting the server at port :4000")

	err := http.ListenAndServe(":4000", server)

	log.Fatal(err)

}
