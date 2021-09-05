package serverService

import (
	consts "github.com/utlai/utl/internal/comm/const"
	"github.com/utlai/utl/internal/comm/domain"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	serverConf "github.com/utlai/utl/internal/server/conf"
	"github.com/utlai/utl/internal/server/domain"
	"github.com/utlai/utl/internal/server/repo"
	"sync"
)

var (
	ProjectChanged sync.Map
	ProjectRules   sync.Map
)

type NluParseService struct {
	ProjectRepo   *repo.ProjectRepo   `inject:""`
	AgentRepo     *repo.AgentRepo     `inject:""`
	NluTaskRepo   *repo.NluTaskRepo   `inject:""`
	NluIntentRepo *repo.NluIntentRepo `inject:""`
	NluRuleRepo   *repo.NluRuleRepo   `inject:""`

	NluSynonymRepo *repo.NluSynonymRepo `inject:""`
	NluLookupRepo  *repo.NluLookupRepo  `inject:""`
	NluRegexRepo   *repo.NluRegexRepo   `inject:""`

	NluSynonymItemRepo *repo.NluSynonymItemRepo `inject:""`
	NluLookupItemRepo  *repo.NluLookupItemRepo  `inject:""`
	NluRegexItemRepo   *repo.NluRegexItemRepo   `inject:""`

	NluParseRasaService    *NluParseRasaService    `inject:""`
	NluParsePatternService *NluParsePatternService `inject:""`

	RpcService *RpcService `inject:""`
}

func NewNluParseService() *NluParseService {
	return &NluParseService{}
}

func (s *NluParseService) Parse(projectId int, req serverDomain.NluReq) (
	nluResp serverDomain.NluResp) {
	if serverConf.Inst.Analyzer == consts.Rasa {
		nluResp = s.NluParseRasaService.Parse(uint(projectId), req)

	} else if serverConf.Inst.Analyzer == consts.Pattern {
		nluResp = s.NluParsePatternService.Parse(uint(projectId), req)

	}

	if req.AgentId == 0 || nluResp.RasaResult == nil {
		return
	}

	agent := s.AgentRepo.Get(uint(req.AgentId))
	_logUtils.Infof("exec instruction on agent %s", agent.Ip)

	rpcResp := s.RpcService.ExecInstruction(nluResp, agent)
	instructionResp := rpcResp.Payload.(domain.InstructionResp)

	nluResp.ExecResult = &instructionResp

	return
}
