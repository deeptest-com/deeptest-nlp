package service

import (
	consts "github.com/utlai/utl/internal/comm/const"
	"github.com/utlai/utl/internal/server/domain"
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/repo"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
	serverVari "github.com/utlai/utl/internal/server/utils/var"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type NluParsePatternService struct {
	ProjectRepo   *repo.ProjectRepo   `inject:""`
	NluTaskRepo   *repo.NluTaskRepo   `inject:""`
	NluIntentRepo *repo.NluIntentRepo `inject:""`
	NluSentRepo   *repo.NluSentRepo   `inject:""`
	NluSlotRepo   *repo.NluSlotRepo   `inject:""`
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

				sections := rgx.FindStringSubmatch(text)

				if len(sections) > 0 { // matched
					sent := s.NluSentRepo.Get(sent.Id)
					intent := s.NluIntentRepo.Get(intent.Id)

					rasaResp.Intent.ID = int64(intent.ID)
					rasaResp.Intent.Name = intent.Name
					rasaResp.Intent.Sent = domain.Sent{
						ID:   int64(sent.ID),
						Name: sent.Text,
					}

					s.popEntities(sections, sent, rasaResp)

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

func (s *NluParsePatternService) popEntities(sections []string, sent model.NluSent, resp domain.RasaResp) {
	slotMap := s.getSlotMap(sent.ID)

	slots := s.NluSlotRepo.ListBySentId(sent.ID)

	for _, slot := range slots {
		slotType := slot.Type
		slotId := slot.ID

		if slotType == serverConst.Synonym {
			slot := slotMap[slotId]

		} else if slotType == serverConst.Lookup {
			slot := slotMap[slotId]

		} else if slotType == serverConst.Regex {
			slot := slotMap[slotId]

		} else if slotType == serverConst.Slot {

		} else if slotType == "" {

		}

		entity := domain.Entity{
			Value:     "",
			Start:     0,
			End:       1,
			Extractor: consts.Pattern.ToString(),
		}

		resp.Entities = append(resp.Entities, entity)
	}
}

func (s *NluParsePatternService) getSlotMap(sentId uint) (ret map[uint]model.NluSlot) {
	slots := s.NluSlotRepo.ListBySentId(sentId)
	for _, slot := range slots {
		ret[slot.ID] = slot
	}

	return
}
