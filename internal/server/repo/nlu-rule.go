package repo

import (
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/utlai/utl/internal/server/model"
	"gorm.io/gorm"
	"time"
)

type NluRuleRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewNluRuleRepo() *NluRuleRepo {
	return &NluRuleRepo{}
}

func (r *NluRuleRepo) Query(keywords, status string, pageNo int, pageSize int) (pos []model.NluRule, total int64) {
	query := r.DB.Model(&model.NluRule{}).Where("NOT deleted").Order("id ASC")
	if status == "true" {
		query = query.Where("NOT disabled")
	} else if status == "false" {
		query = query.Where("disabled")
	}

	if keywords != "" {
		query = query.Where("name LIKE ?", "%"+keywords+"%")
	}
	if pageNo > 0 {
		query = query.Offset((pageNo - 1) * pageSize).Limit(pageSize)
	}

	err := query.Find(&pos).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	err = query.Offset(-1).Limit(-1).Count(&total).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *NluRuleRepo) ListByIntentId(intentId uint) (pos []model.NluRule) {
	query := r.DB.Select("*").
		Where("NOT deleted").
		Where("intent_id = ?", intentId).
		Order("id ASC")

	err := query.Find(&pos).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}
func (r *NluRuleRepo) ProjectId(projectId uint) (pos []model.NluRule) {
	query := r.DB.Select("*").
		Where("NOT deleted").
		Where("project_id = ?", projectId).
		Order("id ASC")

	err := query.Find(&pos).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *NluRuleRepo) Get(id uint) (po model.NluRule) {
	r.DB.Where("id = ?", id).First(&po)
	return
}

func (r *NluRuleRepo) GetWithSlots(id uint) (po model.NluRule) {
	r.DB.Preload("Slots").Where("id = ?", id).First(&po)
	return
}

func (r *NluRuleRepo) Save(po *model.NluRule) (err error) {
	err = r.DB.Omit("Slots").Create(&po).Error
	return
}

func (r *NluRuleRepo) Update(po *model.NluRule) (err error) {
	err = r.DB.Omit("Slots").Updates(&po).Error
	return
}

func (r *NluRuleRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.NluRule{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true, "deleted_at": time.Now()}).Error

	return
}

func (r *NluRuleRepo) Disable(id uint) (err error) {
	err = r.DB.Model(&model.NluRule{}).Where("id = ?", id).
		Updates(map[string]interface{}{"disabled": gorm.Expr("NOT disabled")}).Error

	return
}
