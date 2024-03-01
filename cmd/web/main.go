package main

import (
	"database/sql"
	"embed"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/ujstor/snippetbox/internal/server"
	"github.com/ujstor/snippetbox/internal/models"
	"github.com/pressly/goose/v3"	

	_ "github.com/go-sql-driver/mysql"
)


//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dns", "ujstor:password1234@tcp(localhost)/snippetbox?parseTime=true", "MySql data source name" )

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()


	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("mysql"); err != nil {
        panic(err)
    }
	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatalf("goose: failed goose up migration: %v", err)
	}
	if err := goose.Status(db, "migrations"); err != nil {
		log.Fatalf("goose: failed goose status: %v", err)
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