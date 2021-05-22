package domain

type NluSynonym struct {
	Version    string       `json:"version"`
	SynonymDef []SynonymDef `json:"nlu"`
}

type SynonymDef struct {
	Synonym  string `json:"synonym"`
	Examples string `json:"examples"`
}
