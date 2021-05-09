package model

type NluSynonym struct {
	BaseModel

	Name  string           `json:"name"`
	Items []NluSynonymItem `json:"items" gorm:"-"`
}
type NluSynonymItem struct {
	BaseModel

	Content   string `json:"content"`
	SynonymId uint   `json:"synonymId"`
}

func (NluSynonym) TableName() string {
	return "nlu_synonym"
}
func (NluSynonymItem) TableName() string {
	return "nlu_synonym_item"
}
