package repo

import (
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/utlai/utl/internal/server/model"
	"gorm.io/gorm"
)

type NluSynonymRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewNluSynonymRepo() *NluSynonymRepo {
	return &NluSynonymRepo{}
}

func (r *NluSynonymRepo) Query(keywords string, pageNo int, pageSize int) (pos []model.NluSynonym, total int64) {
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
	err = r.DB.Model(&model.NluSynonym{}).Count(&total).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *NluSynonymRepo) Save(po *model.NluSynonym) (err error) {
	err = r.DB.Model(&po).Omit("").Create(&po).Error
	return
}
