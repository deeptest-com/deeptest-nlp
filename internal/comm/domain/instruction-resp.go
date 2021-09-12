package domain

import (
	_const "github.com/utlai/utl/internal/pkg/const"
	"time"
)

type InstructionResult struct {
	Code    _const.ResultCode `json:"code"`
	Msg     string            `json:"msg"`
	Payload interface{}       `json:"payload"`

	StartTime time.Time `json:"startTime,omitempty"`
	EndTime   time.Time `json:"endTime,omitempty"`
}

func (result *InstructionResult) Pass(msg string) {
	result.Code = _const.ResultSuccess

	if msg == "" {
		msg = "success"
	}
	result.Msg = msg
}

func (result *InstructionResult) Fail(msg string) {
	result.Code = _const.ResultFail
	result.Msg = msg
}

func (result *InstructionResult) IsSuccess() bool {
	return result.Code == _const.ResultSuccess
}
