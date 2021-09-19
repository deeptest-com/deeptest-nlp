package serverService

import (
	"fmt"
	consts "github.com/utlai/utl/internal/comm/const"
	_fileUtils "github.com/utlai/utl/internal/pkg/libs/file"
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

type NluCompileRasaService struct {
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

func NewNluCompileRasaService() *NluCompileRasaService {
	return &NluCompileRasaService{}
}

func (s *NluCompileRasaService) RasaCompile(id uint) {
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

	yamlStr := s.NluCompileService.ChangeArrToFlow(nluDomain)
	_fileUtils.WriteFile(filepath.Join(projectDir, consts.Rasa.ToString(), "domain.yml"), yamlStr)
}

func (s *NluCompileRasaService) parserDomain(projectDir string) (nluDomain serverDomain.NluDomain) {
	domainFilePath := filepath.Join(projectDir, consts.Rasa.ToString(), "domain.yml")
	content := _fileUtils.ReadFileBuf(domainFilePath)

	yaml.Unmarshal(content, &nluDomain)

	return
}

func (s *NluCompileRasaService) convertIntent(projectId uint, projectDir string, nluDomain *serverDomain.NluDomain) (files []string) {
	_fileUtils.RmDir(filepath.Join(projectDir, "intent"))

	existSlotCodeMap := map[string]bool{}

	tasks := s.NluTaskRepo.ListByProjectId(projectId)
	for _, task := range tasks {
		intents := s.NluIntentRepo.ListByTaskIdNoDisabled(task.ID)

		nluTask := serverDomain.NluTask{Version: serverConst.NluVersion}
		for _, intent := range intents {
			nluDomain.Intents = append(nluDomain.Intents, intent.Name)

			nluIntent := serverDomain.NluIntent{Intent: intent.Name}

			sents := s.NluSentRepo.ListByIntentId(intent.ID, "ordr")
			for _, sent := range sents {
				slotNameMap := s.getSlotNameMap(sent.ID)

				s.populateSlots(sent.ID, slotNameMap, &existSlotCodeMap, nluDomain)

				intentExamples := s.genIntentSent(sent, slotNameMap)
				nluIntent.Examples += "- " + intentExamples + "\n"
			}

			nluTask.Intents = append(nluTask.Intents, nluIntent)
		}

		yamlContent := s.NluCompileService.ChangeArrToFlow(nluTask)

		intentFilePath := filepath.Join(projectDir, consts.Rasa.ToString(), "data", "intent", task.Name+".yml")
		_fileUtils.WriteFile(intentFilePath, yamlContent)
	}

	return
}

func (s *NluCompileRasaService) convertSynonym(projectId uint, projectDir string, nluDomain *serverDomain.NluDomain) {
	_fileUtils.RmDir(filepath.Join(projectDir, "synonym"))

	synonyms := s.NluSynonymRepo.ListByProjectId(projectId)
	for _, synonym := range synonyms {
		//nluDomain.Entities = append(nluDomain.Entities, synonym.Name)

		nluSynonym := serverDomain.NluSynonym{Version: serverConst.NluVersion}
		synonymDef := serverDomain.NluSynonymDef{Synonym: fmt.Sprintf("%s_%s",
			synonym.Name, serverConst.SlotTypeAbbrMap["synonym"])}

		synonymItems := s.NluSynonymItemRepo.ListBySynonymId(synonym.ID)
		for _, item := range synonymItems {
			synonymDef.Examples += "- " + item.Name + "\n"
		}
		nluSynonym.SynonymDef = append(nluSynonym.SynonymDef, synonymDef)

		yamlContent := s.NluCompileService.ChangeArrToFlow(nluSynonym)
		filePath := filepath.Join(projectDir, consts.Rasa.ToString(), "data", "synonym", synonym.Name+".yml")
		_fileUtils.WriteFile(filePath, yamlContent)
	}

	return
}
func (s *NluCompileRasaService) convertLookup(projectId uint, projectDir string, nluDomain *serverDomain.NluDomain) {
	_fileUtils.RmDir(filepath.Join(projectDir, "lookup"))

	lookups := s.NluLookupRepo.ListByProjectId(projectId)
	for _, lookup := range lookups {
		//nluDomain.Entities = append(nluDomain.Entities, lookup.Name)

		nluLookup := serverDomain.NluLookup{Version: serverConst.NluVersion}
		lookupItem := serverDomain.NluLookupItem{Lookup: fmt.Sprintf("%s_%s",
			lookup.Name, serverConst.SlotTypeAbbrMap["lookup"])}

		lookupItems := s.NluLookupItemRepo.ListByLookupId(lookup.ID)
		for _, item := range lookupItems {
			lookupItem.Examples += "- " + item.Name + "\n"
		}
		nluLookup.Items = append(nluLookup.Items, lookupItem)

		yamlContent := s.NluCompileService.ChangeArrToFlow(nluLookup)
		filePath := filepath.Join(projectDir, consts.Rasa.ToString(), "data", "lookup", lookup.Name+".yml")
		_fileUtils.WriteFile(filePath, yamlContent)
	}

	return
}
func (s *NluCompileRasaService) convertRegex(projectId uint, projectDir string, nluDomain *serverDomain.NluDomain) {
	_fileUtils.RmDir(filepath.Join(projectDir, "regex"))

	regexes := s.NluRegexRepo.ListByProjectId(projectId)
	for _, regex := range regexes {
		//nluDomain.Entities = append(nluDomain.Entities, regex.Name)

		nluRegex := serverDomain.NluRegex{Version: serverConst.NluVersion}
		regexItem := serverDomain.NluRegexItem{Regex: fmt.Sprintf("%s_%s",
			regex.Name, serverConst.SlotTypeAbbrMap["regex"])}

		regexItems := s.NluRegexItemRepo.ListByRegexId(regex.ID)
		for _, item := range regexItems {
			regexItem.Examples += "- " + item.Name + "\n"
		}
		nluRegex.Items = append(nluRegex.Items, regexItem)

		yamlContent := s.NluCompileService.ChangeArrToFlow(nluRegex)
		filePath := filepath.Join(projectDir, consts.Rasa.ToString(), "data", "regex", regex.Name+".yml")
		_fileUtils.WriteFile(filePath, yamlContent)
	}

	return
}

func (s *NluCompileRasaService) getSlotNameMap(sentId uint) (ret map[string]map[string]string) {
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

func (s *NluCompileRasaService) getSlotTypeAndVal(tp serverConst.NluSlotType, val string) (ret map[string]string) {
	ret = map[string]string{}

	id, _ := strconv.Atoi(val)

	if tp == serverConst.Synonym {
		entity := s.NluSynonymRepo.Get(uint(id))
		ret["code"] = entity.Name
		ret["name"] = entity.Name

	} else if tp == serverConst.Lookup {
		entity := s.NluLookupRepo.Get(uint(id))
		ret["code"] = entity.Name
		ret["name"] = entity.Name

	} else if tp == serverConst.Regex {
		entity := s.NluRegexRepo.Get(uint(id))
		ret["code"] = entity.Name
		ret["name"] = entity.Name
	} else if tp == serverConst.Slot {
		ret["code"] = val
		ret["name"] = val
	}

	return
}

func (s *NluCompileRasaService) populateSlots(sentId uint, slotMap map[string]map[string]string, codeMap *map[string]bool, nluDomain *serverDomain.NluDomain) {
	slots := s.NluSlotRepo.ListBySentId(sentId)
	for _, slot := range slots {
		slotCode := slotMap[fmt.Sprintf("%s-%s", slot.Type, slot.Value)]["code"]
		if slotCode == "" || (*codeMap)[slotCode] {
			continue
		}

		//if slot.Type == "_slot_" && !_stringUtils.StrInArr(slotCode, nluDomain.Entities) {
		nluDomain.Entities = append(nluDomain.Entities, slotCode)
		//}

		slotItem := serverDomain.SlotItem{Type: "text", InfluenceConversation: false}
		mapItem := yaml.MapItem{Key: slotCode, Value: slotItem}
		nluDomain.Slots = append(nluDomain.Slots, mapItem)

		(*codeMap)[slotCode] = true
	}
}

func (s *NluCompileRasaService) genIntentSent(sent model.NluSent, slotMap map[string]map[string]string) (ret string) {
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
