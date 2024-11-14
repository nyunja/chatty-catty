package main

import (
	h "chatty/web/handlers"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server on port 8000...")
	staticDir := "./static/"
	staticURL := "/static/"
	http.Handle(staticURL, h.CustomStaticServer(staticDir))
	http.HandleFunc("/", h.HomeHandler)
	http.HandleFunc("/chat", h.ChatHandler)
	http.ListenAndServe(":8000", nil)
}
