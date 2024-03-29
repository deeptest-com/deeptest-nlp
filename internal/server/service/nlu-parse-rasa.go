package serverService

import (
	"fmt"
	"github.com/utlai/utl/internal/comm/domain"
	_httpUtils "github.com/utlai/utl/internal/pkg/libs/http"
	_i118Utils "github.com/utlai/utl/internal/pkg/libs/i118"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/utlai/utl/internal/server/domain"
	"github.com/utlai/utl/internal/server/repo"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
	"regexp"
	"strings"
)

type NluParseRasaService struct {
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
}

func NewNluParseRasaService() *NluParseRasaService {
	return &NluParseRasaService{}
}

func (s *NluParseRasaService) Parse(projectId uint, req serverDomain.NluReq) (ret serverDomain.NluResp) {
	ret.Code = -1
	msg := map[string]string{}

	project := s.ProjectRepo.Get(projectId)
	serviceStatus := project.ServiceStatus
	port := project.ServicePort

	if serviceStatus != serverConst.StartService {
		msg["msg"] = _i118Utils.Sprintf("service.not.start")
		msg["content"] = ""
		ret.SetMsg(msg)
		return
	}

	regexMap := s.GetRuleRegex(projectId)
	originMap := map[string]string{}
	req.TextOrigin = req.Text
	req.Text, originMap = s.ReplaceWithRegex(req.Text, regexMap)

	url := fmt.Sprintf("http://127.0.0.1:%d/%s", port, "model/parse")
	resp, success := _httpUtils.PostRasa(url, req)
	if !success {
		msg["msg"] = _i118Utils.Sprintf("rasa.request.fail")
		msg["content"] = fmt.Sprintf("%#v", resp)
		ret.SetMsg(msg)
		return
	}

	rasaResp := resp.Payload.(domain.RasaResp)
	rasaResp.TextOrigin = req.TextOrigin
	for index, entity := range rasaResp.Entities {
		str, ok := originMap[entity.Entity]
		if ok {
			rasaResp.Entities[index].ValueOrigin = str
		}
	}

	ret.SetResult(rasaResp)
	ret.Code = 1

	return
}

func (s *NluParseRasaService) ReplaceWithRegex(sent string, regexMap []map[string]map[int][]string) (
	ret string, originMap map[string]string) {
	originMap = map[string]string{}

	for _, item := range regexMap {
		for regex, slotMap := range item {
			_logUtils.Infof("%s, %v", regex, slotMap)

			re := regexp.MustCompile(regex)
			items := re.FindAllStringSubmatch(sent, -1)

			if len(items) > 0 {
				for index, item := range items[0] {
					if index == 0 {
						continue
					}

					strs, ok := slotMap[index-1]
					if ok {
						ret += strs[0]
						originMap[strs[1]] = item
					} else {
						ret += item
					}
				}
			}
		}
	}

	return
}

func (s *NluParseRasaService) GetRuleRegex(projectId uint) (ret []map[string]map[int][]string) {
	boolObj, ok1 := ProjectChanged.Load(projectId)
	changed := false
	if !ok1 || boolObj.(bool) {
		changed = true
	}

	regexBoj, ok2 := ProjectRules.Load(projectId)
	if ok2 && !changed {
		ret = regexBoj.([]map[string]map[int][]string)
		return
	}

	rules := s.NluRuleRepo.ProjectId(projectId)
	for _, rule := range rules {
		if rule.Disabled {
			continue
		}

		regex, slotMap := s.GenRegexStr(rule.Text)
		mp := map[string]map[int][]string{}
		mp[regex] = slotMap
		ret = append(ret, mp)
	}
	ProjectRules.Store(projectId, ret)

	return
}

func (s *NluParseRasaService) GenRegexStr(ruleText string) (regex string, slotMap map[int][]string) {
	// {打印:S}{日志级别:L}(内容:_)
	slotMap = map[int][]string{}

	re := regexp.MustCompile(`(?siU)(\{|\()(.*)(\}|\))`)
	items := re.FindAllStringSubmatch(ruleText, -1)
	for index, item := range items {
		content := item[2]
		arr := strings.Split(content, ":")
		text := arr[0]
		tag := strings.ToLower(arr[1])

		list := make([]string, 0)
		if tag == "s" {
			dict := s.NluSynonymRepo.GetByCode(text)
			dictItems := s.NluSynonymItemRepo.ListBySynonymId(dict.ID)

			for _, i := range dictItems {
				list = append(list, i.Name)
			}
		} else if tag == "l" {
			dict := s.NluLookupRepo.GetByCode(text)
			dictItems := s.NluLookupItemRepo.ListByLookupId(dict.ID)

			for _, i := range dictItems {
				list = append(list, i.Name)
			}
		} else if tag == "r" {
			dict := s.NluRegexRepo.GetByCode(text)
			dictItems := s.NluRegexItemRepo.ListByRegexId(dict.ID)

			for _, i := range dictItems {
				list = append(list, i.Name)
			}
		} else {
			list = append(list, ".*")
			slotMap[index] = []string{text, tag}
		}

		regex += "(" + strings.Join(list, "|") + ")"
	}

	return
}
