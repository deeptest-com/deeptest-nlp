package testService

import (
	appiumService "github.com/utlai/utl/internal/agent/service/appium"
	seleniumService "github.com/utlai/utl/internal/agent/service/selenium"
	_const "github.com/utlai/utl/internal/pkg/const"
	_domain "github.com/utlai/utl/internal/pkg/domain"
)

func Exec(build _domain.BuildTo) {
	StartTask()

	if build.BuildType == _const.SeleniumTest {
		seleniumService.ExecTest(&build)
	} else if build.BuildType == _const.AppiumTest {
		appiumService.ExecTest(&build)
	}

	RemoveTask()
	EndTask()
}
