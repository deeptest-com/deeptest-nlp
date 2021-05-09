package repo

import (
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/utlai/utl/internal/server/model"
	"gorm.io/gorm"
	"time"
)

type NluSentRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewNluSentRepo() *NluSentRepo {
	return &NluSentRepo{}
}

func (r *NluSentRepo) Query(keywords, status string, pageNo int, pageSize int) (pos []model.NluSent, total int64) {
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
	err = r.DB.Model(&model.NluSent{}).Count(&total).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *NluSentRepo) Get(id uint) (po model.NluSent) {
	r.DB.Where("id = ?", id).First(&po)
	return
}

func (r *NluSentRepo) GetWithSlots(id uint) (po model.NluSent) {
	r.DB.Preload("Slots").Where("id = ?", id).First(&po)
	return
}

func (r *NluSentRepo) Save(po *model.NluSent) (err error) {
	err = r.DB.Session(&gorm.Session{FullSaveAssociations: true}).Create(&po).Error
	return
}

func (r *NluSentRepo) Update(po *model.NluSent) (err error) {
	err = r.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&po).Error
	return
}

func (r *NluSentRepo) SetDefault(id uint) (err error) {
	r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.DB.Model(&model.NluSent{}).Where("id = ?", id).
			Updates(map[string]interface{}{"is_default": true}).Error
		if err != nil {
			return err
		}

		err = r.DB.Model(&model.NluSent{}).Where("id != ?", id).
			Updates(map[string]interface{}{"is_default": false}).Error

		return nil
	})

	return
}

func (r *NluSentRepo) Disable(id uint) (err error) {
	err = r.DB.Model(&model.NluSent{}).Where("id = ?", id).
		Updates(map[string]interface{}{"disabled": gorm.Expr("NOT disabled")}).Error

	return
}

func (r *NluSentRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.NluSent{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted_at": time.Now()}).Error

	return
}

func (r *NluSentRepo) BatchDelete(ids []int) (err error) {
	err = r.DB.Model(&model.NluSent{}).Where("id IN (?)", ids).
		Updates(map[string]interface{}{"deleted_at": time.Now()}).Error

	return
}

func (r *NluSentRepo) ListByIntent(intentId uint) (pos []model.NluSent) {
	r.DB.Where("intent_id = ?", intentId).Find(&pos)

	return
}
