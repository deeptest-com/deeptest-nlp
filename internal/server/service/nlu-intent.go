package service

import (
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/repo"
)

type NluIntentService struct {
	NluIntentRepo *repo.NluIntentRepo `inject:""`
}

func NewNluIntentService() *NluIntentService {
	return &NluIntentService{}
}

func (s *NluIntentService) List(keywords string, pageNo int, pageSize int) (pos []model.NluIntent, total int64) {
	pos, total = s.NluIntentRepo.Query(keywords, pageNo, pageSize)
	return
}

func (s *NluIntentService) Save(pos *model.NluIntent) (err error) {
	err = s.NluIntentRepo.Save(pos)

	return
}
