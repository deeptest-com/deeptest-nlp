package service

import (
	consts "github.com/utlai/utl/internal/comm/const"
	serverConf "github.com/utlai/utl/internal/server/cfg"
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
}

func NewNluParseService() *NluParseService {
	return &NluParseService{}
}

func (s *NluParseService) Parse(projectId int, req domain.NluReq) (ret domain.NluResp) {
	if serverConf.Config.Analyzer == consts.Rasa {
		ret = s.NluParseRasaService.Parse(uint(projectId), req)

	} else if serverConf.Config.Analyzer == consts.Pattern {
		ret = s.NluParsePatternService.Parse(uint(projectId), req)
	}

	return
}
