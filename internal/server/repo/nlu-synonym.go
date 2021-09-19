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
	query := r.DB.Model(&model.NluSynonym{}).Where("NOT deleted")
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
func (r *NluSynonymRepo) GetByCode(code string) (po model.NluSynonym) {
	r.DB.Where("code = ? AND NOT deleted", code).First(&po)
	return
}

func (r *NluSynonymRepo) Save(po *model.NluSynonym) (err error) {
	po.Name = strings.TrimSpace(po.Name)
	po.Ordr = r.GetMaxOrder()

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
		Order("ordr ASC").
		Find(&pos).
		Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *NluSynonymRepo) Resort(srcId, targetId int) (err error) {
	target := r.Get(uint(targetId))

	err = r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.DB.Model(&model.NluSynonym{}).Where("ordr >= ?", target.Ordr).
			Updates(map[string]interface{}{"ordr": gorm.Expr("ordr + 1")}).Error
		if err != nil {
			return err
		}

		err = r.DB.Model(&model.NluSynonym{}).Where("id = ?", srcId).
			Updates(map[string]interface{}{"ordr": target.Ordr}).Error
		if err != nil {
			return err
		}

		return nil
	})

	return
}

func (r *NluSynonymRepo) GetMaxOrder() (order int) {
	var po model.NluSynonym
	r.DB.Model(&po).Where("NOT deleted").
		Order("ordr DESC").
		Limit(1).
		First(&po)

	if po.ID > 0 {
		order = po.Ordr + 1
	}
	return
}
