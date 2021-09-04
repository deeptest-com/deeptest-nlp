package model

import serverConst "github.com/utlai/utl/internal/server/utils/const"

type NluSlot struct {
	BaseModel

	Seq   int                     `json:"seq"`
	Type  serverConst.NluSlotType `json:"type"`
	Value string                  `json:"value"`
	Text  string                  `json:"text"`

	SentRefer uint `json:"sentRefer"`
}

func (NluSlot) TableName() string {
	return "nlu_slot"
}
