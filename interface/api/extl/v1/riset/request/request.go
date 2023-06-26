package request

type MpayCustReq struct {
	Page int `json:"page"`
	Limit int `json:"limit"`
}