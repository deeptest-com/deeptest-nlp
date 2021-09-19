package handler

import (
	"github.com/kataras/iris/v12"
	_httpUtils "github.com/utlai/utl/internal/pkg/libs/http"
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/service"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
)

type NluLookupItemCtrl struct {
	LookupItemService *serverService.NluLookupItemService `inject:""`
}

func NewNluLookupItemCtrl() *NluLookupItemCtrl {
	return &NluLookupItemCtrl{}
}

func (c *NluLookupItemCtrl) List(ctx iris.Context) {
	lookupId, _ := ctx.URLParamInt("lookupId")
	keywords := ctx.URLParam("keywords")
	status := ctx.URLParam("status")
	pageNo, _ := ctx.URLParamInt("pageNo")
	pageSize, _ := ctx.URLParamInt("pageSize")
	if pageSize == 0 {
		pageSize = serverConst.PageSize
	}

	items, total := c.LookupItemService.List(lookupId, keywords, status, pageNo, pageSize)

	_, _ = ctx.JSON(_httpUtils.ApiResPage(200, "请求成功",
		items, pageNo, pageSize, total))
}

func (c *NluLookupItemCtrl) Get(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	model := c.LookupItemService.Get(uint(id))
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", model))
	return
}

func (c *NluLookupItemCtrl) Create(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	model := model.NluLookupItem{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	err := c.LookupItemService.Save(&model)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", model))
	return
}

func (c *NluLookupItemCtrl) Update(ctx iris.Context) {
	model := model.NluLookupItem{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	err := c.LookupItemService.Update(&model)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", model))
}

func (c *NluLookupItemCtrl) SetDefault(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	c.LookupItemService.SetDefault(uint(id))
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", ""))
}

func (c *NluLookupItemCtrl) Disable(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	c.LookupItemService.Disable(uint(id))
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", ""))
}

func (c *NluLookupItemCtrl) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	c.LookupItemService.Delete(uint(id))
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", ""))
}

func (c *NluLookupItemCtrl) BatchRemove(ctx iris.Context) {
	ids := make([]int, 0)
	if err := ctx.ReadJSON(&ids); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	err := c.LookupItemService.BatchDelete(ids)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", nil))
}

func (c *NluLookupItemCtrl) Resort(ctx iris.Context) {
	parentId, err := ctx.Params().GetInt("parentId")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	srcId, err := ctx.URLParamInt("srcId")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	targetId, err := ctx.URLParamInt("targetId")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	c.LookupItemService.Resort(srcId, targetId, parentId)
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", nil))
}
