package seleniumOpt

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/utlai/utl/internal/comm/domain"
	"sync"
)

const ()

type SeleniumNavigation struct {
	syncMap sync.Map
}

func NewSeleniumNavigation() *SeleniumNavigation {
	return &SeleniumNavigation{}
}

func (s *SeleniumNavigation) Load(rasaRep domain.RasaResp, driver selenium.WebDriver) (instructionResult domain.InstructionResult) {
	url := rasaRep.Entities[1].Value
	err := driver.Get(url)

	if err != nil {
		instructionResult.Fail(err.Error())
	} else {
		instructionResult.Pass(fmt.Sprintf("success to load %s", url))
	}

	return
}
