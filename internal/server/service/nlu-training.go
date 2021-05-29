package service

import (
	"context"
	"fmt"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	_shellUtils "github.com/utlai/utl/internal/pkg/libs/shell"
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/repo"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
	"strings"
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
		_logUtils.Infof("--- 2. start training project %s---", project.Path)

		s.ExecTraining(project)

		_logUtils.Infof("--- 3. end training project %s---", project.Path)

		ch <- struct{}{}
	}()

	select {
	case <-ch:
		_logUtils.Infof("--- 4. finish training project %s---", project.Path)
	case <-ctx.Done():
		_logUtils.Infof("--- 0. timeout training project %s---", project.Path)

		s.CancelTraining()
	}
}

func (s *NluTrainingService) ExecTraining(project model.Project) {
	// kill old training process
	pName := fmt.Sprintf("out models_%d", project.ID)
	_shellUtils.KillProcess(pName)

	// stop service
	s.NluServiceService.Stop(project)

	// rm models
	cmdStr := "rm -rf models_*"
	ret, err := _shellUtils.ExeShellInDir(cmdStr, project.Path)

	// start training
	cmdStr = fmt.Sprintf("rasa train --out models_%d", project.ID)
	ret, err = _shellUtils.ExeShellInDir(cmdStr, project.Path)

	if err != nil { // e.x. killed by new one
		_logUtils.Errorf("--- training failed, error %s", err)
		return
	}

	_logUtils.Infof("--- training successfully: \n%s\n%s\n%s",
		strings.Repeat("*", 100), ret, strings.Repeat("*", 100))

	// start service
	s.NluServiceService.Start(project)
}

func (s *NluTrainingService) CancelTraining() (result string, err error) {
	cmdStr := "ps -ef | grep 'rasa train' | grep -v grep | awk '{print $2}' | xargs kill -9"
	result, err = _shellUtils.ExeShell(cmdStr)

	return
}
