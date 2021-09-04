package serverService

import (
	"fmt"
	"github.com/fatih/color"
	_db "github.com/utlai/utl/internal/pkg/db"
	_commonUtils "github.com/utlai/utl/internal/pkg/libs/common"
	"github.com/utlai/utl/internal/server/model"
)

type InitService struct {
	SeederService *SeederService `inject:""`
}

func NewInitService() {
}

func (s *InitService) Init() {
	if !_commonUtils.IsRelease() {
		err := _db.GetInst().DB().AutoMigrate(
			model.Models...,
		)
		if err != nil {
			color.Yellow(fmt.Sprintf("初始化数据表错误 ：%+v", err))
		}

		s.SeederService.AddPerms()
	}
}
