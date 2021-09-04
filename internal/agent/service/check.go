package agentService

import (
	agentConf "github.com/utlai/utl/internal/agent/conf"
	consts "github.com/utlai/utl/internal/comm/const"
	"github.com/utlai/utl/internal/comm/domain"
	_httpUtils "github.com/utlai/utl/internal/pkg/libs/http"
	_i118Utils "github.com/utlai/utl/internal/pkg/libs/i118"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"time"
)

type CheckService struct {
	TimeStamp int64
}

func NewCheckService() *CheckService {
	s := CheckService{}
	s.TimeStamp = time.Now().Unix()

	return &s
}

func (s *CheckService) Register() {
	agent := domain.Agent{
		MacAddress: agentConf.Inst.MacAddress,
		Ip:         agentConf.Inst.Ip,
		Port:       agentConf.Inst.Port,
		Status:     consts.AgentReady,
	}

	url := _httpUtils.GenUrl(agentConf.Inst.Server, "client/agent/register")
	resp, ok := _httpUtils.Post(url, agent)

	if ok {
		_logUtils.Info(_i118Utils.I118Prt.Sprintf("success_to_register", agentConf.Inst.Server))
	} else {
		_logUtils.Info(_i118Utils.I118Prt.Sprintf("fail_to_register", agentConf.Inst.Server, resp))
	}
}
