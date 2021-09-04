package agentService

import (
	"fmt"
	"github.com/tebeka/selenium"
	agentConf "github.com/utlai/utl/internal/agent/conf"
	"github.com/utlai/utl/internal/comm/domain"
	_domain "github.com/utlai/utl/internal/pkg/domain"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"os"
	"path/filepath"
	"sync"
)

const (
	keySelenium = "selenium"
)

type SeleniumService struct {
	syncMap sync.Map
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
	srv, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	s.syncMap.Store(keySelenium, srv)
	if err != nil {
		msg := fmt.Sprintf("start selenium service failed, err %s", err.Error())

		_logUtils.Errorf(msg)
		s.syncMap.Store(keySelenium, srv)

		result.Fail(msg)
		return
	}

	result.Pass("")
	return
}

func (s *SeleniumService) ExecInstruct(instruct *domain.InstructSelenium) (result _domain.RpcResult) {

	result.Pass("")
	return
}
