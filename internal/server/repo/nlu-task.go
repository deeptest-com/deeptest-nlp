package repo

import (
	"fmt"
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

func (r *NluTaskRepo) Query(keywords, status string, pageNo int, pageSize int) (pos []model.NluTask, total int64) {
	sql := "SELECT t.*, p.name project_name FROM nlu_task t" +
		" LEFT JOIN biz_project p ON t.project_id = p.id" +
		" WHERE t.deleted_at IS NULL"

	if status == "true" {
		sql += " AND NOT t.disabled"
	} else if status == "false" {
		sql += " AND t.disabled"
	}

	if keywords != "" {
		sql += " AND t.name LIKE %" + keywords + "%"
	}
	if pageNo > 0 {
		sql += fmt.Sprintf(" LIMIT %d OFFSET %d", pageSize, (pageNo-1)*pageSize)
	}

	sql += " ORDER BY t.id ASC"
	err := r.DB.Raw(sql).Scan(&pos).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}
	err = r.DB.Model(&model.NluTask{}).Count(&total).Error
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
		Updates(map[string]interface{}{"deleted_at": time.Now()}).Error

	return
}

func (r *NluTaskRepo) BatchDelete(ids []int) (err error) {
	err = r.DB.Model(&model.NluTask{}).Where("id IN (?)", ids).
		Updates(map[string]interface{}{"deleted_at": time.Now()}).Error

	return
}
