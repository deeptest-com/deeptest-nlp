package model

type NluLookup struct {
	BaseModel

	//Code  string          `json:"code"`
	Name  string          `json:"name"`
	Items []NluLookupItem `json:"items" gorm:"-"`

	ProjectId uint `json:"projectId"`
}
type NluLookupItem struct {
	BaseModel

	Content  string `json:"content"`
	LookupId uint   `json:"lookupId"`
}

func (NluLookup) TableName() string {
	return "nlu_lookup"
}
func (NluLookupItem) TableName() string {
	return "nlu_lookup_item"
}
