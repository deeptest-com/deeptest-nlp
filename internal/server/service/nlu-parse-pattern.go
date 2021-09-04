package serverService

import (
	"fmt"
	consts "github.com/utlai/utl/internal/comm/const"
	"github.com/utlai/utl/internal/server/domain"
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/repo"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
	serverVari "github.com/utlai/utl/internal/server/utils/var"
	"regexp"
	"strconv"
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

func (s *NluParsePatternService) Parse(projectId uint, req serverDomain.NluReq) (ret serverDomain.NluResp) {
	ret.Code = -1
	rasaResp := serverDomain.RasaRespForPattern{
		Intent: &serverDomain.Intent{
			Confidence: 1,
		},
		StartTime: time.Now(),
	}

	text := req.Text
	if serverVari.PatternData[projectId] == nil {
		s.NluPatternService.Reload(projectId)
	}
	tasks := serverVari.PatternData[projectId]

	found := false

OuterLoop:
	for _, task := range tasks {
		for _, intent := range task.Intents {
			for _, sent := range intent.Sents {
				rgx := regexp.MustCompile("(?i)" + sent.Example)

				indexArr := rgx.FindStringSubmatchIndex(text)
				//contentArr := rgx.FindStringSubmatch(text)
				//_logUtils.Infof("%v", contentArr)

				if len(indexArr) > 0 { // matched
					found = true

					sent := s.NluSentRepo.Get(sent.Id)
					intent := s.NluIntentRepo.Get(intent.Id)

					rasaResp.Intent.ID = int64(intent.ID)
					rasaResp.Intent.Name = intent.Name
					rasaResp.IntentRanking = append(rasaResp.IntentRanking, serverDomain.IntentRanking{
						Name:       intent.Name,
						Confidence: 1,
					})
					rasaResp.Text = text
					rasaResp.Intent.Sent = serverDomain.Sent{
						ID:   int64(sent.ID),
						Name: sent.Text,
					}

					s.popEntities(indexArr, text, sent, &rasaResp)

					break OuterLoop
				}
			}
		}
	}

	rasaResp.EndTime = time.Now()
	if !found {
		rasaResp.Intent.Confidence = 0
		rasaResp.Entities = make([]serverDomain.Entity, 0)
	}

	ret.SetResult(rasaResp)
	ret.Code = 1

	return
}

func (s *NluParsePatternService) popEntities(indexArr []int, text string, sent model.NluSent, resp *serverDomain.RasaRespForPattern) {
	slots := s.NluSlotRepo.ListBySentId(sent.ID)

	index := 2
	for _, item := range slots {
		if item.Type == "" {
			continue
		}

		entity := serverDomain.Entity{Extractor: consts.Pattern.ToString(), ConfidenceEntity: 1}

		if item.Type == serverConst.Synonym {
			synonymId, _ := strconv.Atoi(item.Value)
			synonym := s.NluSynonymRepo.Get(uint(synonymId))
			entity.Entity = fmt.Sprintf("%d-%s-%s", synonym.ID, item.Type, synonym.Name)
			//entity.Value = item.Text

		} else if item.Type == serverConst.Lookup {
			lookupId, _ := strconv.Atoi(item.Value)
			lookup := s.NluLookupRepo.Get(uint(lookupId))
			entity.Entity = fmt.Sprintf("%d-%s-%s", lookup.ID, item.Type, lookup.Name)
			//entity.Value = item.Text

		} else if item.Type == serverConst.Regex {
			regexId, _ := strconv.Atoi(item.Value)
			regex := s.NluRegexRepo.Get(uint(regexId))
			entity.Entity = fmt.Sprintf("%d-%s-%s", regex.ID, item.Type, regex.Name)
			//entity.Value = item.Text

		} else if item.Type == serverConst.Slot {
			entity.Entity = item.Value
			//entity.Value = item.Text

		}

		entity.Start = int64(indexArr[index])
		entity.End = int64(indexArr[index+1])
		entity.Value = text[entity.Start:entity.End]

		resp.Entities = append(resp.Entities, entity)

		index += 2
	}
}

func (s *NluParsePatternService) getSlotMap(sentId uint) (ret map[uint]model.NluSlot) {
	slots := s.NluSlotRepo.ListBySentId(sentId)
	for _, slot := range slots {
		ret[slot.ID] = slot
	}

	return
}
