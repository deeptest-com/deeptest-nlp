package service

import (
	"fmt"
	consts "github.com/utlai/utl/internal/comm/const"
	_fileUtils "github.com/utlai/utl/internal/pkg/libs/file"
	serverConf "github.com/utlai/utl/internal/server/cfg"
	"github.com/utlai/utl/internal/server/domain"
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/repo"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
	"gopkg.in/yaml.v2"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type NluCompileService struct {
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

func NewNluCompileService() *NluCompileService {
	return &NluCompileService{}
}

func (s *NluCompileService) CompileProject(id uint) {
	if serverConf.Config.Analyzer == consts.Rasa {
		s.RasaCompile(id)
	} else if serverConf.Config.Analyzer == consts.Pattern {

	}

	return
}

func (s *NluCompileService) RasaCompile(id uint) {
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
}

func (s *NluCompileService) parserDomain(projectDir string) (nluDomain domain.NluDomain) {
	domainFilePath := filepath.Join(projectDir, "domain.yml")
	content := _fileUtils.ReadFileBuf(domainFilePath)

	yaml.Unmarshal(content, &nluDomain)

	return
}

func (s *NluCompileService) convertIntent(projectId uint, projectDir string, nluDomain *domain.NluDomain) (files []string) {
	_fileUtils.RmDir(filepath.Join(projectDir, "intent"))

	existSlotCodeMap := map[string]bool{}

	tasks := s.NluTaskRepo.ListByProjectId(projectId)
	for _, task := range tasks {
		intents := s.NluIntentRepo.ListByTaskIdNoDisabled(task.ID)

		nluTask := domain.NluTask{Version: serverConst.NluVersion}
		for _, intent := range intents {
			nluDomain.Intents = append(nluDomain.Intents, intent.Name)

			nluIntent := domain.NluIntent{Intent: intent.Name}

			sents := s.NluSentRepo.ListByIntentId(intent.ID)
			for _, sent := range sents {
				slotNameMap := s.getSlotNameMap(sent.ID)

				s.populateSlots(sent.ID, slotNameMap, &existSlotCodeMap, nluDomain)

				intentExamples := s.genIntentSent(sent, slotNameMap)
				nluIntent.Examples += "- " + intentExamples + "\n"
			}

			nluTask.Intents = append(nluTask.Intents, nluIntent)
		}

		yamlContent := changeArrToFlow(nluTask)

		intentFilePath := filepath.Join(projectDir, "data", "intent", task.Name+".yml")
		_fileUtils.WriteFile(intentFilePath, yamlContent)
	}

	return
}

func (s *NluCompileService) convertSynonym(projectId uint, projectDir string, nluDomain *domain.NluDomain) {
	_fileUtils.RmDir(filepath.Join(projectDir, "synonym"))

	synonyms := s.NluSynonymRepo.ListByProjectId(projectId)
	for _, synonym := range synonyms {
		//nluDomain.Entities = append(nluDomain.Entities, synonym.Name)

		nluSynonym := domain.NluSynonym{Version: serverConst.NluVersion}
		synonymDef := domain.NluSynonymDef{Synonym: fmt.Sprintf("%s_%s",
			synonym.Name, serverConst.SlotTypeAbbrMap["synonym"])}

		synonymItems := s.NluSynonymItemRepo.ListBySynonymId(synonym.ID)
		for _, item := range synonymItems {
			synonymDef.Examples += "- " + item.Content + "\n"
		}
		nluSynonym.SynonymDef = append(nluSynonym.SynonymDef, synonymDef)

		yamlContent := changeArrToFlow(nluSynonym)
		filePath := filepath.Join(projectDir, "data", "synonym", synonym.Name+".yml")
		_fileUtils.WriteFile(filePath, yamlContent)
	}

	return
}
func (s *NluCompileService) convertLookup(projectId uint, projectDir string, nluDomain *domain.NluDomain) {
	_fileUtils.RmDir(filepath.Join(projectDir, "lookup"))

	lookups := s.NluLookupRepo.ListByProjectId(projectId)
	for _, lookup := range lookups {
		//nluDomain.Entities = append(nluDomain.Entities, lookup.Name)

		nluLookup := domain.NluLookup{Version: serverConst.NluVersion}
		lookupItem := domain.NluLookupItem{Lookup: fmt.Sprintf("%s_%s",
			lookup.Name, serverConst.SlotTypeAbbrMap["lookup"])}

		lookupItems := s.NluLookupItemRepo.ListByLookupId(lookup.ID)
		for _, item := range lookupItems {
			lookupItem.Examples += "- " + item.Content + "\n"
		}
		nluLookup.Items = append(nluLookup.Items, lookupItem)

		yamlContent := changeArrToFlow(nluLookup)
		filePath := filepath.Join(projectDir, "data", "lookup", lookup.Name+".yml")
		_fileUtils.WriteFile(filePath, yamlContent)
	}

	return
}
func (s *NluCompileService) convertRegex(projectId uint, projectDir string, nluDomain *domain.NluDomain) {
	_fileUtils.RmDir(filepath.Join(projectDir, "regex"))

	regexes := s.NluRegexRepo.ListByProjectId(projectId)
	for _, regex := range regexes {
		//nluDomain.Entities = append(nluDomain.Entities, regex.Name)

		nluRegex := domain.NluRegex{Version: serverConst.NluVersion}
		regexItem := domain.NluRegexItem{Regex: fmt.Sprintf("%s_%s",
			regex.Name, serverConst.SlotTypeAbbrMap["regex"])}

		regexItems := s.NluRegexItemRepo.ListByRegexId(regex.ID)
		for _, item := range regexItems {
			regexItem.Examples += "- " + item.Content + "\n"
		}
		nluRegex.Items = append(nluRegex.Items, regexItem)

		yamlContent := changeArrToFlow(nluRegex)
		filePath := filepath.Join(projectDir, "data", "regex", regex.Name+".yml")
		_fileUtils.WriteFile(filePath, yamlContent)
	}

	return
}

func (s *NluCompileService) getSlotNameMap(sentId uint) (ret map[string]map[string]string) {
	ret = map[string]map[string]string{}

	slots := s.NluSlotRepo.ListBySentId(sentId)
	for _, slot := range slots {
		slotMap := s.getSlotTypeAndVal(slot.Type, slot.Value)
		if slotMap["name"] == "" {
			continue
		}

		ret[fmt.Sprintf("%s-%s", slot.Type, slot.Value)] = slotMap
	}

	return
}

func (s *NluCompileService) getSlotTypeAndVal(tp string, val string) (ret map[string]string) {
	ret = map[string]string{}

	id, _ := strconv.Atoi(val)

	if tp == string(serverConst.Synonym) {
		entity := s.NluSynonymRepo.Get(uint(id))
		ret["code"] = entity.Name
		ret["name"] = entity.Name

	} else if tp == string(serverConst.Lookup) {
		entity := s.NluLookupRepo.Get(uint(id))
		ret["code"] = entity.Name
		ret["name"] = entity.Name

	} else if tp == string(serverConst.Regex) {
		entity := s.NluRegexRepo.Get(uint(id))
		ret["code"] = entity.Name
		ret["name"] = entity.Name
	} else if tp == string(serverConst.Slot) {
		ret["code"] = val
		ret["name"] = val
	}

	return
}

func (s *NluCompileService) populateSlots(sentId uint, slotMap map[string]map[string]string, codeMap *map[string]bool, nluDomain *domain.NluDomain) {
	slots := s.NluSlotRepo.ListBySentId(sentId)
	for _, slot := range slots {
		slotCode := slotMap[fmt.Sprintf("%s-%s", slot.Type, slot.Value)]["code"]
		if slotCode == "" || (*codeMap)[slotCode] {
			continue
		}

		//if slot.Type == "_slot_" && !_stringUtils.StrInArr(slotCode, nluDomain.Entities) {
		nluDomain.Entities = append(nluDomain.Entities, slotCode)
		//}

		slotItem := domain.SlotItem{Type: "text", InfluenceConversation: false}
		mapItem := yaml.MapItem{Key: slotCode, Value: slotItem}
		nluDomain.Slots = append(nluDomain.Slots, mapItem)

		(*codeMap)[slotCode] = true
	}
}

func (s *NluCompileService) genIntentSent(sent model.NluSent, slotMap map[string]map[string]string) (ret string) {
	if strings.Index(sent.Html, "<span") < 0 {
		ret = sent.Html
		return
	}

	regx := regexp.MustCompile(`(?U)<span\b.*>(.*)</span>`)
	spanArr := regx.FindAllString(sent.Html, -1)
	for _, span := range spanArr {
		line := ""

		regx2 := regexp.MustCompile(`\s(\S*)="(\S*)"`)
		arr2 := regx2.FindAllStringSubmatch(span, -1)

		tp := ""
		val := ""
		for _, subArr := range arr2 {
			//all := subArr[0]
			tpTemp := subArr[1]
			valTemp := subArr[2]

			if tpTemp == "data-type" {
				tp = valTemp
			} else if tpTemp == "data-value" {
				val = valTemp
			}
		}

		regx3 := regexp.MustCompile(`>(.*)<`)
		arr3 := regx3.FindAllStringSubmatch(span, -1)
		content := strings.TrimSpace(arr3[0][1])

		if tp == "" || tp == string(serverConst.Text) {
			line += content
		} else {
			slotCode := slotMap[fmt.Sprintf("%s-%s", tp, val)]["code"]
			if tp == "_slot_" {
				line += fmt.Sprintf(`[%s](%s)`, content, val)
			} else {
				line += fmt.Sprintf(`[%s]{"entity":"%s", "value":"%s_%s"}`,
					content, slotCode, slotCode, serverConst.SlotTypeAbbrMap[tp])
			}
		}

		if line != "" {
			ret += line
		}
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
