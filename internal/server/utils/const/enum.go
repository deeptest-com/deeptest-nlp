package serverConst

type NluSlotType string

const (
	Synonym NluSlotType = "synonym"
	Lookup  NluSlotType = "lookup"
	Regex   NluSlotType = "regex"
	Text    NluSlotType = "text"
)
