package locust

import (
	"fmt"
	"os"
	"strings"

	"github.com/goccy/go-json"
	"github.com/harryrford/locust/chat"
)

// ResearchNode represents a node in the research tree.
type ResearchNode struct {
	Depth        int
	Question     string
	Answer       string
	Sources      []string
	Subquestions []*ResearchNode
}

// DeepResearch implements the full DeepSearch framework.
func DeepResearch(client *chat.Client, query string) (string, error) {

	// Step 1: Decomposition - Break down the main question into subquestions
	resp, err := client.ChatCompletions(&chat.Completions{
		Messages: []*chat.Message{
			{
				Role:    "system",
				Content: ResearchSystemMessage,
			},
			{
				Role:    "user",
				Content: fmt.Sprintf(`Break down the following question into subquestions: %s. List each subquestion clearly, then prioritize them based on their relevance or logical sequence and categorize them by the type of information required (e.g., factual, analytical, expert opinion).`, query),
			},
		},
		Model:          "grok-2-latest",
		ResponseFormat: []byte(DecompositionFormat),
	})
	if err != nil {
		return "", fmt.Errorf("decomposition failed: %v", err)
	}
	if len(resp.Choices) != 1 {
		return "", fmt.Errorf("server error: expected 1 choice, got %d", len(resp.Choices))
	}

	var subquestions DecompositionResponse

	if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &subquestions); err != nil {
		return "", fmt.Errorf("unmarshal subquestions failed: %v", err)
	}

	questionTexts := make([]string, len(subquestions.Subquestions))
	for i, sq := range subquestions.Subquestions {
		questionTexts[i] = sq.Text
	}

	rootNode := &ResearchNode{
		Depth:        0,
		Question:     query,
		Subquestions: []*ResearchNode{},
	}

	researchNode, err := researchTree(client, rootNode, questionTexts)
	if err != nil {
		return "", fmt.Errorf("research tree failed: %v", err)
	}

	b, err := json.MarshalIndent(researchNode, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	os.WriteFile("research_tree.json", b, os.ModePerm)

	return "", nil
}

func researchTree(client *chat.Client, node *ResearchNode, questions []string) (*ResearchNode, error) {
	if node.Depth > 3 {
		return node, nil
	}

	for _, subquestion := range questions {

		response, err := client.ChatCompletions(&chat.Completions{
			Messages: []*chat.Message{
				{
					Role:    "system",
					Content: ResearchSystemMessage,
				},
				{
					Role:    "user",
					Content: fmt.Sprintf(`For this subquestion, "%s", determine if further breakdown is necessary based on these criteria: (1) Is it too broad to answer directly? (2) Can it be interpreted in multiple ways? (3) Does it involve multiple distinct aspects? If yes, provide the additional subquestions.`, subquestion),
				},
			},
			ResponseFormat: []byte(FurtherBreakdownFormat),
		})
		if err != nil {
			return nil, err
		}

		if len(response.Choices) != 1 {
			return &ResearchNode{}, fmt.Errorf("server error: expected 1 choice, got %d", len(response.Choices))
		}

		var breakdown FurtherBreakdownResponse
		if err := json.Unmarshal([]byte(response.Choices[0].Message.Content), &breakdown); err != nil {
			return nil, err
		}

		if breakdown.FurtherBreakdownNeeded {
			furtherBreakdownNode := &ResearchNode{
				Depth:        node.Depth + 1,
				Question:     subquestion,
				Subquestions: []*ResearchNode{},
			}

			additionalSubquestions := make([]string, len(breakdown.AdditionalSubquestions))
			for i, additionSubquestion := range breakdown.AdditionalSubquestions {
				additionalSubquestions[i] = additionSubquestion.Text
			}

			subNode, err := researchTree(client, furtherBreakdownNode, additionalSubquestions)
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

// GetFinalAnswer synthesizes leaf answers into the final answer.
func GetFinalAnswer(client *chat.Client, mainQuestion string, leafAnswers []SubquestionAnswer) (string, error) {

	var message strings.Builder

	message.WriteString(fmt.Sprintf(`Generate a final, comprehensive answer to "%s" based on the synthesized data from all subquestions. `, mainQuestion))
	message.WriteString("Ensure its clear, concise, and evidence-backed. Validate the answers coherence and completeness (e.g., via peer review or self-check), and include citations or references to key sources.")
	message.WriteString("\n\nSubquestion Answers:")

	for n, answer := range leafAnswers {
		message.WriteString(fmt.Sprintf("\n%d. Subquestion: %s\n   Answer: %s\n   Sources: %s", n+1, answer.Subquestion, answer.Answer, strings.Join(answer.Sources, ", ")))
	}

	resp, err := client.ChatCompletions(&chat.Completions{
		Messages: []*chat.Message{
			{
				Role:    "system",
				Content: ResearchSystemMessage,
			},
			{
				Role:    "user",
				Content: message.String(),
			},
		},
		ResponseFormat: []byte(FinalAnswerFormat),
	})
	if err != nil {
		return "", err
	}
	if len(resp.Choices) != 1 {
		return "", fmt.Errorf("no final answer response")
	}

	var finalAnswer FinalAnswerResponse
	if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &finalAnswer); err != nil {
		return "", err
	}

	var answer strings.Builder

	answer.WriteString(finalAnswer.FinalAnswer)

	answer.WriteString("\n\nSources:")
	for _, source := range finalAnswer.References {
		answer.WriteString("\n- " + source)
	}

	return answer.String(), nil
}
