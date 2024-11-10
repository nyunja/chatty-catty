package handlers

import (
	"net/http"
)

func CustomStaticServer(dir string) http.Handler {
	fs := http.FileServer(http.Dir(""))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Block static files from directory traversal
		switch r.URL.Path {
		case "/static", "/static/", "/static/styles", "/static/styles/":
			http.NotFound(w, r)
			// functions.ServeError(w, "Invalid path", http.StatusNotFound, "templates/error.html")
			return
		}
		fs.ServeHTTP(w, r)
	})
}
