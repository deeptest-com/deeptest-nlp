package service

import (
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/utlai/utl/internal/server/domain"
	"github.com/utlai/utl/internal/server/repo"
	serverVari "github.com/utlai/utl/internal/server/utils/var"
	"regexp"
	"strconv"
	"strings"
)

type NluParsePatternService struct {
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

	NluPatternService *NluPatternService `inject:""`
}

func NewNluParsePatternService() *NluParsePatternService {
	return &NluParsePatternService{}
}

func (s *NluParsePatternService) Parse(projectId uint, req domain.NluReq) (ret domain.NluResp) {
	ret.Code = -1

	text := req.TextOrigin
	if serverVari.PatternData[projectId] == nil {
		s.NluPatternService.Reload(projectId)
	}
	patternMap := serverVari.PatternData[projectId]

	for key, patterns := range patternMap {
		for _, p := range patterns {
			rgx, err := regexp.Compile(p)

			if err != nil {
				_logUtils.Error(err.Error())
				continue
			}

			arr := rgx.FindStringSubmatch(text)
			_logUtils.Infof("%s, %#v", key, arr)

			if len(arr) > 0 {
				idStr := strings.Split(key, "-")[0]
				id, _ := strconv.Atoi(idStr)

				intent := s.NluIntentRepo.Get(uint(id))
				_logUtils.Infof("intent %s", intent.Name)
			}
		}
	}

	//ret.SetResult(rasaResp)
	ret.Code = 1

	return
}
