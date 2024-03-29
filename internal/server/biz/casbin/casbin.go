package bizCasbin

import (
	"errors"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/fatih/color"
	"github.com/kataras/iris/v12"
	_httpUtils "github.com/utlai/utl/internal/pkg/libs/http"
	"github.com/utlai/utl/internal/server/biz/jwt"
	"github.com/utlai/utl/internal/server/repo"
	"net/http"
)

type CasbinService struct {
	TokenService *jwt.TokenService `inject:""`
	Enforcer     *casbin.Enforcer  `inject:""`
	UserRepo     *repo.UserRepo    `inject:""`
	TokenRepo    *repo.TokenRepo   `inject:""`
}

func NewCasbinService() *CasbinService {
	return &CasbinService{}
}

func (m *CasbinService) Serve(ctx iris.Context) {
	ctx.StatusCode(http.StatusOK)
	value := m.TokenService.Get(ctx)

	credentials, _ := m.TokenService.GetCredentials(value, ctx)

	if credentials == nil {
		ctx.StopExecution()
		_, _ = ctx.JSON(_httpUtils.ApiRes(401, "", nil))
		ctx.StopExecution()
		return
	} else {
		check, err := m.Check(ctx.Request(), credentials.UserId)
		if !check {
			_, _ = ctx.JSON(_httpUtils.ApiRes(403, err.Error(), nil))
			ctx.StopExecution()
			return
		} else {
			ctx.Values().Set("sess", credentials)
		}
	}

	ctx.Next()
}

// Check checks the username, request's method and path and
// returns true if permission grandted otherwise false.
func (c *CasbinService) Check(r *http.Request, userId string) (bool, error) {
	method := r.Method
	path := r.URL.Path
	ok, err := c.Enforcer.Enforce(userId, path, method)
	if err != nil {
		color.Red("验证权限报错：%v;%s-%s-%s", err.Error(), userId, path, method)
		return false, err
	}
	if !ok {
		msg := fmt.Sprintf("你未拥有 %s:%s 操作权限", method, path)
		color.Red(msg)
		return ok, errors.New(msg)
	}
	return ok, nil
}
