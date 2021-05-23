package serverConst

type NluSlotType string

const (
	Synonym NluSlotType = "synonym"
	Lookup  NluSlotType = "lookup"
	Regex   NluSlotType = "regex"
	text    NluSlotType = "text"
)
