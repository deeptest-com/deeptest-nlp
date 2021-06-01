package service

import "github.com/utlai/utl/internal/server/domain"

type NluParseService struct {
}

func NewNluParseService() *NluParseService {
	return &NluParseService{}
}

func (s *NluParseService) Parse(projectId int, req domain.NluReq) (resp domain.NluResp) {
	resp.Code = -1
	resp.Result = "{\"key\": \"value\"}"

	return
}
