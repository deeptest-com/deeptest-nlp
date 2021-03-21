package domain

import (
	_const "github.com/utlai/utl/internal/pkg/const"
	"github.com/utlai/utl/internal/server/model"
)

type ResItem struct {
	Ident    string `json:"ident,omitempty"`
	Computer string `json:"computer,omitempty"`
	Cluster  string `json:"cluster,omitempty"`

	Name     string         `json:"name"`
	Path     string         `json:"path"`
	Type     _const.ResType `json:"type"`
	Key      string         `json:"key"`
	Children []*ResItem     `json:"children"`

	IsTemplate bool `json:"isTemplate"`

	Ip       string `json:"ip,omitempty"`
	Port     int    `json:"port,omitempty"`
	Username string `json:"-"`
	Password string `json:"-"`

	ComputerObj model.Computer `gorm:"-" json:"computer"` // only for computer data persistence
}
