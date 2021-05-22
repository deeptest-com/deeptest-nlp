package handler

import (
	"github.com/kataras/iris/v12"
	"github.com/utlai/utl/internal/pkg/utils"
	"github.com/utlai/utl/internal/server/service"
)

type NluConvertCtrl struct {
	NluConvertService *service.NluConvertService `inject:""`
}

func NewNluConvertCtrl() *NluConvertCtrl {
	return &NluConvertCtrl{}
}

func (c *NluConvertCtrl) Convert(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_utils.ApiRes(400, err.Error(), nil))
		return
	}

	_, _ = ctx.JSON(_utils.ApiRes(200, "请求成功", id))
}
