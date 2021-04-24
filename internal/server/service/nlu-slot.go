package service

import (
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/repo"
)

type NluSlotService struct {
	NluSlotRepo *repo.NluSlotRepo `inject:""`
}

func NewNluSlotService() *NluSlotService {
	return &NluSlotService{}
}

func (s *NluSlotService) List(keywords, status string, pageNo int, pageSize int) (pos []model.NluSlot, total int64) {
	pos, total = s.NluSlotRepo.Query(keywords, status, pageNo, pageSize)
	return
}

func (s *NluSlotService) Get(id uint) (po model.NluSlot) {
	po = s.NluSlotRepo.Get(id)
	return
}

func (s *NluSlotService) Save(po *model.NluSlot) (err error) {
	err = s.NluSlotRepo.Save(po)

	return
}

func (s *NluSlotService) Update(po *model.NluSlot) (err error) {
	err = s.NluSlotRepo.Update(po)

	return
}

func (s *NluSlotService) SetDefault(id uint) (err error) {
	err = s.NluSlotRepo.SetDefault(id)

	return
}

func (s *NluSlotService) Disable(id uint) (err error) {
	err = s.NluSlotRepo.Disable(id)

	return
}

func (s *NluSlotService) Delete(id uint) (err error) {
	err = s.NluSlotRepo.Delete(id)

	return
}

func (s *NluSlotService) BatchDelete(ids []int) (err error) {
	err = s.NluSlotRepo.BatchDelete(ids)

	return
}
