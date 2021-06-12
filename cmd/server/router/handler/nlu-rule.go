package handler

import (
	"github.com/kataras/iris/v12"
	_httpUtils "github.com/utlai/utl/internal/pkg/libs/http"
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/service"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
)

type NluRuleCtrl struct {
	BaseCtrl

	RuleService *service.NluRuleService `inject:""`
}

func NewNluRuleCtrl() *NluRuleCtrl {
	return &NluRuleCtrl{}
}

func (c *NluRuleCtrl) List(ctx iris.Context) {
	keywords := ctx.URLParam("keywords")
	status := ctx.URLParam("status")
	pageNo, _ := ctx.URLParamInt("pageNo")
	pageSize, _ := ctx.URLParamInt("pageSize")
	if pageSize == 0 {
		pageSize = serverConst.PageSize
	}

	sents, total := c.RuleService.List(keywords, status, pageNo, pageSize)

	_, _ = ctx.JSON(_httpUtils.ApiResPage(200, "请求成功",
		sents, pageNo, pageSize, total))
}

func (c *NluRuleCtrl) Get(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	model := c.RuleService.Get(uint(id))

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", model))
	return
}

func (c *NluRuleCtrl) Create(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	model := model.NluRule{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	if c.Validate(model, ctx) {
		return
	}

	err := c.RuleService.Save(&model)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, "操作失败", nil))
		return
	}

	sents := c.RuleService.ListByIntent(model.IntentId)
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", sents))
	return
}

func (c *NluRuleCtrl) Update(ctx iris.Context) {
	model := model.NluRule{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	err := c.RuleService.Update(&model)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, "操作失败", nil))
		return
	}

	sents := c.RuleService.ListByIntent(model.IntentId)
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", sents))
}

func (c *NluRuleCtrl) Disable(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	intentId, err := ctx.URLParamInt("intentId")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	c.RuleService.Disable(uint(id))
	rules := c.RuleService.ListByIntent(uint(intentId))
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", rules))
}

func (c *NluRuleCtrl) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	intentId, err := ctx.URLParamInt("intentId")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	c.RuleService.Delete(uint(id))
	sents := c.RuleService.ListByIntent(uint(intentId))
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", sents))
}
