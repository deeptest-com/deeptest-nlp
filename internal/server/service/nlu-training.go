package service

import (
	"context"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	_shellUtils "github.com/utlai/utl/internal/pkg/libs/shell"
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/repo"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
	"time"
)

type NluTrainingService struct {
	ProjectRepo       *repo.ProjectRepo  `inject:""`
	NluServiceService *NluServiceService `inject:""`
}

func NewNluTrainingService() *NluTrainingService {
	return &NluTrainingService{}
}

func (s *NluTrainingService) TrainingProject(id uint) (files []string) {
	project := s.ProjectRepo.GetDetail(id)

	s.AsyncCall(project)

	return
}

func (s *NluTrainingService) AsyncCall(project model.Project) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*serverConst.TrainingTimeout)
	defer cancel()
	go func(ctx context.Context) {

		s.Training(project)

	}(ctx)

	select {
	case <-ctx.Done():
		_logUtils.Infof("---async training completed---")
		return
	case <-time.After(time.Millisecond * serverConst.TrainingTimeout):
		_logUtils.Infof("---async training timeout---")
		return
	}
}

func (s *NluTrainingService) Training(project model.Project) {
	s.NluServiceService.Stop(project)

	cmdStr := "rm -rf models"
	ret, err := _shellUtils.ExeShellInDir(cmdStr, project.Path)

	cmdStr = "rasa train"
	ret, err = _shellUtils.ExeShellInDir(cmdStr, project.Path)
	if err != nil {
		_logUtils.Errorf("training failed, error %s", err)
	} else {
		_logUtils.Infof("training successfully, %s", ret)
	}

	s.NluServiceService.Start(project)
}
