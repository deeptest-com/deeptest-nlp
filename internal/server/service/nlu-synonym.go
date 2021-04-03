package service

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

func (s *NluSynonymService) List(keywords string, pageNo int, pageSize int) (pos []model.NluSynonym, total int64) {
	pos, total = s.NluSynonymRepo.Query(keywords, pageNo, pageSize)
	return
}

func (s *NluSynonymService) Save(po *model.NluSynonym) (err error) {
	err = s.NluSynonymRepo.Save(po)

	return
}
