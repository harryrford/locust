package locust

import "github.com/goccy/go-json"

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Completions struct {
	Model          string          `json:"model"`
	Stream         bool            `json:"stream"`
	Temperature    int             `json:"temperature"`
	Messages       []Message       `json:"messages"`
	ResponseFormat json.RawMessage `json:"response_format"`
}

type CompletionResponse struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Index   int     `json:"index"`
	Message Message `json:"message"`
}
