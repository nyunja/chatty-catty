package core

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
)

func SendMessage(ctx context.Context, cs *genai.ChatSession, message string) (string, error) {
	resp, err := cs.SendMessage(ctx, genai.Text(message))
	if err != nil {
		return "", fmt.Errorf("unable to recieve response from the client: %v", err)
	}
	if len(resp.Candidates) == 0  || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no response received from the client")
	}
	return fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0]), nil
}
