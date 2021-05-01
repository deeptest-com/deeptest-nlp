package handler

import (
	"github.com/kataras/iris/v12"
	"github.com/utlai/utl/internal/pkg/utils"
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/service"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
)

type NluIntentCtrl struct {
	BaseCtrl

	IntentService *service.NluIntentService `inject:""`
}

func NewNluIntentCtrl() *NluIntentCtrl {
	return &NluIntentCtrl{}
}

func (c *NluIntentCtrl) List(ctx iris.Context) {
	keywords := ctx.URLParam("keywords")
	status := ctx.URLParam("status")
	pageNo, _ := ctx.URLParamInt("pageNo")
	pageSize, _ := ctx.URLParamInt("pageSize")

	if pageSize == 0 {
		pageSize = serverConst.PageSize
	}

	intents, total := c.IntentService.List(keywords, status, pageNo, pageSize)

	_, _ = ctx.JSON(_utils.ApiResPage(200, "请求成功",
		intents, pageNo, pageSize, total))
}

func (c *NluIntentCtrl) Get(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), nil))
		return
	}

	model := c.IntentService.Get(uint(id))
	_, _ = ctx.JSON(_utils.ApiRes(200, "操作成功", model))
	return
}

func (c *NluIntentCtrl) Create(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	model := model.NluIntent{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), nil))
		return
	}

	if c.Validate(model, ctx) {
		return
	}

	err := c.IntentService.Save(&model)
	if err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_utils.ApiRes(200, "操作成功", model))
	return
}

func (c *NluIntentCtrl) Update(ctx iris.Context) {
	model := model.NluIntent{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), nil))
		return
	}

	err := c.IntentService.Update(&model)
	if err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_utils.ApiRes(200, "操作成功", model))
}

func (c *NluIntentCtrl) SetDefault(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), nil))
		return
	}

	c.IntentService.SetDefault(uint(id))
	_, _ = ctx.JSON(_utils.ApiRes(200, "操作成功", ""))
}

func (c *NluIntentCtrl) Disable(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), nil))
		return
	}

	c.IntentService.Disable(uint(id))
	_, _ = ctx.JSON(_utils.ApiRes(200, "操作成功", ""))
}

func (c *NluIntentCtrl) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), nil))
		return
	}

	c.IntentService.Delete(uint(id))
	_, _ = ctx.JSON(_utils.ApiRes(200, "操作成功", ""))
}

func (c *NluIntentCtrl) BatchRemove(ctx iris.Context) {
	ids := make([]int, 0)
	if err := ctx.ReadJSON(&ids); err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), nil))
		return
	}

	err := c.IntentService.BatchDelete(ids)
	if err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_utils.ApiRes(200, "操作成功", nil))
}
