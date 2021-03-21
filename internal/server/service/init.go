package service

import (
	"fmt"
	"github.com/fatih/color"
	middlewareUtils "github.com/utlai/utl/internal/server/biz/middleware/misc"
	"github.com/utlai/utl/internal/server/db"
	"github.com/utlai/utl/internal/server/model"
)

type InitService struct {
}

func NewInitService() {
}

func (s *InitService) Init() {
	err := db.GetInst().DB().AutoMigrate(
		model.Models...,
	)
	if err != nil {
		color.Yellow(fmt.Sprintf("初始化数据表错误 ：%+v", err))
	}

	err = db.GetInst().DB().AutoMigrate(
		&middlewareUtils.CasbinRule{},
	)
}
