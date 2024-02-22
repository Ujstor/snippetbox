package main

import (
	"log"
	"net/http"
)

// Handler
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Home page"))
}

func SnippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("View Snippet"))
}

func SnippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		return
	}
	w.Write([]byte("Create Snippet"))
}

func main() {
	//Router
	mux := http.NewServeMux()

	mux.HandleFunc("/", Home)
	mux.HandleFunc("/snippet/view", SnippetView)
	mux.HandleFunc("/snippet/create", SnippetCreate)

	log.Printf("Starting server on port 8088")

	//Server
	err := http.ListenAndServe(":8088", mux)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
