package seleniumOpt

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/utlai/utl/internal/agent/service/comm"
	"github.com/utlai/utl/internal/comm/domain"
)

const ()

type SeleniumNavigation struct {
	InstructionService *comm.InstructionService `inject:""`
}

func NewSeleniumNavigation() *SeleniumNavigation {
	return &SeleniumNavigation{}
}

func (s *SeleniumNavigation) Load(instruction domain.RasaResp, driver selenium.WebDriver) (instructionResult domain.InstructionResult) {
	mp := s.InstructionService.Parer(instruction)
	url := mp["expression"].(string)

	err := driver.Get(url)

	if err != nil {
		instructionResult.Fail(err.Error())
	} else {
		instructionResult.Pass(fmt.Sprintf("success to load %s", url))
	}

	return
}
