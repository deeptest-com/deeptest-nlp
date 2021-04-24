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

func (s *NluIntentService) List(keywords, status string, pageNo int, pageSize int) (pos []model.NluIntent, total int64) {
	pos, total = s.NluIntentRepo.Query(keywords, status, pageNo, pageSize)
	return
}

func (s *NluIntentService) Get(id uint) (po model.NluIntent) {
	po = s.NluIntentRepo.Get(id)
	return
}

func (s *NluIntentService) Save(po *model.NluIntent) (err error) {
	err = s.NluIntentRepo.Save(po)

	return
}

func (s *NluIntentService) Update(po *model.NluIntent) (err error) {
	err = s.NluIntentRepo.Update(po)

	return
}

func (s *NluIntentService) SetDefault(id uint) (err error) {
	err = s.NluIntentRepo.SetDefault(id)

	return
}

func (s *NluIntentService) Disable(id uint) (err error) {
	err = s.NluIntentRepo.Disable(id)

	return
}

func (s *NluIntentService) Delete(id uint) (err error) {
	err = s.NluIntentRepo.Delete(id)

	return
}

func (s *NluIntentService) BatchDelete(ids []int) (err error) {
	err = s.NluIntentRepo.BatchDelete(ids)

	return
}
