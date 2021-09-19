package repo

import (
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/utlai/utl/internal/server/model"
	"gorm.io/gorm"
	"time"
)

type NluTaskRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewNluTaskRepo() *NluTaskRepo {
	return &NluTaskRepo{}
}

func (r *NluTaskRepo) Query(projectId int, keywords, status string, pageNo int, pageSize int) (pos []model.NluTask, total int64) {
	query := r.DB.Model(&model.NluTask{}).
		Where("NOT deleted").
		Order("ordr ASC")

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

	err := query.Find(&pos).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	err = query.Offset(-1).Limit(-1).Count(&total).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return

	return
}

func (r *NluTaskRepo) ListByProjectId(projectId uint) (pos []model.NluTask) {
	query := r.DB.Select("*").
		Where("NOT deleted AND NOT disabled").
		Where("project_id = ?", projectId).
		Order("ordr ASC")

	err := query.Find(&pos).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *NluTaskRepo) Get(id uint) (po model.NluTask) {
	r.DB.Where("id = ?", id).First(&po)
	return
}

func (r *NluTaskRepo) Save(po *model.NluTask) (err error) {
	err = r.DB.Model(&po).Omit("").Create(&po).Error
	return
}

func (r *NluTaskRepo) Update(po *model.NluTask) (err error) {
	err = r.DB.Omit("").Save(&po).Error
	return
}

func (r *NluTaskRepo) SetDefault(id uint) (err error) {
	r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.DB.Model(&model.NluTask{}).Where("id = ?", id).
			Updates(map[string]interface{}{"is_default": true}).Error
		if err != nil {
			return err
		}

		err = r.DB.Model(&model.NluTask{}).Where("id != ?", id).
			Updates(map[string]interface{}{"is_default": false}).Error

		return nil
	})

	return
}

func (r *NluTaskRepo) Disable(id uint) (err error) {
	err = r.DB.Model(&model.NluTask{}).Where("id = ?", id).
		Updates(map[string]interface{}{"disabled": gorm.Expr("NOT disabled")}).Error

	return
}

func (r *NluTaskRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.NluTask{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true, "deleted_at": time.Now()}).Error

	return
}

func (r *NluTaskRepo) BatchDelete(ids []int) (err error) {
	err = r.DB.Model(&model.NluTask{}).Where("id IN (?)", ids).
		Updates(map[string]interface{}{"deleted": true, "deleted_at": time.Now()}).Error

	return
}

func (r *NluTaskRepo) Resort(srcId, targetId, parentId int) (err error) {
	target := r.Get(uint(targetId))

	err = r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.DB.Model(&model.NluTask{}).Where("project_id = ? AND ordr >= ?", parentId, target.Ordr).
			Updates(map[string]interface{}{"ordr": gorm.Expr("ordr + 1")}).Error
		if err != nil {
			return err
		}

		err = r.DB.Model(&model.NluTask{}).Where("project_id = ? AND id = ?", parentId, srcId).
			Updates(map[string]interface{}{"ordr": target.Ordr}).Error
		if err != nil {
			return err
		}

		return nil
	})

	return
}
