package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {

		w.Write([]byte("page not found"))
		return
	}
    
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
		}


	ts, err := template.ParseFiles(files...)
	if err != nil{
		log.Println(err.Error())
		http.Error(w,"Internal serve Error",500)
	}

	err = ts.ExecuteTemplate(w,"base",nil)
	if err != nil{
		log.Println(err.Error())
		http.Error(w,"Internal serve Error",500)

	}

	w.Write([]byte("hello from snippetbox"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return

	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("create snippet"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "showing snippet with Id: %v", id)
	w.Write([]byte("view snippets"))
}
