package serverCron

import (
	"fmt"
	consts "github.com/utlai/utl/internal/comm/const"
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
	_cronUtils.AddTask(
		"check",
		fmt.Sprintf("@every %ds", consts.WebCheckInterval),
		func() {

		},
	)
}
