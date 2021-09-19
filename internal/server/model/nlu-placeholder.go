package model

type NluPlaceholder struct {
	BaseModel

	Code      string `json:"code"`
	Name      string `json:"name"`
	Ordr      int    `json:"ordr"`
	ProjectId uint   `json:"projectId"`
}

func (NluPlaceholder) TableName() string {
	return "nlu_placeholder"
}
