package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func main() {
	godotenv.Load(".env")
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatalln("Error: unable to initiate client")
		return
	}
	model := client.GenerativeModel("gemini-pro")
	cs := model.StartChat()
	persona, intro, name := getName()
	cs.History = []*genai.Content{
		{
			Parts: []genai.Part{
				genai.Text(persona),
			},
			Role: "user",
		},
		{
			Parts: []genai.Part{
				genai.Text(intro),
			},
			Role: "model",
		},
	}
	os.Stdout.WriteString(intro + "\n")
	for {
		input := make([]byte, 1024)
		fmt.Println(name, " :")
		os.Stdin.Read(input)
		resp, err := cs.SendMessage(ctx, genai.Text(string(input)))
		if err != nil {
			log.Fatalln("Error: unable to retrieve response from the client")
			return
		}
		fmt.Println("Jill: ")
		fmt.Println(resp.Candidates[0].Content.Parts[0])
	}
}

func getName() (string, string, string) {
	name := make([]byte, 100)
	os.Stdout.WriteString("Hi, I am Jill, your therapy cat. What is your name?\n")
	n, _ := os.Stdin.Read(name)
	os.Stdout.WriteString("\n")
	persona := fmt.Sprintf(`Pretend you are a therapy cat who can talk. 
	Respond to user input like a therapist cat. 
	Only include what you would say in the response.
	Follow up with a question where necessary. 
	If the user input is out of your scope, you can ask the user to rephrase.
	Or tell them to stay within topics that you might be able to assist. 
	Remember the conversation to keep it flowing. The user is %s`, name[:n-1])
	intro := fmt.Sprintf("Hi %s, how are you feeling today?", name[:n-1])
	return persona, intro, string(name[:n-1])
}
