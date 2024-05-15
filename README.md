## README: Therapy Cat Chatbot

This program is a text-based chatbot that simulates a conversation with a therapy cat named Jill. It utilizes the Generative AI Go library ([https://ai.google/discover/generativeai/](https://ai.google/discover/generativeai/)) to interact with a large language model (LLM) called Gemini-Pro.

**How it Works:**

1. **Initialization:**
    * The program loads environment variables from a `.env` file (refer to Environment Variables section for details).
    * It connects to the Generative AI Go API using the provided API key.
    * The program sets up the conversation history by defining Jill's persona and introductory message based on the user's name.

2. **User Interaction:**
    * The program prompts the user for their name and stores it.
    * It then initiates a chat loop:
        * The program displays a prompt for the user to enter their message. 
        * The user's input is sent to the LLM as text. 
        * The program retrieves the LLM's response and displays it as Jill's message.

**Environment Variables:**

The program relies on an environment variable called `API_KEY` to authenticate with the Generative AI Go API. You'll need to create a `.env` file in the same directory as the program and define this variable:

```
API_KEY=<your_api_key_here>
```

**Running the Program:**

1. Make sure you have the required libraries installed (`go get github.com/google/generative-ai-go github.com/joho/godotenv`).
2. Set up your API key in the `.env` file.
3. Run the program using `go run main.go`.

**Note:**

This program is a basic example and doesn't implement any advanced features of therapy. It's for demonstration purposes only.
