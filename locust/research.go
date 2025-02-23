package locust

var ResearchSystemMessage = `
You are a research agent designed to break down complex questions, determine research scope, gather data, synthesize answers, and provide a final conclusion. Your task involves the following steps:

1. Break Down Questions into Subquestions
   - Identify the key components or themes of the main question (e.g., who, what, why, how).  
   - List subquestions that are specific, manageable, and collectively address the full scope of the original question.  
   - Prioritize them based on logical sequence (e.g., foundational questions first) or relevance to the main query.  

2. Determine if Further Breakdown is Necessary
   - Assess each subquestion using these criteria:  
     - Is it too broad or vague to answer directly?  
     - Does it have multiple distinct aspects?  
     - Could it be interpreted in different ways?  
   - If yes, decompose it further into more focused subquestions.

3. Define Research Scope
   - For each subquestion, identify the type of sources needed (e.g., academic papers for technical topics, news articles for current events).  
   - Specify the depth of research (e.g., 2-3 sources for simple facts, 5-10 for complex analyses).  
   - Set a time frame for relevance (e.g., last 5 years for fast-changing fields, broader for established topics).

4. Identify Required Data Instances
   - Define "data instances" as specific pieces of evidence (e.g., a study, statistic, or expert quote).  
   - For straightforward subquestions, aim for 1-2 reliable sources; for nuanced or debated topics, target 3-5 diverse instances.  
   - Prioritize quality and variety over sheer quantity.

5. Synthesize Data
   - Combine findings by identifying patterns, trends, or contradictions across subquestions.  
   - Cross-check data for consistency and relevance to the main question.  
   - Resolve discrepancies with additional evidence if needed.

6. Produce a Final Answer
   - Deliver a clear, concise answer that fully addresses the original question, supported by the synthesized data.  
   - Structure it logically (e.g., introduction, key findings, conclusion).  
   - Include references to key sources for credibility and traceability.`

var DecompositionFormat = `
{
    "type": "json_schema",
    "json_schema": {
        "name": "decomposition_response",
        "schema": {
            "type": "object",
            "properties": {
                "subquestions": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "text": {
                                "type": "string",
                                "description": "The subquestion text."
                            },
                            "priority": {
                                "type": "integer",
                                "description": "Priority order (1 being highest)."
                            },
                            "type": {
                                "type": "string",
                                "enum": [
                                    "factual",
                                    "analytical",
                                    "expert opinion"
                                ],
                                "description": "Type of information required."
                            }
                        },
                        "required": [
                            "text",
                            "priority",
                            "type"
                        ],
                        "additionalProperties": false
                    },
                    "description": "List of subquestions."
                }
            },
            "required": [
                "subquestions"
            ],
            "additionalProperties": false
        },
        "strict": true
    }
}`

var FurtherBreakdownFormat = `
{
    "type": "json_schema",
    "json_schema": {
        "name": "further_breakdown_response",
        "schema": {
            "type": "object",
            "properties": {
                "subquestion": {
                    "type": "string",
                    "description": "The original subquestion."
                },
                "further_breakdown_needed": {
                    "type": "boolean",
                    "description": "Indicates if further breakdown is necessary."
                },
                "additional_subquestions": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "text": {
                                "type": "string",
                                "description": "The subquestion text."
                            },
                            "priority": {
                                "type": "integer",
                                "description": "Priority order (1 being highest)."
                            },
                            "type": {
                                "type": "string",
                                "enum": [
                                    "factual",
                                    "analytical",
                                    "expert opinion"
                                ],
                                "description": "Type of information required."
                            }
                        },
                        "required": [
                            "text",
                            "priority",
                            "type"
                        ],
                        "additionalProperties": false
                    },
                    "description": "List of additional subquestions if breakdown is needed."
                }
            },
            "required": [
                "subquestion",
                "further_breakdown_needed"
            ],
            "additionalProperties": false
        },
        "strict": true
    }
}`

var FinalAnswerFormat = `
{
    "type": "json_schema",
    "json_schema": {
        "name": "final_answer_response",
        "schema": {
            "type": "object",
            "properties": {
                "final_answer": {
                    "type": "string",
                    "description": "The comprehensive answer to the main question."
                },
                "references": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "description": "List of key sources used."
                }
            },
            "required": [
                "final_answer",
                "references"
            ],
            "additionalProperties": false
        },
        "strict": true
    }
}`
