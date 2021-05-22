package service

import (
	_const "github.com/utlai/utl/internal/pkg/const"
	_fileUtils "github.com/utlai/utl/internal/pkg/libs/file"
	"github.com/utlai/utl/internal/server/domain"
)

type ValidService struct {
}

func NewValidService() *ValidService {
	return &ValidService{}
}

func (s *ValidService) Valid(model domain.ValidRequest) (result domain.ValidResp) {
	if model.Method == _const.ValidPath {
		result = s.ValidPath(model.Value)
	}

	return
}

func (s *ValidService) ValidPath(value string) (result domain.ValidResp) {
	if _fileUtils.FileExist(value) {
		result.Pass = true
	}

	return
}
