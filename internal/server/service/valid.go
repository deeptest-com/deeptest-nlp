package service

import (
	_const "github.com/utlai/utl/internal/pkg/const"
	_fileUtils "github.com/utlai/utl/internal/pkg/libs/file"
	"github.com/utlai/utl/internal/server/domain"
	"path/filepath"
)

type ValidService struct {
}

func NewValidService() *ValidService {
	return &ValidService{}
}

func (s *ValidService) Valid(model domain.ValidRequest) (result domain.ValidResp) {
	if model.Method == _const.ValidProjectPath {
		result = s.ValidProjectPath(model.Value)
	}

	return
}

func (s *ValidService) ValidProjectPath(value string) (result domain.ValidResp) {
	if !_fileUtils.FileExist(value) {
		result.Pass = false
		return
	}

	configFile := filepath.Join(value, "config.yml")
	if !_fileUtils.FileExist(configFile) {
		result.Pass = false
		return
	}

	result.Pass = true
	return
}
