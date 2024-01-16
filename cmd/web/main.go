package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// Define an application struct to hold the application-wide dependencies for the
// web application.
type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

// use the http.NewServerMux() function to initiallize a new server
// register the home func as the handler for the "/" URL pattern
func main() {

	// define a new command line flag named "port" and a default value of ":4000"
	port := flag.String("addr", ":4000", "port to listen on")

	flag.Parse()

	// create a custom logger to better manage errors and info
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errLog,
		infoLog:  infoLog,
	}

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil.

	srv := &http.Server{
		Addr:     *port,
		ErrorLog: errLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on port %s", *port)

	err := srv.ListenAndServe()

	errLog.Fatal(err)

	// go run ./cmd/web >>/tmp/info.log 2>>/tmp/error.log will help to direct the logs to an on-disk file
}
