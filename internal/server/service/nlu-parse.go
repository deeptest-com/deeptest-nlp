package service

import (
	"fmt"
	_httpUtils "github.com/utlai/utl/internal/pkg/libs/http"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/utlai/utl/internal/server/domain"
	"github.com/utlai/utl/internal/server/repo"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
	"regexp"
	"strings"
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
}

func NewNluParseService() *NluParseService {
	return &NluParseService{}
}

func (s *NluParseService) Parse(projectId int, req domain.NluReq) (ret domain.NluResp) {
	ret.Code = -1
	msg := map[string]string{}

	project := s.ProjectRepo.Get(uint(projectId))
	serviceStatus := project.ServiceStatus
	port := project.ServicePort

	if serviceStatus != serverConst.StartService {
		msg["msg"] = "service not started"
		ret.SetResult(msg)
		return
	}

	regexMap := s.GetRuleRegex(uint(projectId))
	req.Text = s.ReplaceWithRegex(req.Text, regexMap)

	url := fmt.Sprintf("http://127.0.0.1:%d/%s", port, "model/parse")
	resp, success := _httpUtils.PostRasa(url, req)
	if !success {
		msg["msg"] = fmt.Sprintf("rasa request failed, response %v", resp)
		ret.SetResult(msg)
		return
	}

	ret.Result = resp.Payload
	ret.Code = 1

	return
}

func (s *NluParseService) ReplaceWithRegex(sent string, regexMap []map[string]map[int][]string) (ret string) {
	for _, item := range regexMap {
		for regex, slotMap := range item {
			_logUtils.Infof("%s, %v", regex, slotMap)
		}
	}

	return
}

func (s *NluParseService) GetRuleRegex(projectId uint) (ret []map[string]map[int][]string) {
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

func (s *NluParseService) GenRegexStr(ruleText string) (regex string, slotMap map[int][]string) {
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
			dict := s.NluSynonymRepo.GetByName(text)
			dictItems := s.NluSynonymItemRepo.ListBySynonymId(dict.ID)

			for _, i := range dictItems {
				list = append(list, i.Content)
			}
		} else if tag == "l" {
			dict := s.NluLookupRepo.GetByName(text)
			dictItems := s.NluLookupItemRepo.ListByLookupId(dict.ID)

			for _, i := range dictItems {
				list = append(list, i.Content)
			}
		} else if tag == "r" {
			dict := s.NluRegexRepo.GetByName(text)
			dictItems := s.NluRegexItemRepo.ListByRegexId(dict.ID)

			for _, i := range dictItems {
				list = append(list, i.Content)
			}
		} else if tag == "_" {
			list = append(list, ".*")
			slotMap[index] = []string{text}
		}

		regex += "(" + strings.Join(list, "|") + ")"
	}

	return
}
