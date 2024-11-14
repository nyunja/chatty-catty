package handlers

import (
	"fmt"
	"net/http"
)

func CustomStaticServer(dir string) http.Handler {
	fs := http.FileServer(http.Dir("./"))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Block static files from directory traversal
		switch r.URL.Path {
		case "/static", "/static/", "/static/css", "/static/css/", "/static/js", "/static/js/":
			http.NotFound(w, r)
			fmt.Println("static file issue")
			// functions.ServeError(w, "Invalid path", http.StatusNotFound, "templates/error.html")
			return
		}
		fs.ServeHTTP(w, r)
	})
}
