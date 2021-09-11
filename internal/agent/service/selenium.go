package agentService

import (
	seleniumOpt "github.com/utlai/utl/internal/agent/service/selenium"
	consts "github.com/utlai/utl/internal/comm/const"
	"github.com/utlai/utl/internal/comm/domain"
	_domain "github.com/utlai/utl/internal/pkg/domain"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
)

const ()

type SeleniumService struct {
	SeleniumBrowser    *seleniumOpt.SeleniumBrowser    `inject:""`
	SeleniumNavigation *seleniumOpt.SeleniumNavigation `inject:""`
}

func NewSeleniumService() *RegisterService {
	return &RegisterService{}
}

func (s *SeleniumService) Exec(instruction *domain.RasaResp, reply *_domain.RpcResult) {
	instructionResp := domain.InstructionResp{}

	if instruction.Intent == nil || instruction.Intent.Name == "" {
		reply.Pass("no instruction")
		reply.Payload = instructionResp
		return
	}

	cmd := instruction.Intent.Name

	// init driver
	if cmd == consts.SeleniumStart.ToString() {
		instructionResp = s.SeleniumBrowser.Restart(*instruction)
		reply.Pass("")
		reply.Payload = instructionResp
		return
	} else if cmd == consts.SeleniumStop.ToString() {
		instructionResp = s.SeleniumBrowser.Stop()
		reply.Pass("")
		reply.Payload = instructionResp
		return
	}

	//srv := driverCache.(selenium.Service)
	driver := s.SeleniumBrowser.GetDriver()

	switch cmd {
	case consts.Load.ToString():
		instructionResp = s.SeleniumNavigation.Load(*instruction, driver)

	default:
		_logUtils.Infof("unknown instruction %s.", cmd)
	}

	reply.Pass("")
	reply.Payload = instructionResp
	return
}
