package service

import (
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/utlai/utl/internal/server/repo"
)

type NluConvertService struct {
	ProjectRepo   *repo.ProjectRepo   `inject:""`
	NluTaskRepo   *repo.NluTaskRepo   `inject:""`
	NluIntentRepo *repo.NluIntentRepo `inject:""`
	NluSentRepo   *repo.NluSentRepo   `inject:""`
}

func NewNluConvertService() *NluConvertService {
	return &NluConvertService{}
}

func (s *NluConvertService) ConvertProject(id uint) (files []string) {
	tasks := s.NluTaskRepo.ListByProjectId(id)
	for _, task := range tasks {
		intents := s.NluIntentRepo.ListByTaskId(task.ID)

		for _, intent := range intents {
			sents := s.NluSentRepo.ListByIntentId(intent.ID)

			for _, sent := range sents {
				_logUtils.Info(sent.Text)
			}
		}
	}

	return
}
