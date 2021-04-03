package repo

import (
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/utlai/utl/internal/server/model"
	"gorm.io/gorm"
)

type NluIntentRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewNluIntentRepo() *NluIntentRepo {
	return &NluIntentRepo{}
}

func (r *NluIntentRepo) Query(keywords string, pageNo int, pageSize int) (pos []model.NluIntent, total int64) {
	query := r.DB.Select("*").Order("id ASC")
	if keywords != "" {
		query = query.Where("name LIKE ?", "%"+keywords+"%")
	}
	if pageNo > 0 {
		query = query.Offset((pageNo) * pageSize).Limit(pageSize)
	}

	err := query.Find(&pos).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}
	err = r.DB.Model(&model.NluIntent{}).Count(&total).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *NluIntentRepo) Save(po *model.NluIntent) (err error) {
	err = r.DB.Model(&po).Omit("").Create(&po).Error
	return
}
