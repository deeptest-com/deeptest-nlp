package service

import (
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/repo"
)

type NluSentService struct {
	NluSentRepo *repo.NluSentRepo `inject:""`
}

func NewNluSentService() *NluSentService {
	return &NluSentService{}
}

func (s *NluSentService) List(keywords, status string, pageNo int, pageSize int) (pos []model.NluSent, total int64) {
	pos, total = s.NluSentRepo.Query(keywords, status, pageNo, pageSize)
	return
}
func (s *NluSentService) ListByIntent(intentId uint) (pos []model.NluSent) {
	pos = s.NluSentRepo.ListByIntent(intentId)
	return
}

func (s *NluSentService) Get(id uint) (po model.NluSent) {
	po = s.NluSentRepo.Get(id)
	return
}

func (s *NluSentService) GetWithSlots(id uint) (po model.NluSent) {
	po = s.NluSentRepo.GetWithSlots(id)

	return
}

func (s *NluSentService) Save(po *model.NluSent) (err error) {
	err = s.NluSentRepo.Save(po)

	return
}

func (s *NluSentService) Update(po *model.NluSent) (err error) {
	err = s.NluSentRepo.Update(po)

	return
}

func (s *NluSentService) SetDefault(id uint) (err error) {
	err = s.NluSentRepo.SetDefault(id)

	return
}

func (s *NluSentService) Disable(id uint) (err error) {
	err = s.NluSentRepo.Disable(id)

	return
}

func (s *NluSentService) Delete(id uint) (err error) {
	err = s.NluSentRepo.Delete(id)

	return
}

func (s *NluSentService) BatchDelete(ids []int) (err error) {
	err = s.NluSentRepo.BatchDelete(ids)

	return
}
