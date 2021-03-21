package model

import (
	_const "github.com/utlai/utl/internal/pkg/const"
	"time"
)

type Project struct {
	BaseModel

	// info
	Name     string `json:"name,omitempty"`
	UserName string `json:"userName,omitempty"`

	// job
	BuildType _const.BuildType `json:"buildType,omitempty"`
	Priority  int              `json:"priority,omitempty"`
	GroupId   uint             `json:"groupId,omitempty"`

	// status
	Progress _const.BuildProgress `json:"progress,omitempty"`
	Status   _const.BuildStatus   `json:"status,omitempty"`

	StartTime   time.Time `json:"startTime,omitempty"`
	PendingTime time.Time `json:"pendingTime,omitempty"`
	ResultTime  time.Time `json:"resultTime,omitempty"`
}

func NewProject() Project {
	project := Project{
		Progress: _const.ProgressCreated,
		Status:   _const.StatusCreated,
	}
	return project
}

func (Project) TableName() string {
	return "biz_project"
}
