package repo

import (
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/utlai/utl/internal/server/model"
	"gorm.io/gorm"
	"time"
)

type NluIntentRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewNluIntentRepo() *NluIntentRepo {
	return &NluIntentRepo{}
}

func (r *NluIntentRepo) Query(keywords, status string, pageNo int, pageSize int) (pos []model.NluIntent, total int64) {
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
	query = query.Where("deleted_at IS NULL")

	err := query.Find(&pos).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}
	err = r.DB.Model(&model.NluIntent{}).Count(&total).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *NluIntentRepo) ListByTaskId(taskId uint) (pos []model.NluIntent) {
	query := r.DB.Select("*").
		Where("deleted_at IS NULL").
		Where("task_id = ?", taskId).
		Order("id ASC")

	err := query.Find(&pos).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *NluIntentRepo) Get(id uint) (po model.NluIntent) {
	r.DB.Where("id = ?", id).First(&po)
	return
}

func (r *NluIntentRepo) Save(po *model.NluIntent) (err error) {
	err = r.DB.Model(&po).Omit("").Create(&po).Error
	return
}

func (r *NluIntentRepo) Update(po *model.NluIntent) (err error) {
	err = r.DB.Omit("").Save(&po).Error
	return
}

func (r *NluIntentRepo) SetDefault(id uint) (err error) {
	r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.DB.Model(&model.NluIntent{}).Where("id = ?", id).
			Updates(map[string]interface{}{"is_default": true}).Error
		if err != nil {
			return err
		}

		err = r.DB.Model(&model.NluIntent{}).Where("id != ?", id).
			Updates(map[string]interface{}{"is_default": false}).Error

		return nil
	})

	return
}

func (r *NluIntentRepo) Disable(id uint) (err error) {
	err = r.DB.Model(&model.NluIntent{}).Where("id = ?", id).
		Updates(map[string]interface{}{"disabled": gorm.Expr("NOT disabled")}).Error

	return
}

func (r *NluIntentRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.NluIntent{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted_at": time.Now()}).Error

	return
}

func (r *NluIntentRepo) BatchDelete(ids []int) (err error) {
	err = r.DB.Model(&model.NluIntent{}).Where("id IN (?)", ids).
		Updates(map[string]interface{}{"deleted_at": time.Now()}).Error

	return
}
