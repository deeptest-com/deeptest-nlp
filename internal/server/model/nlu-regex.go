package model

type NluRegex struct {
	BaseModel

	Name  string         `json:"name"`
	Items []NluRegexItem `json:"items" gorm:"-"`
}
type NluRegexItem struct {
	Content string `json:"content"`
	RegexId uint   `json:"lookupId"`
}

func (NluRegex) TableName() string {
	return "nlu_lookup"
}
func (NluRegexItem) TableName() string {
	return "nlu_lookup_item"
}
