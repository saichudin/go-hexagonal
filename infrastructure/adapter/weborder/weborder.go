package weborder

import (
	"e-menu-tentakel/core/model"
	port "e-menu-tentakel/core/port/merchant"
	"e-menu-tentakel/interface/api/extl/v1/merchant/response"
	"e-menu-tentakel/utils/config"
	"e-menu-tentakel/utils/net"
	"encoding/json"
	"fmt"
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
