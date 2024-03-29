package serverService

import (
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/repo"
)

type NluRegexService struct {
	NluRegexRepo *repo.NluRegexRepo `inject:""`
}

func NewNluRegexService() *NluRegexService {
	return &NluRegexService{}
}

func (s *NluRegexService) List(keywords, status string, pageNo, pageSize, projectId int) (pos []model.NluRegex, total int64) {
	pos, total = s.NluRegexRepo.Query(keywords, status, pageNo, pageSize, projectId)
	return
}

func (s *NluRegexService) Get(id uint) (po model.NluRegex) {
	po = s.NluRegexRepo.Get(id)
	return
}

func (s *NluRegexService) Save(po *model.NluRegex) (err error) {
	err = s.NluRegexRepo.Save(po)

	return
}

func (s *NluRegexService) Update(po *model.NluRegex) (err error) {
	err = s.NluRegexRepo.Update(po)

	return
}

func (s *NluRegexService) Disable(id uint) (err error) {
	err = s.NluRegexRepo.Disable(id)

	return
}

func (s *NluRegexService) Delete(id uint) (err error) {
	err = s.NluRegexRepo.Delete(id)

	return
}

func (s *NluRegexService) BatchDelete(ids []int) (err error) {
	err = s.NluRegexRepo.BatchDelete(ids)

	return
}

func (s *NluRegexService) Resort(srcId, targetId int) (err error) {
	err = s.NluRegexRepo.Resort(srcId, targetId)

	return
}
