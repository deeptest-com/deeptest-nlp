package service

import (
	_fileUtils "github.com/utlai/utl/internal/pkg/libs/file"
	"github.com/utlai/utl/internal/server/domain"
	"github.com/utlai/utl/internal/server/repo"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
	"gopkg.in/yaml.v2"
	"path/filepath"
	"strconv"
)

type NluConvertService struct {
	NluSynonymRepo     *repo.NluSynonymRepo     `inject:""`
	NluSynonymItemRepo *repo.NluSynonymItemRepo `inject:""`
	NluLookupRepo      *repo.NluLookupRepo      `inject:""`
	NluLookupItemRepo  *repo.NluLookupItemRepo  `inject:""`
	NluRegexRepo       *repo.NluRegexRepo       `inject:""`

	ProjectRepo   *repo.ProjectRepo   `inject:""`
	NluTaskRepo   *repo.NluTaskRepo   `inject:""`
	NluIntentRepo *repo.NluIntentRepo `inject:""`
	NluSentRepo   *repo.NluSentRepo   `inject:""`
	NluSlotRepo   *repo.NluSlotRepo   `inject:""`
}

func NewNluConvertService() *NluConvertService {
	return &NluConvertService{}
}

func (s *NluConvertService) ConvertProject(id uint) (files []string) {
	project := s.ProjectRepo.Get(id)
	projectDir := project.Path

	nluDomain := s.parserDomain(projectDir)

	nluDomain.Intents = make([]string, 0)
	nluDomain.Entities = make([]string, 0)
	nluDomain.Slots = yaml.MapSlice{}

	s.convertIntent(id, projectDir, &nluDomain)
	s.convertSynonym(id, projectDir, &nluDomain)
	s.convertLookup(id, projectDir, &nluDomain)
	s.convertRegex(id, projectDir, &nluDomain)

	yamlStr := changeArrToFlow(nluDomain)
	_fileUtils.WriteFile(filepath.Join(projectDir, "domain.yml"), yamlStr)

	return
}

func (s *NluConvertService) parserDomain(projectDir string) (nluDomain domain.NluDomain) {
	domainFilePath := filepath.Join(projectDir, "domain.yml")
	content := _fileUtils.ReadFileBuf(domainFilePath)

	yaml.Unmarshal(content, &nluDomain)

	return
}

func (s *NluConvertService) convertIntent(projectId uint, projectDir string, nluDomain *domain.NluDomain) (files []string) {
	_fileUtils.RmDir(filepath.Join(projectDir, "intent"))

	tasks := s.NluTaskRepo.ListByProjectId(projectId)
	for _, task := range tasks {
		intents := s.NluIntentRepo.ListByTaskId(task.ID)

		for _, intent := range intents {
			nluDomain.Intents = append(nluDomain.Intents, intent.Name)

			nluIntent := domain.NluIntent{}

			sents := s.NluSentRepo.ListByIntentId(intent.ID)

			for _, sent := range sents {
				slots := s.NluSlotRepo.ListBySentId(sent.ID)
				for _, slot := range slots {
					slotName := s.getSlotNameByTypeAndId(slot.Type, slot.Value)
					if slotName == "" {
						continue
					}

					slotItem := domain.SlotItem{Type: "text", InfluenceConversation: false}
					mapItem := yaml.MapItem{Key: slotName, Value: slotItem}
					nluDomain.Slots = append(nluDomain.Slots, mapItem)
				}
			}

			intentFilePath := filepath.Join(projectDir, "intent", intent.Name+".yml")
			bytes, _ := yaml.Marshal(&nluIntent)
			_fileUtils.WriteFile(intentFilePath, string(bytes))
		}
	}

	return
}

func (s *NluConvertService) getSlotNameByTypeAndId(tp string, idStr string) (ret string) {
	id, _ := strconv.Atoi(idStr)

	if tp == string(serverConst.Synonym) {
		entity := s.NluSynonymRepo.Get(uint(id))
		ret = entity.Name

	} else if tp == string(serverConst.Lookup) {
		entity := s.NluLookupRepo.Get(uint(id))
		ret = entity.Name

	} else if tp == string(serverConst.Regex) {
		entity := s.NluRegexRepo.Get(uint(id))
		ret = entity.Name
	}

	return
}

func (s *NluConvertService) convertSynonym(projectId uint, projectDir string, nluDomain *domain.NluDomain) {
	_fileUtils.RmDir(filepath.Join(projectDir, "synonym"))

	synonyms := s.NluSynonymRepo.ListByProjectId(projectId)
	for _, synonym := range synonyms {
		nluDomain.Entities = append(nluDomain.Entities, synonym.Name)

		nluSynonym := domain.NluSynonym{Version: serverConst.NluVersion}
		synonymDef := domain.NluSynonymDef{Synonym: synonym.Name}

		synonymItems := s.NluSynonymItemRepo.ListBySynonymId(projectId)
		for _, item := range synonymItems {
			synonymDef.Examples += item.Content + "\n"
		}
		nluSynonym.SynonymDef = append(nluSynonym.SynonymDef, synonymDef)

		filePath := filepath.Join(projectDir, "synonym", synonym.Name+".yml")
		bytes, _ := yaml.Marshal(&nluSynonym)
		_fileUtils.WriteFile(filePath, string(bytes))
	}

	return
}
func (s *NluConvertService) convertLookup(projectId uint, projectDir string, nluDomain *domain.NluDomain) {
	_fileUtils.RmDir(filepath.Join(projectDir, "lookup"))

	lookups := s.NluLookupRepo.ListByProjectId(projectId)
	for _, lookup := range lookups {
		nluDomain.Entities = append(nluDomain.Entities, lookup.Name)

		nluLookup := domain.NluLookup{Version: serverConst.NluVersion}
		lookupDef := domain.NluLookupDef{Lookup: lookup.Name}

		lookupItems := s.NluSynonymItemRepo.ListBySynonymId(projectId)
		for _, item := range lookupItems {
			lookupDef.Examples += item.Content + "\n"
		}
		nluLookup.LookupDef = append(nluLookup.LookupDef, lookupDef)

		filePath := filepath.Join(projectDir, "lookup", lookup.Name+".yml")
		bytes, _ := yaml.Marshal(&nluLookup)
		_fileUtils.WriteFile(filePath, string(bytes))
	}

	return
}
func (s *NluConvertService) convertRegex(projectId uint, projectDir string, nluDomain *domain.NluDomain) {
	_fileUtils.RmDir(filepath.Join(projectDir, "regex"))

	regexes := s.NluRegexRepo.ListByProjectId(projectId)
	for _, regex := range regexes {
		nluDomain.Entities = append(nluDomain.Entities, regex.Name)

		nluRegex := domain.NluRegex{Version: serverConst.NluVersion}

		regexDef := domain.NluRegexDef{Regex: regex.Name}

		lookupItems := s.NluSynonymItemRepo.ListBySynonymId(projectId)
		for _, item := range lookupItems {
			regexDef.Examples += item.Content + "\n"
		}
		nluRegex.RegexDef = append(nluRegex.RegexDef, regexDef)

		filePath := filepath.Join(projectDir, "regex", regex.Name+".yml")
		bytes, _ := yaml.Marshal(&nluRegex)
		_fileUtils.WriteFile(filePath, string(bytes))
	}

	return
}

func changeArrToFlow(obj interface{}) (ret string) {
	bytes, _ := yaml.Marshal(&obj)
	m := yaml.MapSlice{}
	yaml.Unmarshal(bytes, &m)

	d, _ := yaml.Marshal(&m)
	ret = string(d)
	return
}
