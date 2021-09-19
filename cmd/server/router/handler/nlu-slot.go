package handler

import (
	"github.com/kataras/iris/v12"
	_httpUtils "github.com/utlai/utl/internal/pkg/libs/http"
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/service"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
)

type NluSlotCtrl struct {
	SlotService *serverService.NluSlotService `inject:""`
}

func NewNluSlotCtrl() *NluSlotCtrl {
	return &NluSlotCtrl{}
}

func (c *NluSlotCtrl) List(ctx iris.Context) {
	keywords := ctx.URLParam("keywords")
	status := ctx.URLParam("status")
	pageNo, _ := ctx.URLParamInt("pageNo")
	pageSize, _ := ctx.URLParamInt("pageSize")
	if pageSize == 0 {
		pageSize = serverConst.PageSize
	}

	slots, total := c.SlotService.List(keywords, status, pageNo, pageSize)

	_, _ = ctx.JSON(_httpUtils.ApiResPage(200, "请求成功",
		slots, pageNo, pageSize, total))
}

func (c *NluSlotCtrl) Get(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	model := c.SlotService.Get(uint(id))
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", model))
	return
}

func (c *NluSlotCtrl) Create(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	model := model.NluSlot{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	err := c.SlotService.Save(&model)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", model))
	return
}

func (c *NluSlotCtrl) Update(ctx iris.Context) {
	model := model.NluSlot{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	err := c.SlotService.Update(&model)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", model))
}

func (c *NluSlotCtrl) SetDefault(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	c.SlotService.SetDefault(uint(id))
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", ""))
}

func (c *NluSlotCtrl) Disable(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	c.SlotService.Disable(uint(id))
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", ""))
}

func (c *NluSlotCtrl) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	c.SlotService.Delete(uint(id))
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", ""))
}

func (c *NluSlotCtrl) BatchRemove(ctx iris.Context) {
	ids := make([]int, 0)
	if err := ctx.ReadJSON(&ids); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	err := c.SlotService.BatchDelete(ids)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", nil))
}
