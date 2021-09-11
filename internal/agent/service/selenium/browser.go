package seleniumOpt

import (
	"fmt"
	"github.com/tebeka/selenium"
	agentConf "github.com/utlai/utl/internal/agent/conf"
	consts "github.com/utlai/utl/internal/comm/const"
	"github.com/utlai/utl/internal/comm/domain"
	_commonUtils "github.com/utlai/utl/internal/pkg/libs/common"
	_fileUtils "github.com/utlai/utl/internal/pkg/libs/file"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	_shellUtils "github.com/utlai/utl/internal/pkg/libs/shell"
	"os"
	"path/filepath"
	"sync"
)

const (
	keySeleniumService = "selenium-service"
	keySeleniumDriver  = "selenium-driver"
)

type SeleniumBrowser struct {
	syncMap sync.Map
}

func NewSeleniumBrowser() *SeleniumBrowser {
	return &SeleniumBrowser{}
}

func (s *SeleniumBrowser) Restart(instruction domain.RasaResp) (result domain.InstructionResp) {
	result = s.Stop()
	if !result.IsSuccess() {
		return
	}
	result = s.Start(instruction)

	return
}

func (s *SeleniumBrowser) Start(instruction domain.RasaResp) (result domain.InstructionResp) {
	driverType := instruction.Entities[1].Value
	driverVersion := instruction.Entities[2].Value
	port := 8848

	seleniumPath, _ := s.DownloadDriver("selenium", "")
	driverPath, _ := s.DownloadDriver(driverType, driverVersion)

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

func (s *SeleniumBrowser) Stop() (result domain.InstructionResp) {
	obj, ok := s.syncMap.Load(keySeleniumDriver)
	if ok && obj != nil {
		driver := obj.(selenium.WebDriver)
		driver.Quit()
	}

	_shellUtils.KillProcess("org.openqa.grid.selenium")

	result.Pass("")
	return
}

func (s *SeleniumBrowser) DownloadDriver(driverType, driverVersion string) (driverPath string, err error) {
	url := ""
	if driverType == "selenium" {
		// https://dl.cnezsoft.com/driver/selenium/driver.exe
		url = fmt.Sprintf("%s%s/driver.jar", consts.DriverDownloadUrl, driverType)

		driverPath = filepath.Join(agentConf.Inst.WorkDir, "driver", "selenium", "driver.jar")
	} else {
		// https://dl.cnezsoft.com/driver/chrome/windows/93/driver.exe
		url = fmt.Sprintf("%s%s/%s/%s/driver", consts.DriverDownloadUrl, driverType, _commonUtils.GetOs(), driverVersion)

		driverPath = filepath.Join(agentConf.Inst.WorkDir, "driver", driverType, driverVersion, "driver")
	}
	if _commonUtils.IsWin() {
		url += ".exe"
		driverPath += ".exe"
	}

	if _fileUtils.FileExist(driverPath) {
		return
	}

	if _commonUtils.IsWin() {
		url += ".exe"
	}

	_fileUtils.RmDir(driverPath)
	err = _fileUtils.Download(url, driverPath)
	if err != nil {
		_logUtils.Errorf("fail to download driver , error %s.", err.Error())
	}

	_shellUtils.ExeShell(fmt.Sprintf("chmod +x %s", driverPath), "")

	return
}

func (s *SeleniumBrowser) GetDriver() (driver selenium.WebDriver) {
	_, ok1 := s.syncMap.Load(keySeleniumService)
	obj, ok2 := s.syncMap.Load(keySeleniumDriver)

	if ok1 && ok2 && obj != nil {
		driver = obj.(selenium.WebDriver)
	}

	return
}
