package agentHandler

import (
	agentService "github.com/utlai/utl/internal/agent/service"
	"github.com/utlai/utl/internal/comm/domain"
	_domain "github.com/utlai/utl/internal/pkg/domain"
	"golang.org/x/net/context"
)

type SeleniumCtrl struct {
	SeleniumService *agentService.SeleniumService `inject:""`
}

func NewSeleniumCtrl() *SeleniumCtrl {
	return &SeleniumCtrl{}
}

func (c *SeleniumCtrl) Exec(ctx context.Context, instruction *domain.RasaResp, reply *_domain.RpcResult) (err error) {
	c.SeleniumService.Exec(instruction, reply)

	return nil
}
