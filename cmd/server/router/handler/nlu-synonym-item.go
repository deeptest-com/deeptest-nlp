package handler

import (
	"github.com/kataras/iris/v12"
	_httpUtils "github.com/utlai/utl/internal/pkg/libs/http"
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/service"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
)

type NluSynonymItemCtrl struct {
	SynonymItemService *serverService.NluSynonymItemService `inject:""`
}

func NewNluSynonymItemCtrl() *NluSynonymItemCtrl {
	return &NluSynonymItemCtrl{}
}

func (c *NluSynonymItemCtrl) List(ctx iris.Context) {
	synonymId, _ := ctx.URLParamInt("synonymId")
	keywords := ctx.URLParam("keywords")
	status := ctx.URLParam("status")
	pageNo, _ := ctx.URLParamInt("pageNo")
	pageSize, _ := ctx.URLParamInt("pageSize")
	if pageSize == 0 {
		pageSize = serverConst.PageSize
	}

	items, total := c.SynonymItemService.List(synonymId, keywords, status, pageNo, pageSize)

	_, _ = ctx.JSON(_httpUtils.ApiResPage(200, "请求成功",
		items, pageNo, pageSize, total))
}

func (c *NluSynonymItemCtrl) Get(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	model := c.SynonymItemService.Get(uint(id))
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", model))
	return
}

func (c *NluSynonymItemCtrl) Create(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	model := model.NluSynonymItem{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	err := c.SynonymItemService.Save(&model)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", model))
	return
}

func (c *NluSynonymItemCtrl) Update(ctx iris.Context) {
	model := model.NluSynonymItem{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	err := c.SynonymItemService.Update(&model)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", model))
}

func (c *NluSynonymItemCtrl) SetDefault(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	c.SynonymItemService.SetDefault(uint(id))
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", ""))
}

func (c *NluSynonymItemCtrl) Disable(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	c.SynonymItemService.Disable(uint(id))
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", ""))
}

func (c *NluSynonymItemCtrl) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	c.SynonymItemService.Delete(uint(id))
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", ""))
}

func (c *NluSynonymItemCtrl) BatchRemove(ctx iris.Context) {
	ids := make([]int, 0)
	if err := ctx.ReadJSON(&ids); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	err := c.SynonymItemService.BatchDelete(ids)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", nil))
}

func (c *NluSynonymItemCtrl) Resort(ctx iris.Context) {
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

	c.SynonymItemService.Resort(srcId, targetId, parentId)
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", nil))
}
