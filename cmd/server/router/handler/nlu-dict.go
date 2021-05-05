package handler

import (
	"github.com/kataras/iris/v12"
	"github.com/utlai/utl/internal/pkg/utils"
	"github.com/utlai/utl/internal/server/service"
)

type NluDictCtrl struct {
	BaseCtrl

	DictService *service.NluDictService `inject:""`
}

func NewNluDictCtrl() *NluDictCtrl {
	return &NluDictCtrl{}
}

func (c *NluDictCtrl) List(ctx iris.Context) {
	tp := ctx.URLParam("type")

	list := c.DictService.List(tp)

	_, _ = ctx.JSON(_utils.ApiRes(200, "请求成功", list))
}
