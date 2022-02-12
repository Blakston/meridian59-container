package main

import (
	"log"
	"net/http"

	"github.com/xi2/httpgzip"
)

func main() {
	// setup the client patcher
	http.Handle("/client/", http.StripPrefix("/client/", httpgzip.NewHandler(http.FileServer(http.Dir("client")), nil)))
	// setup download page
	http.Handle("/download/", http.StripPrefix("/download/", httpgzip.NewHandler(http.FileServer(http.Dir("download")), nil)))
	// start listening
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
