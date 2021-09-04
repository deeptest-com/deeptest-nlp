package serverService

import (
	consts "github.com/utlai/utl/internal/comm/const"
	_const "github.com/utlai/utl/internal/pkg/const"
	_fileUtils "github.com/utlai/utl/internal/pkg/libs/file"
	serverConf "github.com/utlai/utl/internal/server/conf"
	"github.com/utlai/utl/internal/server/domain"
	"github.com/utlai/utl/internal/server/repo"
	"path/filepath"
	"strings"
)

type ValidService struct {
	NluSynonymRepo *repo.NluSynonymRepo `inject:""`
	NluLookupRepo  *repo.NluLookupRepo  `inject:""`
	NluRegexRepo   *repo.NluRegexRepo   `inject:""`
}

func NewValidService() *ValidService {
	return &ValidService{}
}

func (s *ValidService) Valid(model serverDomain.ValidRequest) (result serverDomain.ValidResp) {
	if model.Method == _const.ValidProjectPath {
		result = s.ValidProjectPath(model.Value)
	} else if model.Method == _const.ValidDictName {
		result = s.ValidDictName(strings.TrimSpace(model.Value), uint(model.Id), model.Type)
	}

	return
}

func (s *ValidService) ValidProjectPath(value string) (result serverDomain.ValidResp) {
	if !_fileUtils.FileExist(value) {
		result.Pass = false
		return
	}

	configFile := filepath.Join(value, "config.yml")
	if serverConf.Inst.Analyzer == consts.Rasa && !_fileUtils.FileExist(configFile) {
		result.Pass = false
		return
	}

	result.Pass = true
	return
}

func (s *ValidService) ValidDictName(value string, id uint, tp string) (result serverDomain.ValidResp) {
	if tp == "synonym" {
		po := s.NluSynonymRepo.GetByName(value)
		if po.ID == id {
			result.Pass = true
		}
	} else if tp == "lookup" {
		po := s.NluLookupRepo.GetByName(value)
		if po.ID == id {
			result.Pass = true
		}
	} else if tp == "regex" {
		po := s.NluRegexRepo.GetByName(value)
		if po.ID == id {
			result.Pass = true
		}
	}

	return
}
