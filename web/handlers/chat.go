package handlers

import (
	cfg "chatty/config"
	"chatty/core"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/generative-ai-go/genai"
)

type ChatResponse struct {
	Response string `json:"response"`
}

type ChatRequest struct {
	Message string `json:"message"`
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	var chatReq ChatRequest
	if err := json.NewDecoder(r.Body).Decode(&chatReq); err != nil {
		http.Error(w, "Opps! Invalid request", http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	client, err := core.NewChatClient(ctx)
	if err != nil {
		http.Error(w, "unable to create chat client", http.StatusInternalServerError)
		return
	}
	model := client.GenerativeModel(cfg.ModelName)
	cs := model.StartChat()
	cs.History = append(cs.History, &genai.Content{
		Parts: []genai.Part{genai.Text(chatReq.Message)},
		Role: "user",
	})
	core.InitializeChat(cs, cfg.Persona, cfg.DefaultIntro)
	resp, err := core.SendMessage(ctx, cs, chatReq.Message)
	if err != nil {
		http.Error(w, "Error in processing request", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	cs.History = append(cs.History, &genai.Content{
		Parts: []genai.Part{genai.Text(chatReq.Message)},
		Role: "model",
	})
	res := ChatResponse{Response: resp}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
