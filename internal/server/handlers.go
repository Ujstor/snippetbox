package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

// Handler
func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.NotFound(w)
		return
	}

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.ServerError(w, err)
	}
}

func (app *Application) SnippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}

	fmt.Fprintf(w, "Specific snippet with id: %d", id)
}

func (app *Application) SnippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.ClientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create Snippet"))
}
