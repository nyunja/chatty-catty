package main

import (
	h "chatty/handlers"
	"net/http"
)

func main() {
	staticDir := "./static/"
	staticURL := "/static/"
	http.Handle(staticURL, h.CustomStaticServer(staticDir))
	http.HandleFunc("/", h.HomeHandler)
	http.ListenAndServe(":8000", nil)
}
