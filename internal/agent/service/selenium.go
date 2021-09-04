package agentService

import (
	"github.com/utlai/utl/internal/comm/domain"
	_domain "github.com/utlai/utl/internal/pkg/domain"
)

type SeleniumService struct {
}

func NewSeleniumService() *CheckService {
	return &CheckService{}
}

func (s *SeleniumService) ExecInstruct(instruct *domain.InstructSelenium) (result _domain.RpcResult) {
	result = _domain.RpcResult{}

	return
}
