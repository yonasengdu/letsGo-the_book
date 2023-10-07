package main

import (
	"errors"
	"fmt"

	// "html/template"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/yonasengdu/letsGo-the_book.git/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(
		w,
		http.StatusOK,
		"home.tmpl.html",
		data,
	)

}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {

	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
    



	if err != nil {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrorNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Snippet = snippet

	app.render(
		w,
		http.StatusOK,
		"view.tmpl.html",
		data,
	)

}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display the form for creating a new snippet---"))
}

func (app *application) snippetCreatePos(w http.ResponseWriter, r *http.Request) {
    
	err := r.ParseForm()
	if err != nil {
		app.clientError(w,http.StatusBadRequest)
		return
	}


	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires, err := strconv.Atoi(r.PostForm.Get("expires"))
	if err != nil {
		app.clientError(w,http.StatusBadRequest)
		return
	}


	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
