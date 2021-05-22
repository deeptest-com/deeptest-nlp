package domain

type NluSynonym struct {
	Version    string          `json:"version"`
	SynonymDef []NluSynonymDef `json:"nlu"`
}

type NluSynonymDef struct {
	Synonym  string `json:"synonym"`
	Examples string `json:"examples"`
}
