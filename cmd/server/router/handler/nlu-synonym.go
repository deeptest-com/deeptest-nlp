package handler

import (
	"github.com/kataras/iris/v12"
	"github.com/utlai/utl/internal/pkg/utils"
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/service"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
	"strconv"
)

type NluSynonymCtrl struct {
	BaseCtrl

	NluSynonymService *service.NluSynonymService `inject:""`
}

func NewNluSynonymCtrl() *NluSynonymCtrl {
	return &NluSynonymCtrl{}
}

func (c *NluSynonymCtrl) List(ctx iris.Context) {
	keywords := ctx.FormValue("keywords")
	pageNoStr := ctx.FormValue("pageNo")
	pageSizeStr := ctx.FormValue("pageSize")

	pageNo, _ := strconv.Atoi(pageNoStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	if pageSize == 0 {
		pageSize = serverConst.PageSize
	}

	plans, total := c.NluSynonymService.List(keywords, pageNo, pageSize)

	_, _ = ctx.JSON(_utils.ApiResPage(200, "请求成功",
		plans, pageNo, pageSize, total))
}

func (c *NluSynonymCtrl) Get(ctx iris.Context) {

}

func (c *NluSynonymCtrl) Create(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	plan := new(model.NluSynonym)
	if err := ctx.ReadJSON(plan); err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), nil))
		return
	}

	if c.Validate(*plan, ctx) {
		return
	}

	err := c.NluSynonymService.Save(plan)
	if err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_utils.ApiRes(200, "操作成功", plan))
	return
}

func (c *NluSynonymCtrl) Update(ctx iris.Context) {

}

func (c *NluSynonymCtrl) Delete(ctx iris.Context) {

}
