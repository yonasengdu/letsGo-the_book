package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("starting new server")
	log.Println("server started at :4000")
	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)

}
