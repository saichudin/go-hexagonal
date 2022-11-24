package net

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func HTTPRequest(url string, method string, body string, header map[string]string) (output []byte, err error) {
	var httpParam HttpParam

	httpParam.Url = url
	httpParam.Method = method
	httpParam.Header = header
	httpParam.Body = body
	httpParam.Timeout = 30

	res, err := httpParam.HttpDo()
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		return response, fmt.Errorf("error while sending data to external server with status %d", res.StatusCode)
	}

	return response, nil
}
