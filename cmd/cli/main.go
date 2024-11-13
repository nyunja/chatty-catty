package main

import (
	"chatty/config"
	"chatty/lib"
	"context"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Therapy Cat Chatbot is starting... 🐱‍💻🐱‍💻")
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
	ctx := context.Background()
	client, err := lib.NewChatClient(ctx)
	if err != nil {
		log.Fatalln("Error: unable to initiate client")
		return
	}
	lib.PrintChat(ctx, client)
}
