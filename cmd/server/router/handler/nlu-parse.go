package handler

import (
	"github.com/kataras/iris/v12"
	_httpUtils "github.com/utlai/utl/internal/pkg/libs/http"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/utlai/utl/internal/server/domain"
	"github.com/utlai/utl/internal/server/service"
)

type NluParseCtrl struct {
	NluParseService *service.NluParseService `inject:""`
}

func NewNluParseCtrl() *NluParseCtrl {
	return &NluParseCtrl{}
}

func (c *NluParseCtrl) Parse(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	projectId, _ := ctx.Params().GetInt("id")
	req := domain.NluReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}
	_logUtils.Infof("request: %d, %v", projectId, req)

	resp := c.NluParseService.Parse(projectId, req)

	_logUtils.Infof("response: %v", resp)

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "请求成功", resp))
}
