package model

import (
	serverConst "github.com/utlai/utl/internal/server/utils/const"
	"time"
)

type NluHistory struct {
	BaseModel

	Html string `json:"html"`
	Text string `json:"text"`

	ProjectId   uint   `sql:"index" json:"projectId"`
	UserId      uint   `json:"userId"`
	ProjectName string `json:"projectName"`
	UserName    string `json:"userName"`

	Action      serverConst.NluAction `json:"action"`
	CompletedAt *time.Time            `json:"completedAt"`
}

func (NluHistory) TableName() string {
	return "nlu_history"
}
