package handler

import (
	"github.com/kataras/iris/v12"
	_httpUtils "github.com/utlai/utl/internal/pkg/libs/http"
	"github.com/utlai/utl/internal/server/service"
)

type NluTrainingCtrl struct {
	NluTrainingService *service.NluTrainingService `inject:""`
}

func NewNluTrainingCtrl() *NluTrainingCtrl {
	return &NluTrainingCtrl{}
}

func (c *NluTrainingCtrl) Training(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	c.NluTrainingService.TrainingProject(uint(id))

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "请求成功", nil))
}
