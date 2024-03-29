package handler

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	_httpUtils "github.com/utlai/utl/internal/pkg/libs/http"
	"github.com/utlai/utl/internal/server/biz/validate"
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/repo"
)

type PermCtrl struct {
	UserRepo *repo.UserRepo `inject:""`
	PermRepo *repo.PermRepo `inject:""`
}

func NewPermCtrl() *PermCtrl {
	return &PermCtrl{}
}

/**
* @api {get} /admin/permissions/:id 根据id获取权限信息
* @apiName 根据id获取权限信息
* @apiGroup Permissions
* @apiVersion 1.0.0
* @apiDescription 根据id获取权限信息
* @apiSampleRequest /admin/permissions/:id
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission
 */
func (c *PermCtrl) GetPermission(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	//id, _ := ctx.Params().GetUint("id")

	perm, err := c.PermRepo.GetPermission(nil)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", c.PermRepo.PermTransform(perm)))
}

/**
* @api {post} /admin/permissions/ 新建权限
* @apiName 新建权限
* @apiGroup Permissions
* @apiVersion 1.0.0
* @apiDescription 新建权限
* @apiSampleRequest /admin/permissions/
* @apiParam {string} name 权限名
* @apiParam {string} display_name
* @apiParam {string} description
* @apiParam {string} level
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func (c *PermCtrl) CreatePermission(ctx iris.Context) {

	ctx.StatusCode(iris.StatusOK)
	perm := new(model.Permission)
	if err := ctx.ReadJSON(perm); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}
	err := validate.Validate.Struct(*perm)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validate.ValidateTrans) {
			if len(e) > 0 {
				_, _ = ctx.JSON(_httpUtils.ApiRes(400, e, nil))
				return
			}
		}
	}

	err = c.PermRepo.CreatePermission(perm)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, fmt.Sprintf("Error create prem: %s", err.Error()), nil))
		return
	}

	if perm.ID == 0 {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, "操作失败", perm))
		return
	}
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", c.PermRepo.PermTransform(perm)))

}

/**
* @api {post} /admin/permissions/:id/update 更新权限
* @apiName 更新权限
* @apiGroup Permissions
* @apiVersion 1.0.0
* @apiDescription 更新权限
* @apiSampleRequest /admin/permissions/:id/update
* @apiParam {string} name 权限名
* @apiParam {string} display_name
* @apiParam {string} description
* @apiParam {string} level
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func (c *PermCtrl) UpdatePermission(ctx iris.Context) {

	ctx.StatusCode(iris.StatusOK)
	aul := new(model.Permission)

	if err := ctx.ReadJSON(aul); err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}
	err := validate.Validate.Struct(*aul)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validate.ValidateTrans) {
			if len(e) > 0 {
				_, _ = ctx.JSON(_httpUtils.ApiRes(400, e, nil))
				return
			}
		}
	}

	id, _ := ctx.Params().GetUint("id")
	err = c.PermRepo.UpdatePermission(id, aul)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, fmt.Sprintf("Error update prem: %s", err.Error()), nil))
		return
	}

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", c.PermRepo.PermTransform(aul)))

}

/**
* @api {delete} /admin/permissions/:id/delete 删除权限
* @apiName 删除权限
* @apiGroup Permissions
* @apiVersion 1.0.0
* @apiDescription 删除权限
* @apiSampleRequest /admin/permissions/:id/delete
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func (c *PermCtrl) DeletePermission(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	id, _ := ctx.Params().GetUint("id")
	err := c.PermRepo.DeletePermissionById(id)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "删除成功", nil))
}

/**
* @api {get} /permissions 获取所有的权限
* @apiName 获取所有的权限
* @apiGroup Permissions
* @apiVersion 1.0.0
* @apiDescription 获取所有的权限
* @apiSampleRequest /permissions
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func (c *PermCtrl) GetAllPermissions(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	permissions, count, err := c.PermRepo.GetAllPermissions(nil)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	transform := c.PermRepo.PermsTransform(permissions)
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "操作成功", map[string]interface{}{"items": transform, "total": count, "limit": "s.Limit"}))

}
