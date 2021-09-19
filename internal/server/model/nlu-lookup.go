package model

type NluLookup struct {
	BaseModel

	Code  string          `json:"code"`
	Name  string          `json:"name"`
	Ordr  int             `json:"ordr"`
	Items []NluLookupItem `json:"items" gorm:"-"`

	ProjectId uint `json:"projectId"`
}
type NluLookupItem struct {
	BaseModel

	Code     string `json:"code"`
	Name     string `json:"name"`
	Ordr     int    `json:"ordr"`
	LookupId uint   `json:"lookupId"`
}

func (NluLookup) TableName() string {
	return "nlu_lookup"
}
func (NluLookupItem) TableName() string {
	return "nlu_lookup_item"
}
