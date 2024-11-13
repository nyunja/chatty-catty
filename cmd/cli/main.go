package main

import (
	"chatty/config"
	"chatty/lib"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	fmt.Println("Therapy Cat Chatbot is starting... 🐱‍💻🐱‍💻")
	cfg, err := config.LoadConfig()
	if err!= nil {
        log.Fatalf("Error: %v", err)
        return
    }
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv(cfg.APIKey)))
	if err != nil {
		log.Fatalln("Error: unable to initiate client")
		return
	}
	lib.PrintChat(ctx, client)
}
