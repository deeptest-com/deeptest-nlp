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

	NluLookupService *service.NluLookupService `inject:""`
}

func NewNluLookupCtrl() *NluLookupCtrl {
	return &NluLookupCtrl{}
}

func (c *NluLookupCtrl) List(ctx iris.Context) {
	keywords := ctx.FormValue("keywords")
	pageNoStr := ctx.FormValue("pageNo")
	pageSizeStr := ctx.FormValue("pageSize")

	pageNo, _ := strconv.Atoi(pageNoStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	if pageSize == 0 {
		pageSize = serverConst.PageSize
	}

	plans, total := c.NluLookupService.List(keywords, pageNo, pageSize)

	_, _ = ctx.JSON(_utils.ApiResPage(200, "请求成功",
		plans, pageNo, pageSize, total))
}

func (c *NluLookupCtrl) Get(ctx iris.Context) {

}

func (c *NluLookupCtrl) Create(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	plan := new(model.NluLookup)
	if err := ctx.ReadJSON(plan); err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), nil))
		return
	}

	if c.Validate(*plan, ctx) {
		return
	}

	err := c.NluLookupService.Save(plan)
	if err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_utils.ApiRes(200, "操作成功", plan))
	return
}

func (c *NluLookupCtrl) Update(ctx iris.Context) {

}

func (c *NluLookupCtrl) Delete(ctx iris.Context) {

}
