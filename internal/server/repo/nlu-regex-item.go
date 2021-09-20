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
	query := r.DB.Model(&model.NluRegexItem{})
	query = query.Where("regex_id = ?", regexId).Where("NOT deleted")

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

	query.Order("ordr ASC")
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

func (r *NluRegexItemRepo) ListByRegexId(regexId uint) (pos []model.NluRegexItem) {
	query := r.DB.Select("*").
		Where("NOT deleted AND NOT disabled").
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
	po.Ordr = r.GetMaxOrder(po.RegexId)
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
		Updates(map[string]interface{}{"deleted": true, "deleted_at": time.Now()}).Error

	return
}

func (r *NluRegexItemRepo) BatchDelete(ids []int) (err error) {
	err = r.DB.Model(&model.NluRegexItem{}).Where("id IN (?)", ids).
		Updates(map[string]interface{}{"deleted": true, "deleted_at": time.Now()}).Error

	return
}

func (r *NluRegexItemRepo) List() (pos []map[string]interface{}) {
	err := r.DB.Model(&model.NluRegexItem{}).
		Where("NOT deleted").
		Order("ordr ASC").
		Find(&pos).
		Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *NluRegexItemRepo) Resort(srcId, targetId, parentId int) (err error) {
	target := r.Get(uint(targetId))

	err = r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.DB.Model(&model.NluRegexItem{}).Where("regex_id = ? AND ordr >= ?", parentId, target.Ordr).
			Updates(map[string]interface{}{"ordr": gorm.Expr("ordr + 1")}).Error
		if err != nil {
			return err
		}

		err = r.DB.Model(&model.NluRegexItem{}).Where("regex_id = ? AND id = ?", parentId, srcId).
			Updates(map[string]interface{}{"ordr": target.Ordr}).Error
		if err != nil {
			return err
		}

		return nil
	})

	return
}

func (r *NluRegexItemRepo) GetMaxOrder(parentId uint) (order int) {
	var po model.NluRegexItem
	r.DB.Model(&po).Where("regex_id = ?", parentId).
		Where("NOT deleted").
		Order("ordr DESC").
		Limit(1).
		First(&po)

	if po.ID > 0 {
		order = po.Ordr + 1
	}
	return
}

func (r *NluRegexItemRepo) GetByCode(code string) (po model.NluRegexItem) {
	r.DB.Model(&po).Where("code = ? AND NOT deleted", code).First(&po)
	return
}
