package model

type NluRegex struct {
	BaseModel

	Code  string         `json:"code"`
	Name  string         `json:"name"`
	Items []NluRegexItem `json:"items" gorm:"-"`

	ProjectId uint `json:"projectId"`
}
type NluRegexItem struct {
	BaseModel

	Content string `json:"content"`
	RegexId uint   `json:"regexId"`
}

func (NluRegex) TableName() string {
	return "nlu_regex"
}
func (NluRegexItem) TableName() string {
	return "nlu_regex_item"
}
