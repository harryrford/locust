package web

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly/v2"
	"github.com/harryrford/locust/chat"
)

type Client struct {
	c    *colly.Collector
	chat *chat.Client
}

func NewClient(chat *chat.Client) *Client {
	colly := colly.NewCollector(
		colly.AllowedDomains(
			"scholar.google.com",
			"pubmed.ncbi.nlm.nih.gov",
			"arxiv.org",
			"semanticscholar.org",
			"jstor.org",
			"courtlistener.com",
			"findlaw.com",
			"law.cornell.edu",
			"eur-lex.europa.eu",
			"justia.com",
			"wikipedia.org",
			"worldcat.org",
			"opendoar.org",
		),
		colly.Async(true),
	)
	return &Client{
		c:    colly,
		chat: chat,
	}
}

func (c *Client) Research(question string) (string, error) {
	response, err := c.chat.ChatCompletions(&chat.Completions{
		Messages: []*chat.Message{
			{
				Role:    "system",
				Content: WebResearchSystemMessage,
			},
			{
				Role:    "user",
				Content: question,
			},
		},
		Model:          "grok-2-latest",
		ResponseFormat: []byte(ResearchFormat),
	})
	if err != nil {
		return "", err
	}
	content := response.Choices[0].Message.Content

	var researchResponse ResearchResponse
	if err := json.Unmarshal([]byte(content), &researchResponse); err != nil {
		return "", fmt.Errorf("unmarshal response failed: %v", err)
	}

	b, err := json.Marshal(researchResponse)
	if err != nil {
		return "", fmt.Errorf("marshal response failed: %v", err)
	}

	os.WriteFile("research_response.json", b, os.ModePerm)

	return "", nil
}
