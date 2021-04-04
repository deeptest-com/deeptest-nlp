package repo

import (
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/utlai/utl/internal/server/model"
	"gorm.io/gorm"
)

type ProjectRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewProjectRepo() *ProjectRepo {
	return &ProjectRepo{}
}

func (r *ProjectRepo) Query(keywords, status string, pageNo int, pageSize int) (pos []model.Project, total int64) {
	query := r.DB.Select("*").Order("id ASC")
	if status == "true" {
		query = query.Where("deleted_at IS NULL")
	} else if status == "false" {
		query = query.Where("deleted_at IS NOT NULL")
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
	err = r.DB.Model(&model.Project{}).Count(&total).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *ProjectRepo) Save(po *model.Project) (err error) {
	err = r.DB.Model(&po).Omit("").Create(&po).Error
	return
}

func (r *ProjectRepo) SetDefault(id uint) (err error) {

	return
}
