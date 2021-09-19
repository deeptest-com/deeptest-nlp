package serverService

import (
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/repo"
)

type NluLookupItemService struct {
	NluLookupItemRepo *repo.NluLookupItemRepo `inject:""`
}

func NewNluLookupItemService() *NluLookupItemService {
	return &NluLookupItemService{}
}

func (s *NluLookupItemService) List(lookupId int, keywords, status string, pageNo int, pageSize int) (pos []model.NluLookupItem, total int64) {
	pos, total = s.NluLookupItemRepo.Query(lookupId, keywords, status, pageNo, pageSize)
	return
}

func (s *NluLookupItemService) Get(id uint) (po model.NluLookupItem) {
	po = s.NluLookupItemRepo.Get(id)
	return
}

func (s *NluLookupItemService) Save(po *model.NluLookupItem) (err error) {
	err = s.NluLookupItemRepo.Save(po)

	return
}

func (s *NluLookupItemService) Update(po *model.NluLookupItem) (err error) {
	err = s.NluLookupItemRepo.Update(po)

	return
}

func (s *NluLookupItemService) SetDefault(id uint) (err error) {
	err = s.NluLookupItemRepo.SetDefault(id)

	return
}

func (s *NluLookupItemService) Disable(id uint) (err error) {
	err = s.NluLookupItemRepo.Disable(id)

	return
}

func (s *NluLookupItemService) Delete(id uint) (err error) {
	err = s.NluLookupItemRepo.Delete(id)

	return
}

func (s *NluLookupItemService) BatchDelete(ids []int) (err error) {
	err = s.NluLookupItemRepo.BatchDelete(ids)

	return
}

func (s *NluLookupItemService) Resort(srcId, targetId, parentId int) (err error) {
	err = s.NluLookupItemRepo.Resort(srcId, targetId, parentId)

	return
}
