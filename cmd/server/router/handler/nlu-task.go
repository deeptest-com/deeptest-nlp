package handler

import (
	"github.com/kataras/iris/v12"
	"github.com/utlai/utl/internal/pkg/utils"
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/service"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
)

type NluTaskCtrl struct {
	BaseCtrl

	NluTaskService *service.NluTaskService `inject:""`
}

func NewNluTaskCtrl() *NluTaskCtrl {
	return &NluTaskCtrl{}
}

func (c *NluTaskCtrl) List(ctx iris.Context) {
	keywords := ctx.URLParam("keywords")
	status := ctx.URLParam("status")
	pageNo, _ := ctx.URLParamInt("pageNo")
	pageSize, _ := ctx.URLParamInt("pageSize")

	if pageSize == 0 {
		pageSize = serverConst.PageSize
	}

	tasks, total := c.NluTaskService.List(keywords, status, pageNo, pageSize)

	_, _ = ctx.JSON(_utils.ApiResPage(200, "请求成功",
		tasks, pageNo, pageSize, total))
}

func (c *NluTaskCtrl) Get(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	withIntents, _ := ctx.URLParamBool("withIntents")
	if err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), nil))
		return
	}

	model := c.NluTaskService.Get(uint(id), withIntents)
	_, _ = ctx.JSON(_utils.ApiRes(200, "操作成功", model))
	return
}

func (c *NluTaskCtrl) Create(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	model := model.NluTask{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), nil))
		return
	}

	if c.Validate(model, ctx) {
		return
	}

	err := c.NluTaskService.Save(&model)
	if err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_utils.ApiRes(200, "操作成功", model))
	return
}

func (c *NluTaskCtrl) Update(ctx iris.Context) {
	model := model.NluTask{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), nil))
		return
	}

	err := c.NluTaskService.Update(&model)
	if err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_utils.ApiRes(200, "操作成功", model))
}

func (c *NluTaskCtrl) SetDefault(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), nil))
		return
	}

	c.NluTaskService.SetDefault(uint(id))
	_, _ = ctx.JSON(_utils.ApiRes(200, "操作成功", ""))
}

func (c *NluTaskCtrl) Disable(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), nil))
		return
	}

	c.NluTaskService.Disable(uint(id))
	_, _ = ctx.JSON(_utils.ApiRes(200, "操作成功", ""))
}

func (c *NluTaskCtrl) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), nil))
		return
	}

	c.NluTaskService.Delete(uint(id))
	_, _ = ctx.JSON(_utils.ApiRes(200, "操作成功", ""))
}
