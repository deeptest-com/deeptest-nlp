package repo

import (
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/utlai/utl/internal/server/model"
	"gorm.io/gorm"
	"time"
)

type NluSynonymItemRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewNluSynonymItemRepo() *NluSynonymItemRepo {
	return &NluSynonymItemRepo{}
}

func (r *NluSynonymItemRepo) Query(synonymId int, keywords, status string, pageNo int, pageSize int) (pos []model.NluSynonymItem, total int64) {
	query := r.DB.Select("*").Where("NOT deleted").Order("id ASC")
	query = query.Where("synonym_id = ?", synonymId)

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

func (r *NluSynonymItemRepo) ListBySynonymId(synonymId uint) (pos []model.NluSynonymItem) {
	query := r.DB.Select("*").
		Where("NOT deleted AND NOT disabled").
		Where("synonym_id = ?", synonymId).
		Order("id ASC")

	err := query.Find(&pos).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *NluSynonymItemRepo) Get(id uint) (po model.NluSynonymItem) {
	r.DB.Where("id = ?", id).First(&po)
	return
}

func (r *NluSynonymItemRepo) Save(po *model.NluSynonymItem) (err error) {
	err = r.DB.Model(&po).Omit("").Create(&po).Error
	return
}

func (r *NluSynonymItemRepo) Update(po *model.NluSynonymItem) (err error) {
	err = r.DB.Omit("").Save(&po).Error
	return
}

func (r *NluSynonymItemRepo) SetDefault(id uint) (err error) {
	r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.DB.Model(&model.NluSynonymItem{}).Where("id = ?", id).
			Updates(map[string]interface{}{"is_default": true}).Error
		if err != nil {
			return err
		}

		err = r.DB.Model(&model.NluSynonymItem{}).Where("id != ?", id).
			Updates(map[string]interface{}{"is_default": false}).Error

		return nil
	})

	return
}

func (r *NluSynonymItemRepo) Disable(id uint) (err error) {
	err = r.DB.Model(&model.NluSynonymItem{}).Where("id = ?", id).
		Updates(map[string]interface{}{"disabled": gorm.Expr("NOT disabled")}).Error

	return
}

func (r *NluSynonymItemRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.NluSynonymItem{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true, "deleted_at": time.Now()}).Error

	return
}

func (r *NluSynonymItemRepo) BatchDelete(ids []int) (err error) {
	err = r.DB.Model(&model.NluSynonymItem{}).Where("id IN (?)", ids).
		Updates(map[string]interface{}{"deleted": true, "deleted_at": time.Now()}).Error

	return
}

func (r *NluSynonymItemRepo) List() (pos []map[string]interface{}) {
	err := r.DB.Model(&model.NluSynonymItem{}).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).
		Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}
