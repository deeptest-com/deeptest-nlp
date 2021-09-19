package serverService

import (
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/repo"
)

type NluPlaceholderService struct {
	NluPlaceholderRepo *repo.NluPlaceholderRepo `inject:""`
}

func NewNluPlaceholderService() *NluPlaceholderService {
	return &NluPlaceholderService{}
}

func (s *NluPlaceholderService) List(keywords, status string, pageNo, pageSize, projectId int) (pos []model.NluPlaceholder, total int64) {
	pos, total = s.NluPlaceholderRepo.Query(keywords, status, pageNo, pageSize, projectId)
	return
}

func (s *NluPlaceholderService) Get(id uint) (po model.NluPlaceholder) {
	po = s.NluPlaceholderRepo.Get(id)
	return
}

func (s *NluPlaceholderService) Save(po *model.NluPlaceholder) (err error) {
	err = s.NluPlaceholderRepo.Save(po)

	return
}

func (s *NluPlaceholderService) Update(po *model.NluPlaceholder) (err error) {
	err = s.NluPlaceholderRepo.Update(po)

	return
}

func (s *NluPlaceholderService) Disable(id uint) (err error) {
	err = s.NluPlaceholderRepo.Disable(id)

	return
}

func (s *NluPlaceholderService) Delete(id uint) (err error) {
	err = s.NluPlaceholderRepo.Delete(id)

	return
}

func (s *NluPlaceholderService) BatchDelete(ids []int) (err error) {
	err = s.NluPlaceholderRepo.BatchDelete(ids)

	return
}

func (s *NluPlaceholderService) Resort(srcId, targetId int) (err error) {
	err = s.NluPlaceholderRepo.Resort(srcId, targetId)

	return
}
