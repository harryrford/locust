package locust

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
)

var Endpoint = `https://api.x.ai`

var Client = &fasthttp.Client{
	ReadTimeout:  30 * time.Second,
	WriteTimeout: 30 * time.Second,
}

func ChatCompletions(apiKey string, m *Completions) (*CompletionResponse, error) {
	reqBody, err := json.Marshal(&m)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(reqBody))

	req := fasthttp.AcquireRequest()

	req.SetRequestURI(Endpoint + "/v1/chat/completions")
	req.Header.Add("Authorization", `Bearer `+apiKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.SetMethod(fasthttp.MethodPost)
	req.SetBody(reqBody)

	resp := fasthttp.AcquireResponse()

	err = Client.Do(req, resp)
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	if err != nil {
		return nil, err
	}

	var respBody CompletionResponse
	if err = json.Unmarshal(resp.Body(), &respBody); err != nil {
		return nil, err
	}

	return &respBody, nil
}
