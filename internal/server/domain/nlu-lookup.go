package domain

type NluLookup struct {
	Version   string         `json:"version"`
	LookupDef []NluLookupDef `json:"nlu"`
}

type NluLookupDef struct {
	Lookup   string `json:"lookup"`
	Examples string `json:"examples"`
}
