package repo

import (
	_const "github.com/utlai/utl/internal/pkg/const"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	"github.com/utlai/utl/internal/server/model"
	"gorm.io/gorm"
	"time"
)

func NewProjectRepo() *ProjectRepo {
	return &ProjectRepo{}
}

type ProjectRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func (r *ProjectRepo) Query(keywords string, pageNo int, pageSize int) (models []model.Project, total int64) {
	query := r.DB.Select("*").Order("id ASC")
	if keywords != "" {
		query = query.Where("name LIKE ?", "%"+keywords+"%")
	}
	if pageNo > 0 {
		query = query.Offset((pageNo) * pageSize).Limit(pageSize)
	}

	err := query.Find(&models).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}
	err = r.DB.Model(&model.Project{}).Count(&total).Error
	if err != nil {
		_logUtils.Errorf("sql error %s", err.Error())
	}

	return
}

func (r *ProjectRepo) Save(project *model.Project) (err error) {
	err = r.DB.Model(&project).Omit("").Create(&project).Error
	return
}

func (r *ProjectRepo) SetProgress(projectId uint, progress _const.BuildProgress) (err error) {
	var data map[string]interface{}
	if progress == _const.ProgressInProgress {
		data = map[string]interface{}{"progress": progress, "start_time": time.Now()}
	} else {
		data = map[string]interface{}{"progress": progress, "pending_time": time.Now()}
	}

	r.DB.Model(model.Project{}).Where("id=?", projectId).Updates(data)
	return
}

func (r *ProjectRepo) SetResult(projectId uint, progress _const.BuildProgress, status _const.BuildStatus) (err error) {
	var data = map[string]interface{}{"progress": progress, "result": status, "updatedTime": time.Now()}
	r.DB.Model(model.Project{}).Where("id=?", projectId).Updates(data)
	return
}
