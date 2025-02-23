package robot

// import (
// 	"encoding/json"

// 	"github.com/valyala/fasthttp"
// )

// func (c *Client) ChatCompletions(request *Completions) (*CompletionResponse, error) {
// 	request.Model = c.config.Model

// 	reqBody, err := json.Marshal(&request)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req := fasthttp.AcquireRequest()

// 	req.SetRequestURI(c.config.Endpoint + "/api/chat")
// 	req.Header.Add("Authorization", `Bearer `+c.config.APIKey)
// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.SetMethod(fasthttp.MethodPost)
// 	req.SetBody(reqBody)

// 	resp := fasthttp.AcquireResponse()

// 	err = c.c.Do(req, resp)
// 	defer fasthttp.ReleaseRequest(req)
// 	defer fasthttp.ReleaseResponse(resp)

// 	if err != nil {
// 		return nil, err
// 	}

// 	var respBody CompletionResponse
// 	if err = json.Unmarshal(resp.Body(), &respBody); err != nil {
// 		return nil, err
// 	}

// 	return &respBody, nil
// }
