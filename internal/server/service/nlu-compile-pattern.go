package service

import (
	consts "github.com/utlai/utl/internal/comm/const"
	_fileUtils "github.com/utlai/utl/internal/pkg/libs/file"
	"github.com/utlai/utl/internal/server/domain"
	"github.com/utlai/utl/internal/server/repo"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
	"path/filepath"
	"strconv"
	"strings"
)

type NluCompilePatternService struct {
	NluSynonymRepo     *repo.NluSynonymRepo     `inject:""`
	NluSynonymItemRepo *repo.NluSynonymItemRepo `inject:""`
	NluLookupRepo      *repo.NluLookupRepo      `inject:""`
	NluLookupItemRepo  *repo.NluLookupItemRepo  `inject:""`
	NluRegexRepo       *repo.NluRegexRepo       `inject:""`
	NluRegexItemRepo   *repo.NluRegexItemRepo   `inject:""`

	ProjectRepo   *repo.ProjectRepo   `inject:""`
	NluTaskRepo   *repo.NluTaskRepo   `inject:""`
	NluIntentRepo *repo.NluIntentRepo `inject:""`
	NluSentRepo   *repo.NluSentRepo   `inject:""`
	NluSlotRepo   *repo.NluSlotRepo   `inject:""`

	NluCompileService *NluCompileService `inject:""`
}

func NewNluCompilePatternService() *NluCompilePatternService {
	return &NluCompilePatternService{}
}

func (s *NluCompilePatternService) PatternCompile(id uint) {
	project := s.ProjectRepo.Get(id)
	projectDir := project.Path

	synonymMap := s.getSynonymMap(id)
	lookupMap := s.getLookupMap(id)
	regexMap := s.getRegexMap(id)

	tasks := s.NluTaskRepo.ListByProjectId(id)
	for _, task := range tasks {
		nluTask := domain.NluPatternTask{Version: serverConst.NluVersion, Name: task.Name}

		intents := s.NluIntentRepo.ListByTaskIdNoDisabled(task.ID)
		for _, intent := range intents {
			lines := make([]string, 0)
			lineMap := map[string]bool{}

			sents := s.NluSentRepo.ListByIntentId(intent.ID)
			for _, sent := range sents {
				slots := s.NluSlotRepo.ListBySentId(sent.ID)

				line := ""
				for _, slot := range slots {
					slotType := slot.Type
					slotVale := slot.Value
					slotText := slot.Text

					if slotType == serverConst.Synonym {
						synonymId, _ := strconv.Atoi(slotVale)
						items := synonymMap[uint(synonymId)]

						section := strings.Join(items, "|")
						line += "(" + section + ")"

					} else if slotType == serverConst.Lookup {
						lookupId, _ := strconv.Atoi(slotVale)
						items := lookupMap[uint(lookupId)]

						section := strings.Join(items, "|")
						line += "(" + section + ")"

					} else if slotType == serverConst.Regex {
						regexId, _ := strconv.Atoi(slotVale)
						items := regexMap[uint(regexId)]

						section := strings.Join(items, "|")
						line += "(" + section + ")"

					} else if slotType == serverConst.Slot {
						line += "((?U:.+))"

					} else if slotType == "" {
						line += slotText

					}
				}

				line = strings.TrimSpace(line)
				if line != "" && !lineMap[line] {
					lines = append(lines, line)
					lineMap[line] = true
				}
			}

			pattern := domain.NluPattern{Id: intent.ID, Name: intent.Name}
			pattern.Examples = strings.Join(lines, "\n")
			nluTask.Intents = append(nluTask.Intents, pattern)
		}

		yamlContent := s.NluCompileService.ChangeArrToFlow(nluTask)

		patternFilePath := filepath.Join(projectDir, consts.Pattern.ToString(), task.Name+".yml")
		_fileUtils.WriteFile(patternFilePath, yamlContent)
	}
}

func (s *NluCompilePatternService) getSynonymMap(projectId uint) (ret map[uint][]string) {
	ret = map[uint][]string{}

	synonyms := s.NluSynonymRepo.ListByProjectId(projectId)
	for _, synonym := range synonyms {
		synonymItems := s.NluSynonymItemRepo.ListBySynonymId(synonym.ID)
		for _, item := range synonymItems {
			ret[synonym.ID] = append(ret[synonym.ID], item.Content)
		}
	}

	return
}
func (s *NluCompilePatternService) getLookupMap(projectId uint) (ret map[uint][]string) {
	ret = map[uint][]string{}

	lookups := s.NluLookupRepo.ListByProjectId(projectId)
	for _, lookup := range lookups {
		lookupItems := s.NluLookupItemRepo.ListByLookupId(lookup.ID)
		for _, item := range lookupItems {
			ret[lookup.ID] = append(ret[lookup.ID], item.Content)
		}
	}

	return
}
func (s *NluCompilePatternService) getRegexMap(projectId uint) (ret map[uint][]string) {
	ret = map[uint][]string{}

	regexes := s.NluRegexRepo.ListByProjectId(projectId)
	for _, regex := range regexes {
		regexItems := s.NluRegexItemRepo.ListByRegexId(regex.ID)
		for _, item := range regexItems {
			ret[regex.ID] = append(ret[regex.ID], item.Content)
		}
	}

	return
}
