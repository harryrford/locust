# DeepSearch Framework

# System Content
```
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
   - Include references to key sources for credibility and traceability.
```
# User Content

## Step 1. Decomposition
```
"Break down the following question into subquestions: [Main Question]. List each subquestion clearly, then prioritize them based on their relevance or logical sequence and categorize them by the type of information required (e.g., factual, analytical, expert opinion)."
```
## Response Format
```json
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
                "enum": ["factual", "analytical", "expert opinion"],
                "description": "Type of information required."
              }
            },
            "required": ["text", "priority", "type"],
            "additionalProperties": false
          },
          "description": "List of subquestions."
        }
      },
      "required": ["subquestions"],
      "additionalProperties": false
    },
    "strict": true
  }
}
```
## Step 2. Criteria
```
"For this subquestion, [Subquestion], determine if further breakdown is necessary based on these criteria: (1) Is it too broad to answer directly? (2) Can it be interpreted in multiple ways? (3) Does it involve multiple distinct aspects? If yes, provide the additional subquestions."
```
## Response Format
```json
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
            "type": "string"
          },
          "description": "List of additional subquestions if breakdown is needed."
        }
      },
      "required": ["subquestion", "further_breakdown_needed"],
      "additionalProperties": false
    },
    "strict": true
  }
}
```
## Step 3.  Scope
```
"Determine the scope of research for [Subquestion or Main Question]. Suggest specific source types (e.g., academic journals, industry reports, expert interviews) based on the question’s nature, define the depth (e.g., number of sources, primary vs. secondary research), and set a time frame (e.g., last 5 years for emerging topics, broader for stable fields)."
```
## Response Format
```json
{
  "type": "json_schema",
  "json_schema": {
    "name": "scope_response",
    "schema": {
      "type": "object",
      "properties": {
        "subquestion": {
          "type": "string",
          "description": "The subquestion or main question."
        },
        "source_types": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Recommended source types."
        },
        "depth": {
          "type": "string",
          "description": "Depth of research (e.g., number of sources)."
        },
        "time_frame": {
          "type": "string",
          "description": "Relevance time frame."
        }
      },
      "required": ["subquestion", "source_types", "depth", "time_frame"],
      "additionalProperties": false
    },
    "strict": true
  }
}
```
## Step 4. Context
```
"How many instances or data points are required for [Subquestion or Main Question]? Base your answer on the question’s context: use 1-2 sources for factual questions, a larger sample for complex or controversial topics, and prioritize diversity over quantity for emerging fields. Consider statistical significance only for quantitative questions."
```
## Response Format
```json
{
  "type": "json_schema",
  "json_schema": {
    "name": "context_response",
    "schema": {
      "type": "object",
      "properties": {
        "subquestion": {
          "type": "string",
          "description": "The subquestion or main question."
        },
        "number_of_instances": {
          "type": "integer",
          "description": "Number of data instances required."
        }
      },
      "required": ["subquestion", "number_of_instances"],
      "additionalProperties": false
    },
    "strict": true
  }
}
```
## Step 5. Synthesis
```
"Compile the answers from these subquestions [list of subquestions] to form an answer to [Main Question]. Use a synthesis method (e.g., thematic analysis for patterns, comparative analysis for contrasts) to integrate findings. Cross-verify for consistency, relevance, and completeness, resolving any contradictions with additional evidence if needed."
```
## Response Format
```json
{
  "type": "json_schema",
  "json_schema": {
    "name": "synthesis_response",
    "schema": {
      "type": "object",
      "properties": {
        "subquestion_answers": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "subquestion": {
                "type": "string",
                "description": "The subquestion."
              },
              "answer": {
                "type": "string",
                "description": "Answer to the subquestion."
              },
              "sources": {
                "type": "array",
                "items": {
                  "type": "string"
                },
                "description": "Sources used for the answer."
              }
            },
            "required": ["subquestion", "answer", "sources"],
            "additionalProperties": false
          },
          "description": "List of subquestion answers."
        }
      },
      "required": ["subquestion_answers"],
      "additionalProperties": false
    },
    "strict": true
  }
}
```
## Step 6. Final Answer
```
"Generate a final, comprehensive answer to [Main Question] based on the synthesized data from all subquestions. Ensure it’s clear, concise, and evidence-backed. Validate the answer’s coherence and completeness (e.g., via peer review or self-check), and include citations or references to key sources."
```
## Response Format
```json
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
      "required": ["final_answer", "references"],
      "additionalProperties": false
    },
    "strict": true
  }
}
```