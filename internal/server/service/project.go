package serverService

import (
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/repo"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
)

type ProjectService struct {
	ProjectRepo       *repo.ProjectRepo  `inject:""`
	NluHistoryService *NluHistoryService `inject:""`
}

func NewProjectService() *ProjectService {
	return &ProjectService{}
}

func (s *ProjectService) List(keywords, status string, pageNo int, pageSize int) (pos []model.Project, total int64) {
	pos, total = s.ProjectRepo.Query(keywords, status, pageNo, pageSize)
	return
}

func (s *ProjectService) Get(id uint) (po model.Project) {
	po = s.ProjectRepo.Get(id)
	return
}

func (s *ProjectService) GetDetail(id uint) (po model.Project) {
	po = s.ProjectRepo.GetDetail(id)
	return
}

func (s *ProjectService) Save(po *model.Project, userId uint) (err error) {
	po.UserId = userId
	err = s.ProjectRepo.Save(po)

	s.NluHistoryService.Add(userId, po.ID, serverConst.Create)

	return
}

func (s *ProjectService) Update(po *model.Project) (err error) {
	err = s.ProjectRepo.Update(po)

	return
}

func (s *ProjectService) SetDefault(id uint) (err error) {
	err = s.ProjectRepo.SetDefault(id)

	return
}

func (s *ProjectService) Disable(id uint) (err error) {
	err = s.ProjectRepo.Disable(id)

	return
}

func (s *ProjectService) Delete(id uint) (err error) {
	err = s.ProjectRepo.Delete(id)

	return
}

func (s *ProjectService) GetDefault() (project model.Project, err error) {
	project, err = s.ProjectRepo.GetDefault()

	return
}
