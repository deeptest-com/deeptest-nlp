package repo

import (
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/utlai/utl/internal/server/model"
	"gorm.io/gorm"
	"strings"
	"time"
)

type NluSentRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewNluSentRepo() *NluSentRepo {
	return &NluSentRepo{}
}

func (r *NluSentRepo) Query(keywords, status string, pageNo int, pageSize int) (pos []model.NluSent, total int64) {
	query := r.DB.Model(&model.NluSent{}).Where("NOT deleted").Order("id ASC")
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
}

func (r *NluSentRepo) ListByIntentId(intentId uint, sort string) (pos []model.NluSent) {
	if sort == "" {
		sort = "id"
	}
	query := r.DB.Select("*").
		Where("NOT deleted AND NOT disabled").
		Where("intent_id = ?", intentId).
		Order(sort + " ASC")

	err := query.Find(&pos).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *NluSentRepo) Get(id uint) (po model.NluSent) {
	r.DB.Where("id = ?", id).First(&po)
	return
}

func (r *NluSentRepo) GetWithSlots(id uint) (po model.NluSent) {
	r.DB.Preload("Slots").Where("id = ?", id).First(&po)
	return
}

func (r *NluSentRepo) Save(po *model.NluSent) (err error) {
	po.Html = strings.TrimSpace(po.Html)
	po.Ordr = r.GetMaxOrder(po.IntentId)

	err = r.DB.Omit("Slots").Create(&po).Error

	for i := 0; i < len(po.Slots); i++ {
		po.Slots[i].ID = 0
		po.Slots[i].SentRefer = po.ID
	}
	err = r.DB.CreateInBatches(po.Slots, len(po.Slots)).Error
	return
}

func (r *NluSentRepo) Update(po *model.NluSent) (err error) {
	po.Html = strings.TrimSpace(po.Html)

	err = r.DB.Omit("Slots").Updates(&po).Error

	for i := 0; i < len(po.Slots); i++ {
		po.Slots[i].ID = 0
		po.Slots[i].SentRefer = po.ID
	}

	err = r.DB.Where("sent_refer = ?", po.ID).Delete(&model.NluSlot{}).Error
	err = r.DB.CreateInBatches(po.Slots, len(po.Slots)).Error

	return
}

func (r *NluSentRepo) Disable(id uint) (err error) {
	err = r.DB.Model(&model.NluSent{}).Where("id = ?", id).
		Updates(map[string]interface{}{"disabled": gorm.Expr("NOT disabled")}).Error

	return
}

func (r *NluSentRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.NluSent{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true, "deleted_at": time.Now()}).Error

	return
}

func (r *NluSentRepo) ListByIntent(intentId uint) (pos []model.NluSent) {
	r.DB.Where("intent_id = ?", intentId).
		Where("NOT deleted").
		Order("ordr ASC").
		Find(&pos)

	return
}

func (r *NluSentRepo) Resort(srcId, targetId, intentId int) (err error) {
	target := r.Get(uint(targetId))

	err = r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.DB.Model(&model.NluSent{}).Where("intent_id = ? AND ordr >= ?", intentId, target.Ordr).
			Updates(map[string]interface{}{"ordr": gorm.Expr("ordr + 1")}).Error
		if err != nil {
			return err
		}

		err = r.DB.Model(&model.NluSent{}).Where("intent_id = ? AND id = ?", intentId, srcId).
			Updates(map[string]interface{}{"ordr": target.Ordr}).Error
		if err != nil {
			return err
		}

		return nil
	})

	return
}

func (r *NluSentRepo) GetMaxOrder(parentId uint) (order int) {
	var po model.NluSent
	r.DB.Model(&po).Where("intent_id = ?", parentId).
		Where("NOT deleted").
		Order("ordr DESC").
		Limit(1).
		First(&po)

	if po.ID > 0 {
		order = po.Ordr + 1
	}
	return
}
