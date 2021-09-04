package seleniumOpt

import (
	"github.com/tebeka/selenium"
	"github.com/utlai/utl/internal/comm/domain"
	_domain "github.com/utlai/utl/internal/pkg/domain"
	"sync"
)

const ()

type SeleniumNavigation struct {
	syncMap sync.Map
}

func NewSeleniumNavigation() *SeleniumNavigation {
	return &SeleniumNavigation{}
}

func (s *SeleniumNavigation) Open(*domain.InstructSelenium, selenium.WebDriver) (result _domain.RpcResult) {

	return
}
