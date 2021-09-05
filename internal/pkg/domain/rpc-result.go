package _domain

import _const "github.com/utlai/utl/internal/pkg/const"

type RpcResult struct {
	Code    _const.ResultCode `json:"code"`
	Msg     string            `json:"msg"`
	Payload interface{}       `json:"payload"`
}

func (result *RpcResult) Pass(msg string) {
	result.Code = _const.ResultSuccess
	result.Msg = msg
}

func (result *RpcResult) Fail(msg string) {
	result.Code = _const.ResultFail
	result.Msg = msg
}

func (result *RpcResult) IsSuccess() bool {
	return result.Code == _const.ResultSuccess
}
