package weborder

import (
	"encoding/json"
	"fmt"
	"go-hexagonal/core/model"
	port "go-hexagonal/core/port/merchant"
	"go-hexagonal/interface/api/extl/v1/merchant/response"
	"go-hexagonal/utils/config"
	"go-hexagonal/utils/net"
	"net/http"
)

type WeborderAdapter struct {
	baseUrl string
}

func NewWeborderAdapter() port.WeborderAdapter {
	return &WeborderAdapter{
		baseUrl: config.AppEnv.UrlWeborder,
	}
}

func (adapter *WeborderAdapter) GetDetailWeblink(weblinkUrl string) (*model.WebLinkUri, error) {
	url := fmt.Sprintf("%sweblink/%s/detail", adapter.baseUrl, weblinkUrl)

	header := make(map[string]string)
	header["Content-Type"] = "application/json"

	responseWeborder, err := net.HTTPRequest(url, http.MethodGet, "", header)
	if err != nil && err.Error() != "error while sending data to external server with status 400" {
		return nil, err
	}

	var response response.RespDetailWeblinkUrl
	if responseWeborder != nil {
		err = json.Unmarshal(responseWeborder, &response)
		if err != nil {
			return nil, err
		}
	}

	return response.Data, nil
}
