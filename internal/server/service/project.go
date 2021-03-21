package service

import (
	_const "github.com/utlai/utl/internal/pkg/const"
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/repo"
)

type ProjectService struct {
	ProjectRepo *repo.ProjectRepo `inject:""`
}

func NewProjectService() *ProjectService {
	return &ProjectService{}
}

func (s *ProjectService) List(keywords string, pageNo int, pageSize int) (projects []model.Project, total int64) {
	projects, total = s.ProjectRepo.Query(keywords, pageNo, pageSize)
	return
}

func (s *ProjectService) Save(project *model.Project) (err error) {
	err = s.ProjectRepo.Save(project)

	return
}

func (s *ProjectService) SetProgress(id uint, progress _const.BuildProgress) {
	s.ProjectRepo.SetProgress(id, progress)
}
