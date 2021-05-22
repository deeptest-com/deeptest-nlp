package handler

import (
	"github.com/kataras/iris/v12"
	"github.com/utlai/utl/internal/pkg/utils"
	"github.com/utlai/utl/internal/server/domain"
	"github.com/utlai/utl/internal/server/service"
)

type ValidCtrl struct {
	ValidService *service.ValidService `inject:""`
}

func NewValidCtrl() *ValidCtrl {
	return &ValidCtrl{}
}

func (c *ValidCtrl) Valid(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	model := domain.ValidRequest{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), nil))
		return
	}

	result := c.ValidService.Valid(model)

	_, _ = ctx.JSON(_utils.ApiRes(200, "请求成功", result))
}
