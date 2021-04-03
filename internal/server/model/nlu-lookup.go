package model

type NluLookup struct {
	Name  string          `yaml:"name"`
	Items []NluLookupItem `yaml:"items"`
}
type NluLookupItem struct {
	Content string `yaml:"content"`

	LookupId uint `yaml:"lookupId"`
}

func (NluLookup) TableName() string {
	return "nlu_lookup"
}
func (NluLookupItem) TableName() string {
	return "nlu_lookup_item"
}
