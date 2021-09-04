package vmCron

import (
	"fmt"
	"github.com/kataras/iris/v12"
	agentService "github.com/utlai/utl/internal/agent/service"
	consts "github.com/utlai/utl/internal/comm/const"
	_cronUtils "github.com/utlai/utl/internal/pkg/libs/cron"
	_dateUtils "github.com/utlai/utl/internal/pkg/libs/date"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"sync"
	"time"
)

type CronService struct {
	syncMap      sync.Map
	CheckService *agentService.CheckService `inject:""`
}

func NewAgentCron() *CronService {
	inst := &CronService{}
	inst.Init()
	return inst
}

func (s *CronService) Init() {
	s.syncMap.Store("isRunning", false)
	s.syncMap.Store("lastCompletedTime", int64(0))

	_cronUtils.AddTask(
		"check",
		fmt.Sprintf("@every %ds", consts.AgentCheckInterval),
		func() {
			isRunning, _ := s.syncMap.Load("isRunning")
			lastCompletedTime, _ := s.syncMap.Load("lastCompletedTime")

			if isRunning.(bool) || time.Now().Unix()-lastCompletedTime.(int64) < consts.AgentCheckInterval {
				_logUtils.Infof("skip this iteration " + _dateUtils.DateTimeStr(time.Now()))
				return
			}
			s.syncMap.Store("isRunning", true)

			s.CheckService.Register()

			s.syncMap.Store("isRunning", false)
			s.syncMap.Store("lastCompletedTime", time.Now().Unix())
		},
	)

	iris.RegisterOnInterrupt(func() {
		_cronUtils.Stop()
	})
}
