package serverService

import (
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/repo"
)

type NluSynonymService struct {
	NluSynonymRepo *repo.NluSynonymRepo `inject:""`
}

func NewNluSynonymService() *NluSynonymService {
	return &NluSynonymService{}
}

func (s *NluSynonymService) List(keywords, status string, pageNo int, pageSize int) (pos []model.NluSynonym, total int64) {
	pos, total = s.NluSynonymRepo.Query(keywords, status, pageNo, pageSize)
	return
}

func (s *NluSynonymService) Get(id uint) (po model.NluSynonym) {
	po = s.NluSynonymRepo.Get(id)
	return
}

func (s *NluSynonymService) Save(po *model.NluSynonym) (err error) {
	err = s.NluSynonymRepo.Save(po)

	return
}

func (s *NluSynonymService) Update(po *model.NluSynonym) (err error) {
	err = s.NluSynonymRepo.Update(po)

	return
}

func (s *NluSynonymService) SetDefault(id uint) (err error) {
	err = s.NluSynonymRepo.SetDefault(id)

	return
}

func (s *NluSynonymService) Disable(id uint) (err error) {
	err = s.NluSynonymRepo.Disable(id)

	return
}

func (s *NluSynonymService) Delete(id uint) (err error) {
	err = s.NluSynonymRepo.Delete(id)

	return
}

func (s *NluSynonymService) BatchDelete(ids []int) (err error) {
	err = s.NluSynonymRepo.BatchDelete(ids)

	return
}
