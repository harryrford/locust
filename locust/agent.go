package locust

import (
	"fmt"
	"strings"
)

var SystemMessage = `You are a research agent designed to break down complex questions, determine research scope, gather data, synthesize answers, and provide a final conclusion. Your task involves these steps: 1. Break down questions into subquestions. 2. Determine if further breakdown is necessary. 3. Define research scope. 4. Identify required data instances. 5. Synthesize data. 6. Produce a final answer.`

func NewLocustQuery(apiKey string, queryQuestion string) string {
	resp, err := ChatCompletions(apiKey, &Completions{
		Messages: []Message{
			{
				Role:    "system",
				Content: SystemMessage,
			},
			{
				Role:    "user",
				Content: fmt.Sprintf("Break down the following question into subquestions: %s. Please list each subquestion clearly.", queryQuestion),
			},
		},
		Stream:      false,
		Temperature: 0,
		Model:       "grok-2-latest",
		ResponseFormat: []byte(`{
			"type": "json_schema",
			"json_schema": {
				"name": "subquestions_response",
				"schema": {
					"type": "object",
					"properties": {
						"responses": {
							"type": "array",
							"items": {"type": "string"}
						}
					},
					"required": ["subquestions"],
					"additionalProperties": false
				},
				"strict": true
			}
		}`),
	})
	if err != nil {
		return err.Error()
	}

	var subQuestions strings.Builder
	for _, v := range resp.Choices {
		subQuestions.WriteString(v.Message.Content)
	}

	return subQuestions.String()
}
