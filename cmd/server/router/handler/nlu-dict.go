package handler

import (
	"github.com/kataras/iris/v12"
	_httpUtils "github.com/utlai/utl/internal/pkg/libs/http"
	"github.com/utlai/utl/internal/server/service"
)

type NluDictCtrl struct {
	BaseCtrl

	DictService *serverService.NluDictService `inject:""`
}

func NewNluDictCtrl() *NluDictCtrl {
	return &NluDictCtrl{}
}

func (c *NluDictCtrl) List(ctx iris.Context) {
	tp := ctx.URLParam("type")

	list := c.DictService.List(tp)

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "请求成功", list))
}
