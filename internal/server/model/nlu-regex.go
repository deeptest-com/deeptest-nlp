package model

type NluRegex struct {
	BaseModel

	Code  string         `json:"code"`
	Name  string         `json:"name"`
	Ordr  int            `json:"ordr"`
	Items []NluRegexItem `json:"items" gorm:"-"`

	ProjectId uint `json:"projectId"`
}
type NluRegexItem struct {
	BaseModel
	Code    string `json:"code"`
	Name    string `json:"name"`
	Ordr    int    `json:"ordr"`
	RegexId uint   `json:"regexId"`
}

func (NluRegex) TableName() string {
	return "nlu_regex"
}
func (NluRegexItem) TableName() string {
	return "nlu_regex_item"
}
