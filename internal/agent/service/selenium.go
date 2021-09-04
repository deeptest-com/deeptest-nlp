package agentService

import (
	"fmt"
	"github.com/tebeka/selenium"
	agentConf "github.com/utlai/utl/internal/agent/conf"
	seleniumOpt "github.com/utlai/utl/internal/agent/service/selenium"
	consts "github.com/utlai/utl/internal/comm/const"
	"github.com/utlai/utl/internal/comm/domain"
	_domain "github.com/utlai/utl/internal/pkg/domain"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"os"
	"path/filepath"
	"sync"
)

const (
	keySeleniumDriver = "selenium-driver"
)

type SeleniumService struct {
	syncMap sync.Map

	SeleniumBrowser    *seleniumOpt.SeleniumBrowser    `inject:""`
	SeleniumNavigation *seleniumOpt.SeleniumNavigation `inject:""`
}

func NewSeleniumService() *RegisterService {
	return &RegisterService{}
}

func (s *SeleniumService) StartService(driverType, driverVersion string, port int) (result _domain.RpcResult) {
	seleniumPath := filepath.Join(agentConf.Inst.WorkDir, "driver", "selenium")

	driverPath := "" // download if needed

	selenium.SetDebug(true)
	opts := []selenium.ServiceOption{
		//selenium.StartFrameBuffer(),
		selenium.ChromeDriver(driverPath),
		selenium.Output(os.Stderr),
	}
	_, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		msg := fmt.Sprintf("fail to start selenium service, err %s", err.Error())
		_logUtils.Errorf(msg)
		result.Fail(msg)
		return
	}

	caps := selenium.Capabilities{"browserName": driverType}
	driver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))

	s.syncMap.Store(keySeleniumDriver, driver)

	if err != nil {
		msg := fmt.Sprintf("fail to create selenium driver, err %s", err.Error())
		_logUtils.Errorf(msg)
		result.Fail(msg)
		return
	}

	result.Pass("")
	return
}

func (s *SeleniumService) ExecInstruct(instruct *domain.InstructSelenium) (result _domain.RpcResult) {
	driverCache, ok := s.syncMap.Load(keySeleniumDriver)
	if !ok {
		msg := "fail to get selenium driver"
		_logUtils.Errorf(msg)
		result.Fail(msg)
		return
	}

	driver := driverCache.(selenium.WebDriver)

	if instruct.Opt == consts.Navigation {
		s.SeleniumNavigation.Open(instruct, driver)
	}

	result.Pass("")
	return
}
