package handler

import (
	"github.com/kataras/iris/v12"
	_httpUtils "github.com/utlai/utl/internal/pkg/libs/http"
	_stringUtils "github.com/utlai/utl/internal/pkg/libs/string"
	"github.com/utlai/utl/internal/server/biz/jwt"
	"github.com/utlai/utl/internal/server/service"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
)

type NluRasaCtrl struct {
	NluCompileService  *service.NluCompileService  `inject:""`
	NluTrainingService *service.NluTrainingService `inject:""`
	NluServiceService  *service.NluServiceService  `inject:""`
	NluPatternService  *service.NluPatternService  `inject:""`

	NluHistoryService *service.NluHistoryService `inject:""`
}

func NewNluRasaCtrl() *NluRasaCtrl {
	return &NluRasaCtrl{}
}

func (c *NluRasaCtrl) Compile(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	c.NluCompileService.CompileProject(uint(id))

	cred := jwt.GetCredentials(ctx)
	c.NluHistoryService.Add(_stringUtils.ParseUint(cred.UserId), uint(id), serverConst.Compile)

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "请求成功", nil))
}

func (c *NluRasaCtrl) Training(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	project := c.NluTrainingService.TrainingProject(uint(id))

	cred := jwt.GetCredentials(ctx)
	c.NluHistoryService.Add(_stringUtils.ParseUint(cred.UserId), uint(id), serverConst.StartTraining)

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "请求成功", project))
}

func (c *NluRasaCtrl) ReloadRes(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	c.NluPatternService.Reload(uint(id))

	cred := jwt.GetCredentials(ctx)
	c.NluHistoryService.Add(_stringUtils.ParseUint(cred.UserId), uint(id), serverConst.StartTraining)

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "请求成功", nil))
}

func (c *NluRasaCtrl) Start(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	project, _ := c.NluServiceService.ReStart(uint(id))

	cred := jwt.GetCredentials(ctx)
	c.NluHistoryService.Add(_stringUtils.ParseUint(cred.UserId), uint(id), serverConst.StartService)

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "请求成功", project))
}

func (c *NluRasaCtrl) Stop(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	project := c.NluServiceService.Get(uint(id))
	project, _ = c.NluServiceService.Stop(project)

	cred := jwt.GetCredentials(ctx)
	c.NluHistoryService.Add(_stringUtils.ParseUint(cred.UserId), uint(id), serverConst.StartService)

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "请求成功", project))
}
