package handler

import (
	"github.com/kataras/iris/v12"
	_httpUtils "github.com/utlai/utl/internal/pkg/libs/http"
	"github.com/utlai/utl/internal/server/domain"
	"github.com/utlai/utl/internal/server/service"
)

type ValidCtrl struct {
	ValidService *serverService.ValidService `inject:""`
}

func NewValidCtrl() *ValidCtrl {
	return &ValidCtrl{}
}

func (c *ValidCtrl) Valid(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	model := serverDomain.ValidRequest{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	result := c.ValidService.Valid(model)

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "请求成功", result))
}
