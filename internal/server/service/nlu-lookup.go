package serverService

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

func (s *NluLookupService) List(keywords, status string, pageNo int, pageSize int) (pos []model.NluLookup, total int64) {
	pos, total = s.NluLookupRepo.Query(keywords, status, pageNo, pageSize)
	return
}

func (s *NluLookupService) Get(id uint) (po model.NluLookup) {
	po = s.NluLookupRepo.Get(id)
	return
}

func (s *NluLookupService) Save(po *model.NluLookup) (err error) {
	err = s.NluLookupRepo.Save(po)

	return
}

func (s *NluLookupService) Update(po *model.NluLookup) (err error) {
	err = s.NluLookupRepo.Update(po)

	return
}

func (s *NluLookupService) SetDefault(id uint) (err error) {
	err = s.NluLookupRepo.SetDefault(id)

	return
}

func (s *NluLookupService) Disable(id uint) (err error) {
	err = s.NluLookupRepo.Disable(id)

	return
}

func (s *NluLookupService) Delete(id uint) (err error) {
	err = s.NluLookupRepo.Delete(id)

	return
}

func (s *NluLookupService) BatchDelete(ids []int) (err error) {
	err = s.NluLookupRepo.BatchDelete(ids)

	return
}

func (s *NluLookupService) Resort(srcId, targetId int) (err error) {
	err = s.NluLookupRepo.Resort(srcId, targetId)

	return
}
