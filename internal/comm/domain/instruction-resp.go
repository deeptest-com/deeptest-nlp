package domain

import (
	_const "github.com/utlai/utl/internal/pkg/const"
	"time"
)

type InstructionResp struct {
	Code    _const.ResultCode `json:"code"`
	Msg     string            `json:"msg"`
	Payload interface{}       `json:"payload"`

	StartTime time.Time `json:"startTime,omitempty"`
	EndTime   time.Time `json:"endTime,omitempty"`
}

func (result *InstructionResp) Pass(msg string) {
	result.Code = _const.ResultSuccess

	if msg == "" {
		msg = "success"
	}
	result.Msg = msg
}

func (result *InstructionResp) Fail(msg string) {
	result.Code = _const.ResultFail
	result.Msg = msg
}

func (result *InstructionResp) IsSuccess() bool {
	return result.Code == _const.ResultSuccess
}
