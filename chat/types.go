package chat

import "encoding/json"

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Completions struct {
	Model          string          `json:"model"`
	Stream         bool            `json:"stream"`
	Temperature    int             `json:"temperature"`
	Messages       []*Message      `json:"messages"`
	ResponseFormat json.RawMessage `json:"response_format"`
}

type Choice struct {
	Index   int     `json:"index"`
	Message Message `json:"message"`
}

type CompletionResponse struct {
	Choices []*Choice `json:"choices"`
}
