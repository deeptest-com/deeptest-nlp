package repo

import (
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/utlai/utl/internal/server/model"
	"gorm.io/gorm"
)

type NluLookupRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewNluLookupRepo() *NluLookupRepo {
	return &NluLookupRepo{}
}

func (r *NluLookupRepo) Query(keywords string, pageNo int, pageSize int) (pos []model.NluLookup, total int64) {
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
	err = r.DB.Model(&model.NluLookup{}).Count(&total).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *NluLookupRepo) Save(po *model.NluLookup) (err error) {
	err = r.DB.Model(&po).Omit("").Create(&po).Error
	return
}
