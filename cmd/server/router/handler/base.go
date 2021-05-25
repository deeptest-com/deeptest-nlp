package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	_httpUtils "github.com/utlai/utl/internal/pkg/libs/http"
	"github.com/utlai/utl/internal/server/biz/validate"
)

type BaseCtrl struct {
	Ctx iris.Context
}

func (c *BaseCtrl) Validate(s interface{}, ctx iris.Context) bool {
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
