package serverDomain

import (
	"github.com/utlai/utl/internal/comm/domain"
	_const "github.com/utlai/utl/internal/pkg/const"
)

type NluReq struct {
	AgentId    int    `json:"agentId"`
	Text       string `json:"text"`
	TextOrigin string `json:"textOrigin"`
}

type NluResp struct {
	Text string            `json:"text"`
	Code _const.ResultCode `json:"code"`

	Result *domain.RasaResp   `json:"result,omitempty"`
	Msg    *map[string]string `json:"msg,omitempty"`
}

func (resp *NluResp) SetResult(result domain.RasaResp) {
	resp.Result = &result
}

func (resp *NluResp) SetMsg(msg map[string]string) {
	resp.Msg = &msg
}
