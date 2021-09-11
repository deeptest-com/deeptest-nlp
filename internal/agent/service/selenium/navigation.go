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

func (s *SeleniumNavigation) Load(rasaRep domain.RasaResp, driver selenium.WebDriver) (result domain.InstructionResp) {
	url := rasaRep.Entities[1].Value
	err := driver.Get(url)

	if err != nil {
		result.Fail(err.Error())
	} else {
		result.Pass(fmt.Sprintf("success to load %s", url))
	}

	return
}
