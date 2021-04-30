package service

import (
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/repo"
)

type NluTaskService struct {
	NluTaskRepo *repo.NluTaskRepo `inject:""`
}

func NewNluTaskService() *NluTaskService {
	return &NluTaskService{}
}

func (s *NluTaskService) List(keywords, status string, pageNo int, pageSize int) (pos []model.NluTask, total int64) {
	pos, total = s.NluTaskRepo.Query(keywords, status, pageNo, pageSize)
	return
}

func (s *NluTaskService) Get(id uint) (po model.NluTask) {
	po = s.NluTaskRepo.Get(id)
	return
}

func (s *NluTaskService) Save(po *model.NluTask) (err error) {
	err = s.NluTaskRepo.Save(po)

	return
}

func (s *NluTaskService) Update(po *model.NluTask) (err error) {
	err = s.NluTaskRepo.Update(po)

	return
}

func (s *NluTaskService) SetDefault(id uint) (err error) {
	err = s.NluTaskRepo.SetDefault(id)

	return
}

func (s *NluTaskService) Disable(id uint) (err error) {
	err = s.NluTaskRepo.Disable(id)

	return
}

func (s *NluTaskService) Delete(id uint) (err error) {
	err = s.NluTaskRepo.Delete(id)

	return
}

func (s *NluTaskService) BatchDelete(ids []int) (err error) {
	err = s.NluTaskRepo.BatchDelete(ids)

	return
}
