package repo

import (
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/utlai/utl/internal/server/model"
	"gorm.io/gorm"
	"time"
)

type NluSlotRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewNluSlotRepo() *NluSlotRepo {
	return &NluSlotRepo{}
}

func (r *NluSlotRepo) Query(keywords, status string, pageNo int, pageSize int) (pos []model.NluSlot, total int64) {
	query := r.DB.Select("*").Where("NOT deleted").Order("id ASC")
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

	err = query.Count(&total).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *NluSlotRepo) ListBySentId(sentId uint) (pos []model.NluSlot) {
	query := r.DB.Select("*").
		Where("NOT deleted AND NOT disabled").
		Where("sent_refer = ?", sentId).
		Order("id ASC")

	err := query.Find(&pos).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *NluSlotRepo) Get(id uint) (po model.NluSlot) {
	r.DB.Where("id = ?", id).First(&po)
	return
}

func (r *NluSlotRepo) Save(po *model.NluSlot) (err error) {
	err = r.DB.Model(&po).Omit("").Create(&po).Error
	return
}

func (r *NluSlotRepo) Update(po *model.NluSlot) (err error) {
	err = r.DB.Omit("").Save(&po).Error
	return
}

func (r *NluSlotRepo) SetDefault(id uint) (err error) {
	r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.DB.Model(&model.NluSlot{}).Where("id = ?", id).
			Updates(map[string]interface{}{"is_default": true}).Error
		if err != nil {
			return err
		}

		err = r.DB.Model(&model.NluSlot{}).Where("id != ?", id).
			Updates(map[string]interface{}{"is_default": false}).Error

		return nil
	})

	return
}

func (r *NluSlotRepo) Disable(id uint) (err error) {
	err = r.DB.Model(&model.NluSlot{}).Where("id = ?", id).
		Updates(map[string]interface{}{"disabled": gorm.Expr("NOT disabled")}).Error

	return
}

func (r *NluSlotRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.NluSlot{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true, "deleted_at": time.Now()}).Error

	return
}

func (r *NluSlotRepo) BatchDelete(ids []int) (err error) {
	err = r.DB.Model(&model.NluSlot{}).Where("id IN (?)", ids).
		Updates(map[string]interface{}{"deleted": true, "deleted_at": time.Now()}).Error

	return
}
