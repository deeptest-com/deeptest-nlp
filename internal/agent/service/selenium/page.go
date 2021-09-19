package seleniumOpt

import (
	"github.com/tebeka/selenium"
	"github.com/utlai/utl/internal/agent/service/comm"
	"github.com/utlai/utl/internal/comm/domain"
)

const ()

type SeleniumPage struct {
	InstructionService *comm.InstructionService `inject:""`
}

func NewSeleniumPage() *SeleniumPage {
	return &SeleniumPage{}
}

func (s *SeleniumPage) GetPageSource(rasaRep domain.RasaResp, driver selenium.WebDriver) (instructionResult domain.InstructionResult) {
	src, err := driver.PageSource()

	if err != nil {
		instructionResult.Fail(err.Error())
	} else {
		instructionResult.Pass("")
		instructionResult.Payload = src
	}

	return
}
