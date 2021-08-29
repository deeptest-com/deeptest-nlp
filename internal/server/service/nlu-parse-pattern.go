package service

import (
	consts "github.com/utlai/utl/internal/comm/const"
	"github.com/utlai/utl/internal/server/domain"
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/repo"
	serverVari "github.com/utlai/utl/internal/server/utils/var"
	"regexp"
	"time"
)

type NluParsePatternService struct {
	ProjectRepo   *repo.ProjectRepo   `inject:""`
	NluTaskRepo   *repo.NluTaskRepo   `inject:""`
	NluIntentRepo *repo.NluIntentRepo `inject:""`
	NluSentRepo   *repo.NluSentRepo   `inject:""`
	NluRuleRepo   *repo.NluRuleRepo   `inject:""`

	NluSynonymRepo *repo.NluSynonymRepo `inject:""`
	NluLookupRepo  *repo.NluLookupRepo  `inject:""`
	NluRegexRepo   *repo.NluRegexRepo   `inject:""`

	NluSynonymItemRepo *repo.NluSynonymItemRepo `inject:""`
	NluLookupItemRepo  *repo.NluLookupItemRepo  `inject:""`
	NluRegexItemRepo   *repo.NluRegexItemRepo   `inject:""`

	NluPatternService *NluPatternService `inject:""`
}

func NewNluParsePatternService() *NluParsePatternService {
	return &NluParsePatternService{}
}

func (s *NluParsePatternService) Parse(projectId uint, req domain.NluReq) (ret domain.NluResp) {
	ret.Code = -1
	rasaResp := domain.RasaResp{
		Intent: domain.Intent{
			Confidence: 1,
		},
		StartTime: time.Now(),
	}

	text := req.Text
	if serverVari.PatternData[projectId] == nil {
		s.NluPatternService.Reload(projectId)
	}
	tasks := serverVari.PatternData[projectId]

OuterLoop:
	for _, task := range tasks {
		for _, intent := range task.Intents {
			for _, sent := range intent.Sents {
				rgx := regexp.MustCompile(sent.Example)

				arr := rgx.FindStringSubmatch(text)

				if len(arr) > 0 { // matched
					sent := s.NluSentRepo.Get(sent.Id)
					intent := s.NluIntentRepo.Get(intent.Id)

					rasaResp.Intent.ID = int64(intent.ID)
					rasaResp.Intent.Name = intent.Name
					rasaResp.Intent.Sent = domain.Sent{
						ID:   int64(sent.ID),
						Name: sent.Text,
					}

					s.popEntities(sent, rasaResp)

					break OuterLoop
				}
			}
		}
	}

	rasaResp.EndTime = time.Now()
	ret.SetResult(rasaResp)
	ret.Code = 1

	return
}

func (s *NluParsePatternService) popEntities(sent model.NluSent, resp domain.RasaResp) {

	resp.Entities = []domain.Entity{
		{Extractor: consts.Pattern.ToString()},
	}
}
