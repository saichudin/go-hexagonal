package response

import "e-menu-tentakel/core/model"

type RespDetailWeblinkUrl struct {
	Message string            `json:"message"`
	Code    int               `json:"code"`
	Data    *model.WebLinkUri `json:"data"`
}
