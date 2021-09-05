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
	keySeleniumService = "selenium-service"
	keySeleniumDriver  = "selenium-driver"
)

type SeleniumService struct {
	syncMap sync.Map

	SeleniumBrowser    *seleniumOpt.SeleniumBrowser    `inject:""`
	SeleniumNavigation *seleniumOpt.SeleniumNavigation `inject:""`
}

func NewSeleniumService() *RegisterService {
	return &RegisterService{}
}

func (s *SeleniumService) Exec(instruction domain.RasaResp) (resp *domain.InstructionResp) {
	if instruction.Intent == nil || instruction.Intent.Name == "" {
		resp.Pass("no instruction")
		return
	}

	cmd := instruction.Intent.Name

	// init driver
	if cmd == consts.SeleniumStart.ToString() {
		s.start(instruction)
		return
	}

	// exec command
	_, ok1 := s.syncMap.Load(keySeleniumService)
	driverCache, ok2 := s.syncMap.Load(keySeleniumDriver)
	if !ok1 || !ok2 {
		msg := "fail to get selenium driver"
		_logUtils.Errorf(msg)
		resp.Fail(msg)
		return
	}

	//srv := driverCache.(selenium.Service)
	driver := driverCache.(selenium.WebDriver)

	switch cmd {
	case consts.SeleniumStop.ToString():
		s.Stop(*instruction.Intent, driver)

	case consts.Load.ToString():
		s.SeleniumNavigation.Load(*instruction.Intent, driver)

	default:
		_logUtils.Infof("unknown instruction %s.", cmd)
	}

	resp.Pass("")
	return
}

func (s *SeleniumService) start(instruction domain.RasaResp) (result _domain.RpcResult) {
	driverType := instruction.Entities[1].Value
	driverVersion := ""
	port := 0

	seleniumPath := filepath.Join(agentConf.Inst.WorkDir, "driver", "selenium", driverVersion)

	driverPath := "" // download if needed

	selenium.SetDebug(true)
	opts := []selenium.ServiceOption{
		//selenium.StartFrameBuffer(),
		selenium.ChromeDriver(driverPath),
		selenium.Output(os.Stderr),
	}

	srv, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	s.syncMap.Store(keySeleniumService, srv)
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

func (s *SeleniumService) Stop(intent domain.Intent, driver selenium.WebDriver) (result _domain.RpcResult) {
	driver.Quit()

	return
}
