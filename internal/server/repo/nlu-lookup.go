package repo

import (
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/utlai/utl/internal/server/model"
	"gorm.io/gorm"
	"time"
)

type NluLookupRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewNluLookupRepo() *NluLookupRepo {
	return &NluLookupRepo{}
}

func (r *NluLookupRepo) Query(keywords, status string, pageNo int, pageSize int) (pos []model.NluLookup, total int64) {
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
	err = r.DB.Model(&model.NluLookup{}).Count(&total).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *NluLookupRepo) ListByProjectId(projectId uint) (pos []model.NluLookup) {
	query := r.DB.Select("*").
		Where("deleted_at IS NULL AND NOT disabled").
		Where("project_id = ?", projectId).
		Order("id ASC")

	err := query.Find(&pos).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *NluLookupRepo) Get(id uint) (po model.NluLookup) {
	r.DB.Where("id = ?", id).First(&po)
	return
}

func (r *NluLookupRepo) Save(po *model.NluLookup) (err error) {
	err = r.DB.Model(&po).Omit("").Create(&po).Error
	return
}

func (r *NluLookupRepo) Update(po *model.NluLookup) (err error) {
	err = r.DB.Omit("").Save(&po).Error
	return
}

func (r *NluLookupRepo) SetDefault(id uint) (err error) {
	r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.DB.Model(&model.NluLookup{}).Where("id = ?", id).
			Updates(map[string]interface{}{"is_default": true}).Error
		if err != nil {
			return err
		}

		err = r.DB.Model(&model.NluLookup{}).Where("id != ?", id).
			Updates(map[string]interface{}{"is_default": false}).Error

		return nil
	})

	return
}

func (r *NluLookupRepo) Disable(id uint) (err error) {
	err = r.DB.Model(&model.NluLookup{}).Where("id = ?", id).
		Updates(map[string]interface{}{"disabled": gorm.Expr("NOT disabled")}).Error

	return
}

func (r *NluLookupRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.NluLookup{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted_at": time.Now()}).Error

	return
}

func (r *NluLookupRepo) BatchDelete(ids []int) (err error) {
	err = r.DB.Model(&model.NluLookup{}).Where("id IN (?)", ids).
		Updates(map[string]interface{}{"deleted_at": time.Now()}).Error

	return
}

func (r *NluLookupRepo) List() (pos []map[string]interface{}) {
	err := r.DB.Model(&model.NluLookup{}).
		Where("deleted_at IS NULL").
		Order("id ASC").
		Find(&pos).
		Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}
