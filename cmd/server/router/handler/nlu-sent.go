package handler

import (
	"github.com/kataras/iris/v12"
	_httpUtils "github.com/utlai/utl/internal/pkg/libs/http"
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/service"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
)

type NluSentCtrl struct {
	BaseCtrl

	SentService *serverService.NluSentService `inject:""`
}

func NewNluSentCtrl() *NluSentCtrl {
	return &NluSentCtrl{}
}

func (c *NluSentCtrl) List(ctx iris.Context) {
	keywords := ctx.URLParam("keywords")
	status := ctx.URLParam("status")
	pageNo, _ := ctx.URLParamInt("pageNo")
	pageSize, _ := ctx.URLParamInt("pageSize")
	if pageSize == 0 {
		pageSize = serverConst.PageSize
	}

	sents, total := c.SentService.List(keywords, status, pageNo, pageSize)

	_, _ = ctx.JSON(_httpUtils.ApiResPage(200, "请求成功",
		sents, pageNo, pageSize, total))
}

func (c *NluSentCtrl) Get(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	model := c.SentService.Get(uint(id))

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", model))
	return
}

func (c *NluSentCtrl) Create(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	model := model.NluSent{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	if c.Validate(model, ctx) {
		return
	}

	err := c.SentService.Save(&model)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, "操作失败", nil))
		return
	}

	sents := c.SentService.ListByIntent(model.IntentId)
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", sents))
	return
}

func (c *NluSentCtrl) Update(ctx iris.Context) {
	model := model.NluSent{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	err := c.SentService.Update(&model)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, "操作失败", nil))
		return
	}

	sents := c.SentService.ListByIntent(model.IntentId)
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", sents))
}

func (c *NluSentCtrl) Disable(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	intentId, err := ctx.URLParamInt("intentId")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	c.SentService.Disable(uint(id))
	sents := c.SentService.ListByIntent(uint(intentId))
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", sents))
}

func (c *NluSentCtrl) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	intentId, err := ctx.URLParamInt("intentId")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	c.SentService.Delete(uint(id))
	sents := c.SentService.ListByIntent(uint(intentId))
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", sents))
}
