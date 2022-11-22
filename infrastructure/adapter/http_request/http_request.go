package http_request

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type HttpRequestClient struct {
}

func (httpReq *HttpRequestClient) NewRequestClient(method string, url string, auth string, authType string, payload interface{}) (response *http.Response, err error) {
	var payloadEncode []byte
	if payload != nil {
		payloadEncode, err = json.Marshal(payload)
		if err != nil {
			return nil, err
		}
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(payloadEncode))
	request.Header.Add("Content-type", "application/json")
	if err != nil {
		return nil, err
	}

	if authType == "Bearer" || authType == "Basic" {
		request.Header.Add("Authorization", authType+" "+auth)
	}

	client := &http.Client{}
	response, err = client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
