package service

import (
	"github.com/utlai/utl/internal/server/repo"
)

type NluDictService struct {
	NluSynonymRepo *repo.NluSynonymRepo `inject:""`
	NluLookupRepo  *repo.NluLookupRepo  `inject:""`
	NluRegexRepo   *repo.NluRegexRepo   `inject:""`
}

func NewNluDictService() *NluDictService {
	return &NluDictService{}
}

func (s *NluDictService) List(tp string) (pos []map[string]interface{}) {
	if tp == "synonym" {
		pos = s.NluSynonymRepo.List()
	} else if tp == "lookup" {
		pos = s.NluLookupRepo.List()
	} else if tp == "regex" {
		pos = s.NluRegexRepo.List()
	}

	return
}
