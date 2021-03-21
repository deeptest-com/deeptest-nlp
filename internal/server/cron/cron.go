package serverCron

import (
	"fmt"
	_const "github.com/utlai/utl/internal/pkg/const"
	_cronUtils "github.com/utlai/utl/internal/pkg/libs/cron"
)

type ServerCron struct {
}

func NewServerCron() *ServerCron {
	inst := &ServerCron{}
	inst.Init()
	return inst
}

func (s *ServerCron) Init() {
	_cronUtils.AddTaskFuc(
		"check",
		fmt.Sprintf("@every %ds", _const.WebCheckQueueInterval),
		func() {

		},
	)
}
