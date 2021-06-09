package repo

import (
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/utlai/utl/internal/server/model"
	"gorm.io/gorm"
	"time"
)

type NluRegexRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewNluRegexRepo() *NluRegexRepo {
	return &NluRegexRepo{}
}

func (r *NluRegexRepo) Query(keywords, status string, pageNo int, pageSize int) (pos []model.NluRegex, total int64) {
	query := r.DB.Select("*").Order("id ASC")
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
	query = query.Where("NOT deleted")

	err := query.Find(&pos).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}
	err = r.DB.Model(&model.NluRegex{}).Count(&total).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *NluRegexRepo) ListByProjectId(projectId uint) (pos []model.NluRegex) {
	query := r.DB.Select("*").
		Where("NOT deleted AND NOT disabled").
		//Where("project_id = ?", projectId).
		Order("id ASC")

	err := query.Find(&pos).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *NluRegexRepo) Get(id uint) (po model.NluRegex) {
	r.DB.Where("id = ?", id).First(&po)
	return
}

func (r *NluRegexRepo) Save(po *model.NluRegex) (err error) {
	err = r.DB.Model(&po).Omit("").Create(&po).Error
	return
}

func (r *NluRegexRepo) Update(po *model.NluRegex) (err error) {
	err = r.DB.Omit("").Save(&po).Error
	return
}

func (r *NluRegexRepo) SetDefault(id uint) (err error) {
	r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.DB.Model(&model.NluRegex{}).Where("id = ?", id).
			Updates(map[string]interface{}{"is_default": true}).Error
		if err != nil {
			return err
		}

		err = r.DB.Model(&model.NluRegex{}).Where("id != ?", id).
			Updates(map[string]interface{}{"is_default": false}).Error

		return nil
	})

	return
}

func (r *NluRegexRepo) Disable(id uint) (err error) {
	err = r.DB.Model(&model.NluRegex{}).Where("id = ?", id).
		Updates(map[string]interface{}{"disabled": gorm.Expr("NOT disabled")}).Error

	return
}

func (r *NluRegexRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.NluRegex{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true, "deleted_at": time.Now()}).Error

	return
}

func (r *NluRegexRepo) BatchDelete(ids []int) (err error) {
	err = r.DB.Model(&model.NluRegex{}).Where("id IN (?)", ids).
		Updates(map[string]interface{}{"deleted": true, "deleted_at": time.Now()}).Error

	return
}

func (r *NluRegexRepo) List() (pos []map[string]interface{}) {
	err := r.DB.Model(&model.NluRegex{}).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).
		Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}
