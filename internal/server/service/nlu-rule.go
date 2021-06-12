package service

import (
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/repo"
)

type NluRuleService struct {
	NluRuleRepo *repo.NluRuleRepo `inject:""`
}

func NewNluRuleService() *NluRuleService {
	return &NluRuleService{}
}

func (s *NluRuleService) List(keywords, status string, pageNo int, pageSize int) (pos []model.NluRule, total int64) {
	pos, total = s.NluRuleRepo.Query(keywords, status, pageNo, pageSize)
	return
}
func (s *NluRuleService) ListByIntent(intentId uint) (pos []model.NluRule) {
	pos = s.NluRuleRepo.ListByIntent(intentId)
	return
}

func (s *NluRuleService) Get(id uint) (po model.NluRule) {
	po = s.NluRuleRepo.Get(id)
	return
}

func (s *NluRuleService) GetWithSlots(id uint) (po model.NluRule) {
	po = s.NluRuleRepo.GetWithSlots(id)

	return
}

func (s *NluRuleService) Save(po *model.NluRule) (err error) {
	err = s.NluRuleRepo.Save(po)

	return
}

func (s *NluRuleService) Update(po *model.NluRule) (err error) {
	err = s.NluRuleRepo.Update(po)

	return
}

func (s *NluRuleService) Disable(id uint) (err error) {
	err = s.NluRuleRepo.Disable(id)

	return
}

func (s *NluRuleService) Delete(id uint) (err error) {
	err = s.NluRuleRepo.Delete(id)

	return
}
