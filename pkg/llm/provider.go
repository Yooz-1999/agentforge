package llm

import "context"

type Message struct {
	Role    string
	Content string
}

type ChatRequest struct {
	SystemPrompt  string
	ModelProvider string
	ModelName     string
	Temperature   float64
	MaxTokens     int64
	Messages      []Message
}

type ChatResponse struct {
	Content          string
	PromptTokens     int64
	CompletionTokens int64
}

type StreamChunk struct {
	Delta string
}

type StreamHandler func(chunk StreamChunk) error

type Provider interface {
	Chat(ctx context.Context, req *ChatRequest) (*ChatResponse, error)
	Stream(ctx context.Context, req *ChatRequest, handler StreamHandler) error
}
