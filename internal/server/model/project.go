package model

import serverConst "github.com/utlai/utl/internal/server/utils/const"

type Project struct {
	BaseModel

	Name          string                `json:"name"`
	Desc          string                `json:"desc"`
	Path          string                `json:"path"`
	IsDefault     bool                  `json:"isDefault"`
	ServicePort   int                   `json:"servicePort"`
	ServiceStatus serverConst.NluAction `json:"serviceStatus"`

	UserId    uint         `json:"userId"`
	CreatedBy string       `json:"createdBy,omitempty" gorm:"-"`
	Histories []NluHistory `json:"histories,omitempty" gorm:"-"`
}

func (Project) TableName() string {
	return "biz_project"
}
