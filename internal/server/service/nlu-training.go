package service

import (
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/utlai/utl/internal/server/repo"
)

type NluTrainingService struct {
	ProjectRepo *repo.ProjectRepo `inject:""`
}

func NewNluTrainingService() *NluTrainingService {
	return &NluTrainingService{}
}

func (s *NluTrainingService) TrainingProject(id uint) (files []string) {
	project := s.ProjectRepo.Get(id)
	_logUtils.Infof(project.Name)

	return
}
