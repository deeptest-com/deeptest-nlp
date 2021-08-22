package repo

import (
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/utlai/utl/internal/server/model"
	"gorm.io/gorm"
	"strings"
	"time"
)

type NluSynonymRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewNluSynonymRepo() *NluSynonymRepo {
	return &NluSynonymRepo{}
}

func (r *NluSynonymRepo) Query(keywords, status string, pageNo int, pageSize int) (pos []model.NluSynonym, total int64) {
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

func (r *NluSynonymRepo) ListByProjectId(projectId uint) (pos []model.NluSynonym) {
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

func (r *NluSynonymRepo) Get(id uint) (po model.NluSynonym) {
	r.DB.Where("id = ?", id).First(&po)
	return
}
func (r *NluSynonymRepo) GetByName(name string) (po model.NluSynonym) {
	r.DB.Where("name = ?", name).First(&po)
	return
}

func (r *NluSynonymRepo) Save(po *model.NluSynonym) (err error) {
	po.Name = strings.TrimSpace(po.Name)
	err = r.DB.Model(&po).Omit("").Create(&po).Error
	return
}

func (r *NluSynonymRepo) Update(po *model.NluSynonym) (err error) {
	po.Name = strings.TrimSpace(po.Name)
	err = r.DB.Omit("").Save(&po).Error
	return
}

func (r *NluSynonymRepo) SetDefault(id uint) (err error) {
	r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.DB.Model(&model.NluSynonym{}).Where("id = ?", id).
			Updates(map[string]interface{}{"is_default": true}).Error
		if err != nil {
			return err
		}

		err = r.DB.Model(&model.NluSynonym{}).Where("id != ?", id).
			Updates(map[string]interface{}{"is_default": false}).Error

		return nil
	})

	return
}

func (r *NluSynonymRepo) Disable(id uint) (err error) {
	err = r.DB.Model(&model.NluSynonym{}).Where("id = ?", id).
		Updates(map[string]interface{}{"disabled": gorm.Expr("NOT disabled")}).Error

	return
}

func (r *NluSynonymRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.NluSynonym{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true, "deleted_at": time.Now()}).Error

	return
}

func (r *NluSynonymRepo) BatchDelete(ids []int) (err error) {
	err = r.DB.Model(&model.NluSynonym{}).Where("id IN (?)", ids).
		Updates(map[string]interface{}{"deleted": true, "deleted_at": time.Now()}).Error

	return
}

func (r *NluSynonymRepo) List() (pos []map[string]interface{}) {
	err := r.DB.Model(&model.NluSynonym{}).
		Where("NOT deleted").
		Order("id ASC").
		Find(&pos).
		Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}
