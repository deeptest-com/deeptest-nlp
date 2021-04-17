package model

type NluLookup struct {
	BaseModel

	Name  string          `json:"name"`
	Items []NluLookupItem `json:"items" gorm:"-"`
}
type NluLookupItem struct {
	Content  string `json:"content"`
	LookupId uint   `json:"lookupId"`
}

func (NluLookup) TableName() string {
	return "nlu_lookup"
}
func (NluLookupItem) TableName() string {
	return "nlu_lookup_item"
}
