package service

import (
	"fmt"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	_shellUtils "github.com/utlai/utl/internal/pkg/libs/shell"
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/repo"
)

type NluServiceService struct {
	ProjectRepo *repo.ProjectRepo `inject:""`
}

func NewNluServiceService() *NluServiceService {
	return &NluServiceService{}
}

func (s *NluServiceService) ReStart(id uint) (result string, err error) {
	project := s.ProjectRepo.GetDetail(id)

	s.Stop(project)
	s.Start(project)

	if err != nil {
		_logUtils.Errorf("%s, %s", result, err.Error())
	}

	return
}

func (s *NluServiceService) Stop(project model.Project) (result string, err error) {
	cmdStr := fmt.Sprintf("ps -ef | grep 'm models_%d' | grep -v grep | awk '{print $2}' | xargs kill -9", project.ID)
	_logUtils.Infof("--- stop service project %s---", cmdStr)

	result, err = _shellUtils.ExeShell(cmdStr)

	return
}

func (s *NluServiceService) Start(project model.Project) (result string, err error) {
	cmdStr := fmt.Sprintf("nohup rasa run --enable-api -m models_%d --log-file out.log 2>&1 &", project.ID)
	_logUtils.Infof("--- start service project %s---", cmdStr)

	result, err = _shellUtils.ExeShellInDir(cmdStr, project.Path)

	return
}
