package locust

type ResearchNode struct {
	Depth        int             `json:"depth"`
	Question     string          `json:"question"`
	Subquestions []*ResearchNode `json:"subquestions"`
}

type Subquestions struct {
	Subquestions []string `json:"subquestions"`
}

type FurtherBreakdown struct {
	NeedsBreakdown         bool     `json:"needs_breakdown"`
	AdditionalSubquestions []string `json:"additional_subquestions"`
}
