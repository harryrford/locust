package locust

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/harryrford/locust/chat"
)

func DeepResearch(client *chat.Client, query string) (string, error) {
	resp, err := client.ChatCompletions(&chat.Completions{
		Messages: []*chat.Message{
			{
				Role:    "system",
				Content: ResearchSystemMessage,
			},
			{
				Role:    "user",
				Content: fmt.Sprintf("Break down the following question into subquestions: %s. Please list each subquestion clearly.", query),
			},
		},
		Model:          "grok-2-latest",
		ResponseFormat: []byte(SubquestionsFormat),
	})
	if err != nil {
		return "", err
	}
	if len(resp.Choices) != 1 {
		return "", fmt.Errorf("server error")
	}

	rootNode := &ResearchNode{
		Depth:        0,
		Question:     query,
		Subquestions: []*ResearchNode{},
	}

	var subquestions Subquestions
	if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &subquestions); err != nil {
		return "", err
	}

	breakdown, err := researchTree(client, rootNode, subquestions.Subquestions)
	if err != nil {
		return "", err
	}

	jsonResp, err := json.Marshal(breakdown)
	if err != nil {
		return "", err
	}

	return string(jsonResp), nil
}

func researchTree(client *chat.Client, node *ResearchNode, questions []string) (*ResearchNode, error) {
	if node.Depth > 3 {
		return node, nil
	}
	for _, subquestion := range questions {
		fmt.Println(subquestion)
		response, err := client.ChatCompletions(&chat.Completions{
			Messages: []*chat.Message{
				{
					Role:    "system",
					Content: ResearchSystemMessage,
				},
				{
					Role:    "user",
					Content: fmt.Sprintf("For this subquestion, %s, is further breakdown necessary? If yes, please provide the additional subquestions.", subquestion),
				},
			},
			Model:          "grok-2-latest",
			ResponseFormat: []byte(FurtherBreakdownFormat),
		})
		if err != nil {
			return nil, err
		}
		if len(response.Choices) != 1 {
			continue
		}

		var breakdown FurtherBreakdown
		if err := json.Unmarshal([]byte(response.Choices[0].Message.Content), &breakdown); err != nil {
			return nil, err
		}

		if breakdown.NeedsBreakdown {
			furtherBreakdownNode := &ResearchNode{
				Depth:        node.Depth + 1,
				Question:     subquestion,
				Subquestions: []*ResearchNode{},
			}

			subNode, err := researchTree(client, furtherBreakdownNode, breakdown.AdditionalSubquestions)
			if err != nil {
				return nil, err
			}

			node.Subquestions = append(node.Subquestions, subNode)
		} else {

			node.Subquestions = append(node.Subquestions, &ResearchNode{
				Depth:        node.Depth + 1,
				Question:     subquestion,
				Subquestions: nil,
			})
		}
	}

	return node, nil
}
