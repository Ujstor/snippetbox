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

	// f, err := os.OpenFile("/tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
	// if err != nil {
	// log.Fatal(err)
	// }
	// defer f.Close()
	// infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)
	
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", server.Home)
	mux.HandleFunc("/snippet/view", server.SnippetView)
	mux.HandleFunc("/snippet/create", server.SnippetCreate)

	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: mux,
	}

	infoLog.Printf("Starting server on port %s", *addr)

	err := srv.ListenAndServe() 
	if err != nil {
		errorLog.Fatal(err)
	}
}
