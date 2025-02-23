package locust

// Subquestion represents an individual subquestion in the decomposition step.
type Subquestion struct {
	Text     string `json:"text"`
	Priority int    `json:"priority"`
	Type     string `json:"type"`
}

// DecompositionResponse represents the response format for breaking down the main question into subquestions.
type DecompositionResponse struct {
	Subquestions []Subquestion `json:"subquestions"`
}

// FurtherBreakdownResponse represents the response format for determining if further breakdown of a subquestion is necessary.
type FurtherBreakdownResponse struct {
	Subquestion            string         `json:"subquestion"`
	FurtherBreakdownNeeded bool           `json:"further_breakdown_needed"`
	AdditionalSubquestions []*Subquestion `json:"additional_subquestions,omitempty"`
}

// // ScopeResponse represents the response format for defining the research scope for a subquestion or main question.
// type ScopeResponse struct {
// 	Subquestion string   `json:"subquestion"`
// 	SourceTypes []string `json:"source_types"`
// 	Depth       string   `json:"depth"`
// 	TimeFrame   string   `json:"time_frame"`
// }

// // ContextResponse represents the response format for identifying the number of data instances required.
// type ContextResponse struct {
// 	Subquestion       string `json:"subquestion"`
// 	NumberOfInstances int    `json:"number_of_instances"`
// }

// SubquestionAnswer represents an individual subquestion answer in the synthesis step.
type SubquestionAnswer struct {
	Subquestion string   `json:"subquestion"`
	Answer      string   `json:"answer"`
	Sources     []string `json:"sources"`
}

// // SynthesisResponse represents the response format for synthesizing answers from subquestions.
// type SynthesisResponse struct {
// 	SubquestionAnswers []SubquestionAnswer `json:"subquestion_answers"`
// }

// FinalAnswerResponse represents the response format for providing the final comprehensive answer to the main question.
type FinalAnswerResponse struct {
	FinalAnswer string   `json:"final_answer"`
	References  []string `json:"references"`
}
