package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/ujstor/snippetbox/internal/server"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &server.Application{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/snippet/view", app.SnippetView)
	mux.HandleFunc("/snippet/create", app.SnippetCreate)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on port %s", *addr)

	err := srv.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}
