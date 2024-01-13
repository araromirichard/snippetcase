package main

import (
	"log"
	"net/http"
)

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
