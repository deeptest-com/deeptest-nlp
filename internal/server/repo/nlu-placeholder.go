package repo

import (
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/utlai/utl/internal/server/model"
	"gorm.io/gorm"
	"strings"
	"time"
)

type NluPlaceholderRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewNluPlaceholderRepo() *NluPlaceholderRepo {
	return &NluPlaceholderRepo{}
}

func (r *NluPlaceholderRepo) Query(keywords, status string, pageNo, pageSize, projectId int) (pos []model.NluPlaceholder, total int64) {
	query := r.DB.Model(&model.NluPlaceholder{}).Where("NOT deleted")

	if projectId > 0 {
		query = query.Where("project_id = ?", projectId)
	}

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

func (r *NluPlaceholderRepo) ListByProjectId(projectId uint) (pos []model.NluPlaceholder) {
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

func (r *NluPlaceholderRepo) Get(id uint) (po model.NluPlaceholder) {
	r.DB.Where("id = ?", id).First(&po)
	return
}
func (r *NluPlaceholderRepo) GetByCode(code string) (po model.NluPlaceholder) {
	r.DB.Where("code = ? AND NOT deleted", code).First(&po)
	return
}

func (r *NluPlaceholderRepo) Save(po *model.NluPlaceholder) (err error) {
	po.Name = strings.TrimSpace(po.Name)
	po.Ordr = r.GetMaxOrder()
	err = r.DB.Model(&po).Omit("").Create(&po).Error
	return
}

func (r *NluPlaceholderRepo) Update(po *model.NluPlaceholder) (err error) {
	po.Name = strings.TrimSpace(po.Name)
	err = r.DB.Omit("").Save(&po).Error
	return
}

func (r *NluPlaceholderRepo) Disable(id uint) (err error) {
	err = r.DB.Model(&model.NluPlaceholder{}).Where("id = ?", id).
		Updates(map[string]interface{}{"disabled": gorm.Expr("NOT disabled")}).Error

	return
}

func (r *NluPlaceholderRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.NluPlaceholder{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true, "deleted_at": time.Now()}).Error

	return
}

func (r *NluPlaceholderRepo) BatchDelete(ids []int) (err error) {
	err = r.DB.Model(&model.NluPlaceholder{}).Where("id IN (?)", ids).
		Updates(map[string]interface{}{"deleted": true, "deleted_at": time.Now()}).Error

	return
}

func (r *NluPlaceholderRepo) List() (pos []map[string]interface{}) {
	err := r.DB.Model(&model.NluPlaceholder{}).
		Where("NOT deleted").
		Order("ordr ASC").
		Find(&pos).
		Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *NluPlaceholderRepo) Resort(srcId, targetId int) (err error) {
	target := r.Get(uint(targetId))

	err = r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.DB.Model(&model.NluPlaceholder{}).Where("ordr >= ?", target.Ordr).
			Updates(map[string]interface{}{"ordr": gorm.Expr("ordr + 1")}).Error
		if err != nil {
			return err
		}

		err = r.DB.Model(&model.NluPlaceholder{}).Where("id = ?", srcId).
			Updates(map[string]interface{}{"ordr": target.Ordr}).Error
		if err != nil {
			return err
		}

		return nil
	})

	return
}

func (r *NluPlaceholderRepo) GetMaxOrder() (order int) {
	var po model.NluPlaceholder
	r.DB.Model(&po).Where("NOT deleted").
		Order("ordr DESC").
		Limit(1).
		First(&po)

	if po.ID > 0 {
		order = po.Ordr + 1
	}
	return
}
