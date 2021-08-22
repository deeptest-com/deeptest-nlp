package repo

import (
	"fmt"
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

func (r *NluIntentRepo) ListByTaskId(taskId uint) (pos []model.NluIntent) {
	query := r.DB.Select("*").
		Where("NOT deleted").
		Where("task_id = ?", taskId).
		Order("ordr ASC")

	err := query.Find(&pos).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}
func (r *NluIntentRepo) ListByTaskIdNoDisabled(taskId uint) (pos []model.NluIntent) {
	query := r.DB.Select("*").
		Where("NOT disabled AND NOT deleted").
		Where("task_id = ?", taskId).
		Order("ordr ASC")

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
	err = r.DB.Model(&model.NluIntent{}).Where("id = ?", po.ID).
		Updates(map[string]interface{}{"name": po.Name}).Error
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
		Updates(map[string]interface{}{"deleted": true, "deleted_at": time.Now()}).Error

	return
}

func (r *NluIntentRepo) BatchDelete(ids []int) (err error) {
	err = r.DB.Model(&model.NluIntent{}).Where("id IN (?)", ids).
		Updates(map[string]interface{}{"deleted": true, "deleted_at": time.Now()}).Error

	return
}

func (r *NluIntentRepo) AddOrderForTargetAndNext(srcId uint, targetOrder int, taskId uint) (err error) {
	sql := fmt.Sprintf(`UPDATE %s SET ordr = ordr + 1 WHERE ordr >= %d AND task_id = %d AND id!=%d`,
		(&model.NluIntent{}).TableName(), targetOrder, taskId, srcId)
	err = r.DB.Exec(sql).Error

	return
}

func (r *NluIntentRepo) AddOrderForNext(srcId uint, targetOrder int, taskId uint) (err error) {
	sql := fmt.Sprintf(`UPDATE %s SET ordr = ordr + 1 WHERE ordr > %d AND task_id = %d AND id!=%d`,
		(&model.NluIntent{}).TableName(), targetOrder, taskId, srcId)
	err = r.DB.Exec(sql).Error

	return
}

func (r *NluIntentRepo) UpdateOrd(field model.NluIntent) (err error) {
	err = r.DB.Model(&field).UpdateColumn("ordr", field.Ordr).Error

	return
}

func (r *NluIntentRepo) GetMaxOrder(taskId uint) (ordr int) {
	preChild := model.NluIntent{}
	err := r.DB.
		Where("task_id=?", taskId).
		Order("ordr DESC").Limit(1).
		First(&preChild).Error

	if err != nil {
		ordr = 1
	}
	ordr = preChild.Ordr + 1

	return
}
