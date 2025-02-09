package locust

var ResearchSystemMessage = `You are a research agent designed to break down complex questions, determine research scope, gather data, synthesize answers, and provide a final conclusion. Your task involves these steps: 1. Break down questions into subquestions. 2. Determine if further breakdown is necessary. 3. Define research scope. 4. Identify required data instances. 5. Synthesize data. 6. Produce a final answer.`

var SubquestionsFormat = `
{
   "type":"json_schema",
   "json_schema":{
      "name":"subquestions_response",
      "schema":{
         "type":"object",
         "properties":{
            "subquestions":{
               "type":"array",
               "items":{
                  "type":"string"
               },
               "minItems":1,
               "description":"List of subquestions derived from the main question."
            }
         },
         "required":[
            "subquestions"
         ],
         "additionalProperties":false
      },
      "strict":true
   }
}`

var FurtherBreakdownFormat = `
{
   "type":"json_schema",
   "json_schema":{
      "name":"further_breakdown_response",
      "schema":{
         "type":"object",
         "properties":{
            "needs_breakdown":{
               "type":"boolean",
               "description":"Indicates if further breakdown is necessary."
            },
            "additional_subquestions":{
               "type":"array",
               "items":{
                  "type":"string"
               },
               "description":"List of additional subquestions if breakdown is needed."
            }
         },
         "required":[
            "needs_breakdown"
         ],
         "additionalProperties":false
      },
      "strict":true
   }
}`
