package service

import (
	_fileUtils "github.com/utlai/utl/internal/pkg/libs/file"
	"github.com/utlai/utl/internal/server/domain"
	"github.com/utlai/utl/internal/server/model"
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
	NluRegexItemRepo   *repo.NluRegexItemRepo   `inject:""`

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
			intentItem := domain.NluIntentItem{Intent: intent.Name}

			sents := s.NluSentRepo.ListByIntentId(intent.ID)
			for _, sent := range sents {
				s.populateSlots(sent.ID, nluDomain)

				intentExamples := s.genIntentExamples(sent)
				intentItem.Examples += intentExamples + "\n"
			}

			nluIntent.Items = append(nluIntent.Items, intentItem)

			yamlContent := changeArrToFlow(nluIntent)
			intentFilePath := filepath.Join(projectDir, "data", "intent", intent.Name+".yml")
			_fileUtils.WriteFile(intentFilePath, yamlContent)
		}
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

		synonymItems := s.NluSynonymItemRepo.ListBySynonymId(synonym.ID)
		for _, item := range synonymItems {
			synonymDef.Examples += item.Content + "\n"
		}
		nluSynonym.SynonymDef = append(nluSynonym.SynonymDef, synonymDef)

		yamlContent := changeArrToFlow(nluSynonym)
		filePath := filepath.Join(projectDir, "data", "synonym", synonym.Name+".yml")
		_fileUtils.WriteFile(filePath, yamlContent)
	}

	return
}
func (s *NluConvertService) convertLookup(projectId uint, projectDir string, nluDomain *domain.NluDomain) {
	_fileUtils.RmDir(filepath.Join(projectDir, "lookup"))

	lookups := s.NluLookupRepo.ListByProjectId(projectId)
	for _, lookup := range lookups {
		nluDomain.Entities = append(nluDomain.Entities, lookup.Name)

		nluLookup := domain.NluLookup{Version: serverConst.NluVersion}
		lookupDef := domain.NluLookupItem{Lookup: lookup.Name}

		lookupItems := s.NluLookupItemRepo.ListByLookupId(lookup.ID)
		for _, item := range lookupItems {
			lookupDef.Examples += item.Content + "\n"
		}
		nluLookup.Items = append(nluLookup.Items, lookupDef)

		yamlContent := changeArrToFlow(nluLookup)
		filePath := filepath.Join(projectDir, "data", "lookup", lookup.Name+".yml")
		_fileUtils.WriteFile(filePath, yamlContent)
	}

	return
}
func (s *NluConvertService) convertRegex(projectId uint, projectDir string, nluDomain *domain.NluDomain) {
	_fileUtils.RmDir(filepath.Join(projectDir, "regex"))

	regexes := s.NluRegexRepo.ListByProjectId(projectId)
	for _, regex := range regexes {
		nluDomain.Entities = append(nluDomain.Entities, regex.Name)

		nluRegex := domain.NluRegex{Version: serverConst.NluVersion}

		regexDef := domain.NluRegexItem{Regex: regex.Name}

		regexItems := s.NluRegexItemRepo.ListByRegexId(regex.ID)
		for _, item := range regexItems {
			regexDef.Examples += item.Content + "\n"
		}
		nluRegex.Items = append(nluRegex.Items, regexDef)

		yamlContent := changeArrToFlow(nluRegex)
		filePath := filepath.Join(projectDir, "data", "regex", regex.Name+".yml")
		_fileUtils.WriteFile(filePath, yamlContent)
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

func (s *NluConvertService) populateSlots(sentId uint, nluDomain *domain.NluDomain) {
	slots := s.NluSlotRepo.ListBySentId(sentId)
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

func (s *NluConvertService) genIntentExamples(sent model.NluSent) (ret string) {
	html := sent.Html

	ret = html
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
