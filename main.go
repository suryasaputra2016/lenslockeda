package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/suryasaputra2016/lenslockeda/views"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	path := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, path)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	path := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, path)
}

func executeTemplate(w http.ResponseWriter, path string) {
	tpl, err := views.Parse(path)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error in parsing template", http.StatusInternalServerError)
	}

	tpl.Execute(w, nil)
}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}

}

func main() {
	var router Router

	fmt.Println("Serving on :3000")
	http.ListenAndServe(":3000", router)
}
