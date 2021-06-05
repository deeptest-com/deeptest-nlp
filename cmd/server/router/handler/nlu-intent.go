package handler

import (
	"github.com/kataras/iris/v12"
	_httpUtils "github.com/utlai/utl/internal/pkg/libs/http"
	_stringUtils "github.com/utlai/utl/internal/pkg/libs/string"
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/service"
)

type NluIntentCtrl struct {
	BaseCtrl

	IntentService *service.NluIntentService `inject:""`
}

func NewNluIntentCtrl() *NluIntentCtrl {
	return &NluIntentCtrl{}
}

func (c *NluIntentCtrl) List(ctx iris.Context) {
	taskId, _ := ctx.URLParamInt("taskId")

	intents := c.IntentService.ListByTask(uint(taskId))

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "请求成功", intents))
}

func (c *NluIntentCtrl) Get(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	model := c.IntentService.Get(uint(id))
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", model))
	return
}

func (c *NluIntentCtrl) Create(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	data := map[string]string{}
	if err := ctx.ReadJSON(&data); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	po, err := c.IntentService.Create(data["taskId"], data["targetId"], data["mode"], data["name"])
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, "操作失败", nil))
		return
	}

	intents := c.IntentService.ListByTask(po.TaskId)

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", intents))
	return
}

func (c *NluIntentCtrl) Update(ctx iris.Context) {
	model := model.NluIntent{}
	if err := ctx.ReadJSON(&model); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	err := c.IntentService.Update(&model)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, "操作失败", nil))
		return
	}

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", model))
}

func (c *NluIntentCtrl) SetDefault(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	c.IntentService.SetDefault(uint(id))
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", ""))
}

func (c *NluIntentCtrl) Disable(ctx iris.Context) {
	taskId, _ := ctx.URLParamInt("taskId")
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	c.IntentService.Disable(uint(id))
	intents := c.IntentService.ListByTask(uint(taskId))

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", intents))
}

func (c *NluIntentCtrl) Delete(ctx iris.Context) {
	taskId, _ := ctx.URLParamInt("taskId")
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	c.IntentService.Delete(uint(id))
	intents := c.IntentService.ListByTask(uint(taskId))

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", intents))
}

func (c *NluIntentCtrl) Move(ctx iris.Context) {
	data := map[string]string{}
	if err := ctx.ReadJSON(&data); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	taskId := _stringUtils.ParseUint(data["taskId"])
	srcId := _stringUtils.ParseUint(data["srcId"])
	targetId := _stringUtils.ParseUint(data["targetId"])

	_, err := c.IntentService.Move(srcId, targetId, taskId, data["mode"])
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, "操作失败", nil))
		return
	}

	intents := c.IntentService.ListByTask(taskId)

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", intents))
}
