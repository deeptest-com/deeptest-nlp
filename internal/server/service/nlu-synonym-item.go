package serverService

import (
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/repo"
)

type NluSynonymItemService struct {
	NluSynonymItemRepo *repo.NluSynonymItemRepo `inject:""`
}

func NewNluSynonymItemService() *NluSynonymItemService {
	return &NluSynonymItemService{}
}

func (s *NluSynonymItemService) List(synonymId int, keywords, status string, pageNo int, pageSize int) (pos []model.NluSynonymItem, total int64) {
	pos, total = s.NluSynonymItemRepo.Query(synonymId, keywords, status, pageNo, pageSize)
	return
}

func (s *NluSynonymItemService) Get(id uint) (po model.NluSynonymItem) {
	po = s.NluSynonymItemRepo.Get(id)
	return
}

func (s *NluSynonymItemService) Save(po *model.NluSynonymItem) (err error) {
	err = s.NluSynonymItemRepo.Save(po)

	return
}

func (s *NluSynonymItemService) Update(po *model.NluSynonymItem) (err error) {
	err = s.NluSynonymItemRepo.Update(po)

	return
}

func (s *NluSynonymItemService) SetDefault(id uint) (err error) {
	err = s.NluSynonymItemRepo.SetDefault(id)

	return
}

func (s *NluSynonymItemService) Disable(id uint) (err error) {
	err = s.NluSynonymItemRepo.Disable(id)

	return
}

func (s *NluSynonymItemService) Delete(id uint) (err error) {
	err = s.NluSynonymItemRepo.Delete(id)

	return
}

func (s *NluSynonymItemService) BatchDelete(ids []int) (err error) {
	err = s.NluSynonymItemRepo.BatchDelete(ids)

	return
}

func (s *NluSynonymItemService) Resort(srcId, targetId, parentId int) (err error) {
	err = s.NluSynonymItemRepo.Resort(srcId, targetId, parentId)

	return
}
