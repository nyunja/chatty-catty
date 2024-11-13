package core

import (
	cfg "chatty/config"
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func NewChatClient(ctx context.Context) (*genai.Client, error) {
	if cfg.API_KEY == "" {
		return nil, cfg.ErrNoAPIKey
	}
	client, err := genai.NewClient(ctx, option.WithAPIKey(cfg.API_KEY))
	if err!= nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}
	return client, nil
}
