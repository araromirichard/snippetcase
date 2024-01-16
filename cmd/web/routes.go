package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	server := http.NewServeMux()

	server.HandleFunc("/", app.home)
	server.HandleFunc("/snippet", app.showSnippets)
	server.HandleFunc("/snippet/create", app.createSnippet)

	// create a file server to serve the static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the server.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.

	server.Handle("/static/", http.StripPrefix("/static/", fileServer))

	return server
}
