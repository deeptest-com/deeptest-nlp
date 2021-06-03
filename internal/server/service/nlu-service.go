package service

import (
	"fmt"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	_shellUtils "github.com/utlai/utl/internal/pkg/libs/shell"
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/repo"
	"strconv"
	"strings"
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

	result, err, _ = _shellUtils.ExeShell(cmdStr, "")
	s.ProjectRepo.StopService(project.ID)

	return
}

func (s *NluServiceService) Start(project model.Project) (result string, err error) {
	port := getValidPort()
	cmdStr := fmt.Sprintf("nohup rasa run -p %d -m models_%d --enable-api --log-file out.log > nohup.log 2>&1 &",
		port, project.ID)
	_logUtils.Infof("--- start service project %s---", cmdStr)

	result, err, _ = _shellUtils.ExeShell(cmdStr, project.Path)

	s.ProjectRepo.StartService(project.ID, port)

	return
}

func getValidPort() (port int) {
	cmd := "ps -ef | grep 'rasa run' | grep -v 'grep' | awk '{print $12}'"
	output, err, _ := _shellUtils.ExeShell(cmd, "")

	port = 55005
	if err != nil || output == "" {
		return
	}

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strconv.Itoa(port) == line {
			port += 1
		}
	}

	return
}
