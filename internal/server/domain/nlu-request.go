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

	RasaResult *domain.RasaResp        `json:"nluResult,omitempty"`
	ExecResult *domain.InstructionResp `json:"execResult,omitempty"`
	Msg        *map[string]string      `json:"msg,omitempty"`
}

func (resp *NluResp) SetResult(result domain.RasaResp) {
	resp.RasaResult = &result
}

func (resp *NluResp) SetMsg(msg map[string]string) {
	resp.Msg = &msg
}
