package service

import (
	"github.com/utlai/utl/internal/server/repo"
)

type NluDictService struct {
	NluSynonymRepo *repo.NluSynonymRepo `inject:""`
	NluLookupRepo  *repo.NluLookupRepo  `inject:""`
}

func NewNluDictService() *NluDictService {
	return &NluDictService{}
}

func (s *NluDictService) List(keywords, status string, pageNo int, pageSize int) (pos []interface{}, total int64) {
	//pos, total = s.NluDictRepo.Query(keywords, status, pageNo, pageSize)
	return
}
