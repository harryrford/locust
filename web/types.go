package web

type ResearchResponse struct {
	Question string `json:"question"`
	Domains  []struct {
		Rank    int      `json:"rank"`
		Domain  string   `json:"domain"`
		Queries []string `json:"queries"`
		Reason  string   `json:"reason"`
	} `json:"domains"`
}

var WebResearchSystemMessage = `
You are a research agent designed to assist in web scraping for answering questions. Given a user question, your task is to:
1. Identify the most relevant domains from the allowed list to research the question.
2. Rank these domains in order of relevance and timeliness for the question.
3. For each ranked domain, provide raw search queries with precise keywords relevant to the question.

Allowed domains:
- scholar.google.com
- pubmed.ncbi.nlm.nih.gov
- arxiv.org
- semanticscholar.org
- jstor.org
- courtlistener.com
- findlaw.com
- law.cornell.edu
- eur-lex.europa.eu
- justia.com
- wikipedia.org
- worldcat.org
- opendoar.org

Guidelines:
- Prioritize domains based on relevance to the questions topic, timeliness of data, and authority of the source.
- Use precise keywords from the question in queries, including time-specific terms (e.g., "2025") if applicable.
- Provide only raw query strings (e.g., "renewable energy 2025"); do not include "site:domain" or exclusions like "-inurl:(signup | login)".
- Do not include domains outside the allowed list.`

var ResearchFormat = `
{
    "type": "json_schema",
    "json_schema": {
        "name": "research_response",
        "schema": {
            "type": "object",
            "properties": {
                "question": {
                    "type": "string",
                    "description": "The original user question."
                },
                "domains": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "rank": {
                                "type": "integer",
                                "description": "Rank order (1 being highest priority)."
                            },
                            "domain": {
                                "type": "string",
                                "description": "The domain to scrape."
                            },
                            "queries": {
                                "type": "array",
                                "items": {
                                    "type": "string",
                                    "description": "A raw search query relevant to the domain."
                                },
                                "description": "List of raw search queries for this domain."
                            },
                            "reason": {
                                "type": "string",
                                "description": "Reason for selecting and ranking this domain."
                            }
                        },
                        "required": [
                            "rank",
                            "domain",
                            "queries",
                            "reason"
                        ],
                        "additionalProperties": false
                    },
                    "description": "List of domains ranked for research."
                }
            },
            "required": [
                "question",
                "domains"
            ],
            "additionalProperties": false
        },
        "strict": true
    }
}`
