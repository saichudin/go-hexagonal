package response

import "go-hexagonal/core/model"

type RespDetailWeblinkUrl struct {
	Message string            `json:"message"`
	Code    int               `json:"code"`
	Data    *model.WebLinkUri `json:"data"`
}
