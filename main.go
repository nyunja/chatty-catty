package main

import (
	h "chatty/handlers"
	"chatty/lib"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func main() {
	staticDir := "./static/"
	staticURL := "/static/"
	http.Handle(staticURL, h.CustomStaticServer(staticDir))
	http.HandleFunc("/", h.HomeHandler)
	http.ListenAndServe(":8000", nil)
	fmt.Println("Therapy Cat Chatbot is starting... ���‍��️���‍��️")
	godotenv.Load(".env")
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatalln("Error: unable to initiate client")
		return
	}
	lib.PrintChat(ctx, client)
}
