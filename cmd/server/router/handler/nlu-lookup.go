package handler

import (
	"github.com/kataras/iris/v12"
	"github.com/utlai/utl/internal/pkg/utils"
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/service"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
	"strconv"
)

type NluLookupCtrl struct {
	BaseCtrl

	LookupService *service.NluLookupService `inject:""`
}

func NewNluLookupCtrl() *NluLookupCtrl {
	return &NluLookupCtrl{}
}

func (c *NluLookupCtrl) List(ctx iris.Context) {
	keywords := ctx.FormValue("keywords")
	status := ctx.FormValue("status")
	pageNoStr := ctx.FormValue("pageNo")
	pageSizeStr := ctx.FormValue("pageSize")

	pageNo, _ := strconv.Atoi(pageNoStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	if pageSize == 0 {
		pageSize = serverConst.PageSize
	}

	lookups, total := c.LookupService.List(keywords, status, pageNo, pageSize)

	_, _ = ctx.JSON(_utils.ApiResPage(200, "请求成功",
		lookups, pageNo, pageSize, total))
}

func (c *NluLookupCtrl) Get(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), nil))
		return
	}

	model := c.LookupService.Get(uint(id))
	_, _ = ctx.JSON(_utils.ApiRes(200, "操作成功", model))
	return
}

func (c *NluLookupCtrl) Create(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	model := model.NluLookup{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), nil))
		return
	}

	if c.Validate(model, ctx) {
		return
	}

	err := c.LookupService.Save(&model)
	if err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_utils.ApiRes(200, "操作成功", model))
	return
}

func (c *NluLookupCtrl) Update(ctx iris.Context) {
	model := model.NluLookup{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), nil))
		return
	}

	err := c.LookupService.Update(&model)
	if err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_utils.ApiRes(200, "操作成功", model))
}

func (c *NluLookupCtrl) SetDefault(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), nil))
		return
	}

	c.LookupService.SetDefault(uint(id))
	_, _ = ctx.JSON(_utils.ApiRes(200, "操作成功", ""))
}

func (c *NluLookupCtrl) Disable(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), nil))
		return
	}

	c.LookupService.Disable(uint(id))
	_, _ = ctx.JSON(_utils.ApiRes(200, "操作成功", ""))
}

func (c *NluLookupCtrl) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), nil))
		return
	}

	c.LookupService.Delete(uint(id))
	_, _ = ctx.JSON(_utils.ApiRes(200, "操作成功", ""))
}
