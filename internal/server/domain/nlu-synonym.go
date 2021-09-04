package serverDomain

type NluSynonym struct {
	Version    string          `yaml:"version"`
	SynonymDef []NluSynonymDef `yaml:"nlu,flow"`
}

type NluSynonymDef struct {
	Synonym  string `yaml:"synonym"`
	Examples string `yaml:"examples"`
}
