package serverService

import (
	"github.com/utlai/utl/internal/server/repo"
)

type PermService struct {
	PermRepo *repo.PermRepo `inject:""`
}

func NewPermService() *PermService {
	return &PermService{}
}
