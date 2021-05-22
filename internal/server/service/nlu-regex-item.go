package service

import (
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/repo"
)

type NluRegexItemService struct {
	NluRegexItemRepo *repo.NluRegexItemRepo `inject:""`
}

func NewNluRegexItemService() *NluRegexItemService {
	return &NluRegexItemService{}
}

func (s *NluRegexItemService) List(regexId int, keywords, status string, pageNo int, pageSize int) (pos []model.NluRegexItem, total int64) {
	pos, total = s.NluRegexItemRepo.Query(regexId, keywords, status, pageNo, pageSize)
	return
}

func (s *NluRegexItemService) Get(id uint) (po model.NluRegexItem) {
	po = s.NluRegexItemRepo.Get(id)
	return
}

func (s *NluRegexItemService) Save(po *model.NluRegexItem) (err error) {
	err = s.NluRegexItemRepo.Save(po)

	return
}

func (s *NluRegexItemService) Update(po *model.NluRegexItem) (err error) {
	err = s.NluRegexItemRepo.Update(po)

	return
}

func (s *NluRegexItemService) SetDefault(id uint) (err error) {
	err = s.NluRegexItemRepo.SetDefault(id)

	return
}

func (s *NluRegexItemService) Disable(id uint) (err error) {
	err = s.NluRegexItemRepo.Disable(id)

	return
}

func (s *NluRegexItemService) Delete(id uint) (err error) {
	err = s.NluRegexItemRepo.Delete(id)

	return
}

func (s *NluRegexItemService) BatchDelete(ids []int) (err error) {
	err = s.NluRegexItemRepo.BatchDelete(ids)

	return
}
