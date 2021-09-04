package handler

import (
	"github.com/kataras/iris/v12"
	"github.com/utlai/utl/internal/comm/domain"
	_httpUtils "github.com/utlai/utl/internal/pkg/libs/http"
	serverService "github.com/utlai/utl/internal/server/service"
)

type AgentCtrl struct {
	BaseCtrl

	AgentService *serverService.AgentService `inject:""`
}

func NewAgentCtrl() *AgentCtrl {
	return &AgentCtrl{}
}

func (c *AgentCtrl) Register(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	to := domain.Agent{}
	if err := ctx.ReadJSON(&to); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(iris.StatusInternalServerError, err.Error(), nil))
		return
	}

	rpcResp := c.AgentService.Register(to)

	_, _ = ctx.JSON(_httpUtils.ApiRes(int64(rpcResp.Code), "操作成功", rpcResp))
	return
}
