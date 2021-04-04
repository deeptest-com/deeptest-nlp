package handler

import (
	"github.com/kataras/iris/v12"
	"github.com/utlai/utl/internal/pkg/utils"
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/service"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
	"strconv"
)

type ProjectCtrl struct {
	BaseCtrl

	ProjectService *service.ProjectService `inject:""`
}

func NewProjectCtrl() *ProjectCtrl {
	return &ProjectCtrl{}
}

func (c *ProjectCtrl) List(ctx iris.Context) {
	keywords := ctx.FormValue("keywords")
	status := ctx.FormValue("status")
	pageNoStr := ctx.FormValue("pageNo")
	pageSizeStr := ctx.FormValue("pageSize")

	pageNo, _ := strconv.Atoi(pageNoStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	if pageSize == 0 {
		pageSize = serverConst.PageSize
	}

	projects, total := c.ProjectService.List(keywords, status, pageNo, pageSize)

	_, _ = ctx.JSON(_utils.ApiResPage(200, "请求成功",
		projects, pageNo, pageSize, total))
}

func (c *ProjectCtrl) Get(ctx iris.Context) {

}

func (c *ProjectCtrl) Create(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	model := new(model.Project)
	if err := ctx.ReadJSON(model); err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), nil))
		return
	}

	if c.Validate(*model, ctx) {
		return
	}

	err := c.ProjectService.Save(model)
	if err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_utils.ApiRes(200, "操作成功", model))
	return
}

func (c *ProjectCtrl) Update(ctx iris.Context) {

}

func (c *ProjectCtrl) SetDefault(ctx iris.Context) {

}

func (c *ProjectCtrl) Delete(ctx iris.Context) {

}
