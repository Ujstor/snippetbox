package main

import (
	"log"
	"net/http"

	"github.com/ujstor/snippetbox/internal/server"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", server.Home)
	mux.HandleFunc("/snippet/view", server.SnippetView)
	mux.HandleFunc("/snippet/create", server.SnippetCreate)

	log.Printf("Starting server on port 8088")

	err := http.ListenAndServe(":8088", mux)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}