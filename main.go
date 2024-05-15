package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"runtime"

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

func toLower(s string) string {
	result := ""
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			result += string(ch)
		} else {
			result += string(ch + 32)
		}
	}
	return result
}

func Cap(s string) string {
	result := ""
	if string(s[0]) >= "a" && string(s[0]) <= "z" {
		result += string(s[0] - 32)
	}
	return result  + s[1:]
}

func printJill() {
	file, err := os.Open("jill.txt")
	if err != nil {
		os.Stdout.WriteString("Error: problem opening file")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		os.Stdout.WriteString(scanner.Text() + "\n")
	}

}

func checkOps() string {
	ops := runtime.GOOS
	sep := ""
	switch ops {
	case "windows":
		sep = "\r\n"
	case "linux":
		sep = "\n"
	}
	return sep
}

func getName() (string, string, string) {
	name := make([]byte, 100)
	sep := checkOps()
	printJill()
	os.Stdout.WriteString("Hi, I am Jill, your therapy cat. What is your name?\n")
	n, _ := os.Stdin.Read(name)
	name2 := Cap(string(name[:n-len(sep)]))
	os.Stdout.WriteString("\r\n")
	persona := `Pretend you are a therapy cat who can talk called Jill. 
	Respond to user input like a therapist cat. 
	Only include what you would say in the response.
	Follow up with a question where necessary. 
	If the user input is out of your scope, you can ask the user to rephrase.
	Or tell them to stay within topics that you might be able to assist. 
	Remember the conversation to keep it flowing. The user is ` + name2
	intro := "Hi " + name2 + ", how are you feeling today?"
	return persona, intro, name2
}
