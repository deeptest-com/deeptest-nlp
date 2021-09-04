package serverService

import (
	"fmt"
	"github.com/utlai/utl/internal/comm/domain"
	_domain "github.com/utlai/utl/internal/pkg/domain"
	"github.com/utlai/utl/internal/server/repo"
)

type AgentService struct {
	AgentRepo *repo.AgentRepo `inject:""`

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

	return
}
