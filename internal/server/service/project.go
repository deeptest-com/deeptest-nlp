package service

import (
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/repo"
)

type ProjectService struct {
	ProjectRepo *repo.ProjectRepo `inject:""`
}

func NewProjectService() *ProjectService {
	return &ProjectService{}
}

func (s *ProjectService) List(keywords, status string, pageNo int, pageSize int) (pos []model.Project, total int64) {
	pos, total = s.ProjectRepo.Query(keywords, status, pageNo, pageSize)
	return
}

func (s *ProjectService) Save(pos *model.Project) (err error) {
	err = s.ProjectRepo.Save(pos)

	return
}

func (s *ProjectService) SetDefault(id uint) (err error) {
	err = s.ProjectRepo.SetDefault(id)

	return
}
