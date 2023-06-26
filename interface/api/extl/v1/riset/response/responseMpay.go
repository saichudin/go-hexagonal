package response

import (
	"go-hexagonal/core/model/mpay"
)

type RespDetailWeblinkUrl struct {
	Message string            	`json:"message"`
	Code    int               	`json:"code"`
	Data    *mpay.MpayCustomer 	`json:"data"`
}