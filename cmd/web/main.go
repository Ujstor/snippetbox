package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ujstor/snippetbox/internal/models"
	"github.com/ujstor/snippetbox/internal/server"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dns", "ujstor:password1234@tcp(localhost)/snippetbox?parseTime=true", "MySql data source name")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	if err := models.DBMigrations(db); err != nil {
		errorLog.Fatal(err)
	}

	app := &server.Application{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
		Snippets: &models.SnippetModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.Routes(),
	}

	infoLog.Printf("Starting server on port %s", *addr)

	err = srv.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
