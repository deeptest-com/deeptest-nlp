package repo

import (
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/utlai/utl/internal/server/model"
	"gorm.io/gorm"
	"time"
)

type NluRegexItemRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewNluRegexItemRepo() *NluRegexItemRepo {
	return &NluRegexItemRepo{}
}

func (r *NluRegexItemRepo) Query(regexId int, keywords, status string, pageNo int, pageSize int) (pos []model.NluRegexItem, total int64) {
	query := r.DB.Select("*").Order("id ASC")
	query = query.Where("regex_id = ?", regexId)

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
	query = query.Where("deleted_at IS NULL")

	err := query.Find(&pos).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}
	err = r.DB.Model(&model.NluRegexItem{}).Count(&total).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *NluRegexItemRepo) ListByRegexId(regexId uint) (pos []model.NluRegexItem) {
	query := r.DB.Select("*").
		Where("deleted_at IS NULL AND NOT disabled").
		Where("regex_id = ?", regexId).
		Order("id ASC")

	err := query.Find(&pos).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *NluRegexItemRepo) Get(id uint) (po model.NluRegexItem) {
	r.DB.Where("id = ?", id).First(&po)
	return
}

func (r *NluRegexItemRepo) Save(po *model.NluRegexItem) (err error) {
	err = r.DB.Model(&po).Omit("").Create(&po).Error
	return
}

func (r *NluRegexItemRepo) Update(po *model.NluRegexItem) (err error) {
	err = r.DB.Omit("").Save(&po).Error
	return
}

func (r *NluRegexItemRepo) SetDefault(id uint) (err error) {
	r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.DB.Model(&model.NluRegexItem{}).Where("id = ?", id).
			Updates(map[string]interface{}{"is_default": true}).Error
		if err != nil {
			return err
		}

		err = r.DB.Model(&model.NluRegexItem{}).Where("id != ?", id).
			Updates(map[string]interface{}{"is_default": false}).Error

		return nil
	})

	return
}

func (r *NluRegexItemRepo) Disable(id uint) (err error) {
	err = r.DB.Model(&model.NluRegexItem{}).Where("id = ?", id).
		Updates(map[string]interface{}{"disabled": gorm.Expr("NOT disabled")}).Error

	return
}

func (r *NluRegexItemRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.NluRegexItem{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted_at": time.Now()}).Error

	return
}

func (r *NluRegexItemRepo) BatchDelete(ids []int) (err error) {
	err = r.DB.Model(&model.NluRegexItem{}).Where("id IN (?)", ids).
		Updates(map[string]interface{}{"deleted_at": time.Now()}).Error

	return
}

func (r *NluRegexItemRepo) List() (pos []map[string]interface{}) {
	err := r.DB.Model(&model.NluRegexItem{}).
		Where("deleted_at IS NULL").
		Order("id ASC").
		Find(&pos).
		Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}
