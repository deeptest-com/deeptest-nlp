package agentCron

import (
	"fmt"
	checkService "github.com/utlai/utl/internal/agent/service/check"
	_const "github.com/utlai/utl/internal/pkg/const"
	_cronUtils "github.com/utlai/utl/internal/pkg/libs/cron"
)

func Init() {
	_cronUtils.AddTaskFuc(
		"check",
		fmt.Sprintf("@every %ds", _const.AgentCheckDeviceInterval),
		func() {
			checkService.Check()
		},
	)
}
