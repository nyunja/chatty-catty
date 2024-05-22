package lib

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
)

func PrintChat(ctx context.Context, client *genai.Client) {
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
	fmt.Println("Jill: ")
	os.Stdout.WriteString(intro + "\n")
	for {
		input := make([]byte, 1024)
		sep := checkOps()
		fmt.Println()
		fmt.Println(name + ":")
		n, _ := os.Stdin.Read(input)
		endCommand := string(input)[:n-len(sep)]
		if toLower(endCommand) == "end" {
			fmt.Println()
			fmt.Println("You have ended the session")
			os.Exit(0)
		}
		fmt.Println()
		resp, err := cs.SendMessage(ctx, genai.Text(string(input)))
		if err != nil {
			log.Fatalln("Error: unable to retrieve response from the client")
			return
		}
		fmt.Println("Jill: ")
		fmt.Println(resp.Candidates[0].Content.Parts[0])
	}
}
