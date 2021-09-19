package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	_httpUtils "github.com/utlai/utl/internal/pkg/libs/http"
	"github.com/utlai/utl/internal/server/biz/validate"
	serverService "github.com/utlai/utl/internal/server/service"
)

const (
	keySessionId = "SessionId"
	keyProjectId = "ProjectId"
)

type CommCtrl struct {
	Ctx            iris.Context
	session        *sessions.Session
	ProjectService *serverService.ProjectService `inject:""`
}

func (c *CommCtrl) SetDefaultProject(projectId int, ctx iris.Context) (err error) {
	c.GetSession(ctx)
	c.session.Set(keyProjectId, projectId)

	return
}

func (c *CommCtrl) GetDefaultProject(ctx iris.Context) (projectId int, err error) {
	c.GetSession(ctx)

	projectIdInSession := 0
	if c.session.Get(keyProjectId) != nil {
		projectIdInSession = c.session.Get(keyProjectId).(int)
	}
	projectId, err = ctx.URLParamInt(keyProjectId)

	if projectId > 0 && projectId != projectIdInSession {
		c.ProjectService.SetDefault(uint(projectId))
		c.session.Set(keyProjectId, projectId)

	} else {
		projectId = projectIdInSession

	}

	if projectId <= 0 {
		project, err := c.ProjectService.GetDefault()
		if err == nil {
			projectId = int(project.ID)
			c.session.Set(keyProjectId, projectId)
		}
	}

	return
}

func (c *CommCtrl) Validate(s interface{}, ctx iris.Context) bool {
	err := validate.Validate.Struct(s)

	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validate.ValidateTrans) {
			if len(e) > 0 {
				_, _ = ctx.JSON(_httpUtils.ApiRes(400, e, nil))
				return true
			}
		}
	}

	return false
}

func (c *CommCtrl) GetSession(ctx iris.Context) {
	if c.session == nil {
		sess := sessions.New(sessions.Config{
			Cookie: keySessionId,
		})
		c.session = sess.Start(ctx)
	}

	return
}
