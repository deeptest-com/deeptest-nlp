package handler

import (
	"github.com/kataras/iris/v12"
	_domain "github.com/utlai/utl/internal/pkg/domain"
	_httpUtils "github.com/utlai/utl/internal/pkg/libs/http"
	"github.com/utlai/utl/internal/server/service"
)

type RpcCtrl struct {
	Ctx        iris.Context
	RpcService *service.RpcService `inject:""`
}

func NewRpcCtrl() *RpcCtrl {
	return &RpcCtrl{}
}

func (c *RpcCtrl) Request(ctx iris.Context) {
	rpcReq := _domain.RpcReq{}
	err := ctx.ReadJSON(&rpcReq)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	var obj interface{}

	rpcResult := c.RpcService.Request(rpcReq.ComputerIp, rpcReq.ComputerPort, rpcReq.ApiPath, rpcReq.ApiMethod, &obj)

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "请求成功", rpcResult))
}
