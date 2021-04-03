package model

import (
	"time"
)

type BaseModel struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
}

var (
	Models = []interface{}{
		&NluIntent{},
		&NluSent{},
		&NluSlot{},

		&NluSynonym{},
		&NluSynonymItem{},

		&NluLookup{},
		&NluLookupItem{},

		&User{},
		&Role{},
		&Permission{},
	}
)
