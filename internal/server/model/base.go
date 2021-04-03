package model

import "time"

type BaseModel struct {
	//gorm.Model

	ID        uint       `gorm:"primary_key" sql:"type:INT(10) UNSIGNED NOT NULL" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
}

var (
	Models = []interface{}{
		&Project{},
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
