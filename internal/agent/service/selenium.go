package agentService

import (
	seleniumOpt "github.com/utlai/utl/internal/agent/service/selenium"
	consts "github.com/utlai/utl/internal/comm/const"
	"github.com/utlai/utl/internal/comm/domain"
	_domain "github.com/utlai/utl/internal/pkg/domain"
	_i118Utils "github.com/utlai/utl/internal/pkg/libs/i118"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
)

const ()

type SeleniumService struct {
	SeleniumBrowser    *seleniumOpt.SeleniumBrowser    `inject:""`
	SeleniumNavigation *seleniumOpt.SeleniumNavigation `inject:""`
	SeleniumPage       *seleniumOpt.SeleniumPage       `inject:""`
}

func NewSeleniumService() *RegisterService {
	return &RegisterService{}
}

func (s *SeleniumService) Exec(instruction *domain.RasaResp, rpcResult *_domain.RpcResult) {
	instructionResult := domain.InstructionResult{}

	if instruction.Intent == nil || instruction.Intent.Name == "" {
		rpcResult.Pass("no instruction")
		rpcResult.Payload = instructionResult
		return
	}

	cmd := instruction.Intent.Name

	// init driver
	if cmd == consts.SeleniumStart.ToString() {
		instructionResult = s.SeleniumBrowser.Restart(*instruction)
		rpcResult.Pass("")
		rpcResult.Payload = instructionResult
		return
	} else if cmd == consts.SeleniumStop.ToString() {
		instructionResult = s.SeleniumBrowser.Stop()
		rpcResult.Pass("")
		rpcResult.Payload = instructionResult
		return
	}

	//srv := driverCache.(selenium.Service)
	driver := s.SeleniumBrowser.GetDriver()
	if driver == nil {
		rpcResult.Fail("")
		instructionResult.Fail(_i118Utils.Sprintf("pls.start.selenium"))
		rpcResult.Payload = instructionResult
		return
	}

	switch cmd {
	case consts.Load.ToString():
		instructionResult = s.SeleniumNavigation.Load(*instruction, driver)
	case consts.GetSource.ToString():
		instructionResult = s.SeleniumPage.GetPageSource(*instruction, driver)

	default:
		_logUtils.Infof("unknown instruction %s.", cmd)
	}

	rpcResult.Pass("")
	rpcResult.Payload = instructionResult
	return
}
