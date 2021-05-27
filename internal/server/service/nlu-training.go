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

	go s.CallTraining(project)

	_logUtils.Infof("--- 1. call training project %s---", project.Path)

	return
}

func (s *NluTrainingService) CallTraining(project model.Project) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*serverConst.TrainingTimeout)
	defer cancel()

	ch := make(chan struct{}, 0)

	go func() {
		_logUtils.Infof("--- 2. %s start training project %s---",
			time.Now().Format("2006-01-02 15:04:05"), project.Path)

		s.ExecTraining(project)

		_logUtils.Infof("--- 3. %s end training project %s---",
			time.Now().Format("2006-01-02 15:04:05"), project.Path)

		ch <- struct{}{}
	}()

	select {
	case <-ch:
		_logUtils.Infof("--- 4. %s complete training project %s---",
			time.Now().Format("2006-01-02 15:04:05"), project.Path)
	case <-ctx.Done():
		_logUtils.Infof("--- 0. %s timeout training project %s---",
			time.Now().Format("2006-01-02 15:04:05"), project.Path)

		s.CancelTraining()
	}
}

func (s *NluTrainingService) ExecTraining(project model.Project) {
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

func (s *NluTrainingService) CancelTraining() (result string, err error) {
	cmdStr := "ps -ef | grep 'rasa train' | grep -v grep | awk '{print $2}' | xargs kill -9"
	result, err = _shellUtils.ExeShell(cmdStr)

	return
}
