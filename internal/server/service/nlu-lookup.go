package service

import (
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/repo"
)

type NluLookupService struct {
	NluLookupRepo *repo.NluLookupRepo `inject:""`
}

func NewNluLookupService() *NluLookupService {
	return &NluLookupService{}
}

func (s *NluLookupService) List(keywords string, pageNo int, pageSize int) (pos []model.NluLookup, total int64) {
	pos, total = s.NluLookupRepo.Query(keywords, pageNo, pageSize)
	return
}

func (s *NluLookupService) Save(po *model.NluLookup) (err error) {
	err = s.NluLookupRepo.Save(po)

	return
}
