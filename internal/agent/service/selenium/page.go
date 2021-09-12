package seleniumOpt

import (
	"github.com/tebeka/selenium"
	"github.com/utlai/utl/internal/comm/domain"
	"sync"
)

const ()

type SeleniumPage struct {
	syncMap sync.Map
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
