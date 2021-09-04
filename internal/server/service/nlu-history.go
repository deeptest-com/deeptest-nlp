package serverService

import (
	"github.com/utlai/utl/internal/server/model"
	"github.com/utlai/utl/internal/server/repo"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
)

type NluHistoryService struct {
	NluHistoryRepo *repo.NluHistoryRepo `inject:""`

	ProjectRepo *repo.ProjectRepo `inject:""`
	UserRepo    *repo.UserRepo    `inject:""`
}

func NewNluHistoryService() *NluHistoryService {
	return &NluHistoryService{}
}

func (s *NluHistoryService) Add(userId uint, projectId uint, action serverConst.ServiceStatus) (err error) {
	user, _ := s.UserRepo.Get(userId)
	project := s.ProjectRepo.GetDetail(projectId)
	po := model.NluHistory{
		UserId:      uint(userId),
		ProjectId:   projectId,
		UserName:    user.Name,
		ProjectName: project.Name,
		Action:      action,
	}

	err = s.NluHistoryRepo.Save(&po)

	return
}
