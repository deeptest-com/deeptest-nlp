package service

import (
	"encoding/json"
	"fmt"
	_httpUtils "github.com/utlai/utl/internal/pkg/libs/http"
	"github.com/utlai/utl/internal/server/domain"
	"github.com/utlai/utl/internal/server/repo"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
)

type NluParseService struct {
	ProjectRepo *repo.ProjectRepo `inject:""`
}

func NewNluParseService() *NluParseService {
	return &NluParseService{}
}

func (s *NluParseService) Parse(projectId int, req domain.NluReq) (ret domain.NluResp) {
	ret.Code = -1
	msg := map[string]string{}

	project := s.ProjectRepo.Get(uint(projectId))
	serviceStatus := project.ServiceStatus
	port := project.ServicePort

	if serviceStatus != serverConst.StartService {
		msg["msg"] = "service not started"
		bytes, _ := json.Marshal(msg)
		ret.Result = string(bytes)
		return
	}

	url := fmt.Sprintf("http://127.0.0.1:%d/%s", port, "model/parse")
	resp, success := _httpUtils.PostRasa(url, req)
	if !success {
		msg["msg"] = fmt.Sprintf("rasa request failed, response %v", resp)
		bytes, _ := json.Marshal(msg)
		ret.Result = string(bytes)
		return
	}

	ret.Result = resp.Payload
	ret.Code = 1

	return
}
