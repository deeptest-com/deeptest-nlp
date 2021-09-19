package handler

import (
	"github.com/kataras/iris/v12"
	_httpUtils "github.com/utlai/utl/internal/pkg/libs/http"
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/service"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
)

type NluSynonymCtrl struct {
	CommCtrl       *CommCtrl                        `inject:""`
	SynonymService *serverService.NluSynonymService `inject:""`
}

func NewNluSynonymCtrl() *NluSynonymCtrl {
	return &NluSynonymCtrl{}
}

func (c *NluSynonymCtrl) List(ctx iris.Context) {
	projectId, _ := c.CommCtrl.GetDefaultProject(ctx)

	keywords := ctx.URLParam("keywords")
	status := ctx.URLParam("status")
	pageNo, _ := ctx.URLParamInt("pageNo")
	pageSize, _ := ctx.URLParamInt("pageSize")
	if pageSize == 0 {
		pageSize = serverConst.PageSize
	}

	synonyms, total := c.SynonymService.List(keywords, status, pageNo, pageSize, projectId)

	_, _ = ctx.JSON(_httpUtils.ApiResPage(200, "请求成功",
		synonyms, pageNo, pageSize, total))
}

func (c *NluSynonymCtrl) Get(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	model := c.SynonymService.Get(uint(id))
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", model))
	return
}

func (c *NluSynonymCtrl) Create(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	model := model.NluSynonym{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	err := c.SynonymService.Save(&model)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", model))
	return
}

func (c *NluSynonymCtrl) Update(ctx iris.Context) {
	model := model.NluSynonym{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	err := c.SynonymService.Update(&model)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", model))
}

func (c *NluSynonymCtrl) SetDefault(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	c.SynonymService.SetDefault(uint(id))
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", ""))
}

func (c *NluSynonymCtrl) Disable(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	c.SynonymService.Disable(uint(id))
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", ""))
}

func (c *NluSynonymCtrl) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	c.SynonymService.Delete(uint(id))
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", ""))
}

func (c *NluSynonymCtrl) BatchRemove(ctx iris.Context) {
	ids := make([]int, 0)
	if err := ctx.ReadJSON(&ids); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	err := c.SynonymService.BatchDelete(ids)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", nil))
}

func (c *NluSynonymCtrl) Resort(ctx iris.Context) {
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

	c.SynonymService.Resort(srcId, targetId)
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", nil))
}
