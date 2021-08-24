package repo

import (
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/utlai/utl/internal/server/model"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
	"gorm.io/gorm"
	"time"
)

type ProjectRepo struct {
	CommonRepo
	NluHistoryRepo *NluHistoryRepo `inject:""`
	DB             *gorm.DB        `inject:""`
}

func NewProjectRepo() *ProjectRepo {
	return &ProjectRepo{}
}

func (r *ProjectRepo) Query(keywords, status string, pageNo int, pageSize int) (pos []model.Project, total int64) {
	query := r.DB.Model(&model.Project{}).Where("NOT deleted").Order("id ASC")
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

func (r *ProjectRepo) Get(id uint) (po model.Project) {
	r.DB.Where("id = ?", id).First(&po)

	return
}

func (r *ProjectRepo) GetDetail(id uint) (po model.Project) {
	r.DB.Where("id = ?", id).First(&po)

	histories := r.NluHistoryRepo.ListByProjectId(id)

	if len(histories) > 0 {
		po.Histories = histories
		po.CreatedBy = histories[len(histories)-1].UserName
	}

	return
}

func (r *ProjectRepo) Save(po *model.Project) (err error) {
	err = r.DB.Model(&po).Omit("").Create(&po).Error
	return
}

func (r *ProjectRepo) Update(po *model.Project) (err error) {
	err = r.DB.Omit("").Save(&po).Error
	return
}

func (r *ProjectRepo) SetDefault(id uint) (err error) {
	r.DB.Transaction(func(tx *gorm.DB) error {
		err = r.DB.Model(&model.Project{}).Where("id = ?", id).
			Updates(map[string]interface{}{"is_default": true}).Error
		if err != nil {
			return err
		}

		err = r.DB.Model(&model.Project{}).Where("id != ?", id).
			Updates(map[string]interface{}{"is_default": false}).Error

		return nil
	})

	return
}

func (r *ProjectRepo) Disable(id uint) (err error) {
	err = r.DB.Model(&model.Project{}).Where("id = ?", id).
		Updates(map[string]interface{}{"disabled": gorm.Expr("NOT disabled")}).Error

	return
}

func (r *ProjectRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.Project{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true, "deleted_at": time.Now()}).Error

	return
}

func (r *ProjectRepo) StartService(id uint, port int) (err error) {
	err = r.DB.Model(&model.Project{}).Where("id = ?", id).
		Updates(map[string]interface{}{"service_status": serverConst.StartService, "service_port": port}).Error

	return
}
func (r *ProjectRepo) StopService(id uint) (err error) {
	err = r.DB.Model(&model.Project{}).Where("id = ?", id).
		Updates(map[string]interface{}{"service_status": serverConst.StopService, "service_port": 0}).Error

	return
}

func (r *ProjectRepo) StartTraining(id uint) (err error) {
	err = r.DB.Model(&model.Project{}).Where("id = ?", id).
		Updates(map[string]interface{}{"training_status": serverConst.StartTraining, "service_port": 0}).Error

	return
}

func (r *ProjectRepo) EndTraining(id uint) (err error) {
	err = r.DB.Model(&model.Project{}).Where("id = ?", id).
		Updates(map[string]interface{}{"training_status": serverConst.EndTraining, "service_port": 0}).Error

	return
}

func (r *ProjectRepo) CancelTraining(id uint) (err error) {
	err = r.DB.Model(&model.Project{}).Where("id = ?", id).
		Updates(map[string]interface{}{"service_status": serverConst.CancelTraining, "service_port": 0}).Error

	return
}
