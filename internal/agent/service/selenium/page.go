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

func (s *SeleniumPage) GetPageSource(rasaRep domain.RasaResp, driver selenium.WebDriver) (result domain.InstructionResp) {
	src, err := driver.PageSource()

	if err != nil {
		result.Fail(err.Error())
	} else {
		result.Pass("")
		result.Payload = src
	}

	return
}
