package domain

type NluReq struct {
	Sent string `json:"sent"`
}

type NluResp struct {
	Sent   string `json:"sent"`
	Code   int    `json:"code"`
	Result string `json:"result"`
}
