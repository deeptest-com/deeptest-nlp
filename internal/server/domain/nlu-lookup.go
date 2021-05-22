package domain

type NluLookup struct {
	Version   string      `json:"version"`
	LookupDef []LookupDef `json:"nlu"`
}

type LookupDef struct {
	Lookup   string `json:"lookup"`
	Examples string `json:"examples"`
}
