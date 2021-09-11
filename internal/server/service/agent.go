package serverService

import (
	"fmt"
	"github.com/utlai/utl/internal/comm/domain"
	_domain "github.com/utlai/utl/internal/pkg/domain"
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/repo"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
)

type AgentService struct {
	AgentRepo *repo.AgentRepo `inject:""`

	AgentService     *AgentService     `inject:""`
	WebSocketService *WebSocketService `inject:""`
}

func NewAgentService() *AgentService {
	return &AgentService{}
}

func (s AgentService) Register(agent domain.Agent) (result _domain.RpcResult) {
	err := s.AgentRepo.Register(agent)
	if err != nil {
		result.Fail(fmt.Sprintf("fail to register agent %s ", agent.Ip))
		return
	}

	result.Pass("")

	agents := s.AgentService.List()
	data := map[string]interface{}{"agents": agents, "action": serverConst.UpdateAgent}
	s.WebSocketService.Broadcast(serverConst.WsNamespace, serverConst.WsDefaultRoom, serverConst.WsEvent, data)

	return
}

func (s AgentService) List() (pos []model.Agent) {
	pos = s.AgentRepo.Query()
	return
}
