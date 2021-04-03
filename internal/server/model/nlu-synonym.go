package model

type NluSynonym struct {
	BaseModel

	Name  string           `yaml:"name"`
	Items []NluSynonymItem `yaml:"items"`
}
type NluSynonymItem struct {
	Content string `yaml:"content"`

	SynonymId uint `yaml:"synonymId"`
}

func (NluSynonym) TableName() string {
	return "nlu_synonym"
}
func (NluSynonymItem) TableName() string {
	return "nlu_synonym_item"
}
