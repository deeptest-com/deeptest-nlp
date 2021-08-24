package repo

import (
	_const "github.com/utlai/utl/internal/pkg/const"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/utlai/utl/internal/server/model"
	"gorm.io/gorm"
)

type NluHistoryRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewNluHistoryRepo() *NluHistoryRepo {
	return &NluHistoryRepo{}
}

func (r *NluHistoryRepo) ListByProjectId(projectId uint) (pos []model.NluHistory) {
	query := r.DB.Model(&model.NluHistory{}).
		Where("NOT deleted AND NOT disabled").
		Where("project_id = ?", projectId).
		Order("id DESC").
		Limit(_const.PageSize)

	err := query.Find(&pos).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *NluHistoryRepo) Save(po *model.NluHistory) (err error) {
	err = r.DB.Model(&po).Omit("").Create(&po).Error
	return
}
