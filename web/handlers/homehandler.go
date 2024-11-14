package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		fmt.Println("index.html not found")
		http.NotFound(w, r)
		return
	}
	tmpl.Execute(w, nil)
}