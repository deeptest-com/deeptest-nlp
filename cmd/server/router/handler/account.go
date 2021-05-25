package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	_httpUtils "github.com/utlai/utl/internal/pkg/libs/http"
	bizConst "github.com/utlai/utl/internal/server/biz/const"
	jwt2 "github.com/utlai/utl/internal/server/biz/jwt"
	"github.com/utlai/utl/internal/server/biz/redis"
	"github.com/utlai/utl/internal/server/biz/validate"
	"github.com/utlai/utl/internal/server/cfg"
	"github.com/utlai/utl/internal/server/domain"
	"github.com/utlai/utl/internal/server/repo"
	"github.com/utlai/utl/internal/server/service"
)

type AccountCtrl struct {
	UserService *service.UserService `inject:""`

	UserRepo  *repo.UserRepo  `inject:""`
	TokenRepo *repo.TokenRepo `inject:""`
	RoleRepo  *repo.RoleRepo  `inject:""`
	PermRepo  *repo.PermRepo  `inject:""`
}

func NewAccountCtrl() *AccountCtrl {
	return &AccountCtrl{}
}

/**
* @api {post} /admin/login 用户登陆
* @apiName 用户登陆
* @apiGroup Users
* @apiVersion 1.0.0
* @apiDescription 用户登陆
* @apiSampleRequest /admin/login
* @apiParam {string} username 用户名
* @apiParam {string} password 密码
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func (c *AccountCtrl) UserLogin(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	aul := new(validate.LoginRequest)

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

	ctx.Application().Logger().Infof("%s 登录系统", aul.Username)

	search := &domain.Search{
		Fields: []*domain.Filed{
			{
				Key:       "username",
				Condition: "=",
				Value:     aul.Username,
			},
		},
	}
	user, err := c.UserRepo.GetUser(search)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}

	response, code, msg := c.UserService.CheckLogin(ctx, user, aul.Password)
	if code != 200 {
		_, _ = ctx.JSON(_httpUtils.ApiRes(code, msg, response))
		return
	}
	response.RememberMe = aul.RememberMe

	refreshToken := ""
	if code == 200 && aul.RememberMe {
		refreshToken = response.Token
	}

	c.UserService.UpdateRefreshToken(user.ID, refreshToken)

	_, _ = ctx.JSON(_httpUtils.ApiRes(code, msg, response))
}

func (c *AccountCtrl) UserLogout(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	value := ctx.Values().Get("jwt").(*jwt.Token)

	var (
		credentials *bizConst.UserCredentials
		err         error
	)
	if serverConf.Config.Redis.Enable {
		conn := redisUtils.GetRedisClusterClient()
		defer conn.Close()

		credentials, err = c.TokenRepo.GetRedisSession(conn, value.Raw)
		if err != nil {
			_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
			return
		}
		if credentials != nil {
			if err := c.TokenRepo.DelUserTokenCache(conn, *credentials, value.Raw); err != nil {
				_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
				return
			}
		}
	} else {
		credentials = jwt2.GetCredentials(ctx)
		if credentials == nil {
			_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
			return
		} else {
			jwt2.RemoveCredentials(ctx)
		}
	}

	ctx.Application().Logger().Infof("%d 退出系统", credentials.UserId)
	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "退出", nil))
}

func (c *AccountCtrl) UserExpire(ctx iris.Context) {

	ctx.StatusCode(iris.StatusOK)
	value := ctx.Values().Get("jwt").(*jwt.Token)
	conn := redisUtils.GetRedisClusterClient()
	defer conn.Close()
	sess, err := c.TokenRepo.GetRedisSession(conn, value.Raw)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
		return
	}
	if sess != nil {
		if err := c.TokenRepo.UpdateUserTokenCacheExpire(conn, *sess, value.Raw); err != nil {
			_, _ = ctx.JSON(_httpUtils.ApiRes(400, err.Error(), nil))
			return
		}
	}

	_, _ = ctx.JSON(_httpUtils.ApiRes(200, "", nil))
}
