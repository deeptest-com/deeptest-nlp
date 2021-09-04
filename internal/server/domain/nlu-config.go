package serverDomain

type NluConfig struct {
	Language string     `json:"language"`
	Pipeline []Pipeline `json:"pipeline"`
	Policies []Policy   `json:"policies"`
}

type Pipeline struct {
	Name              string  `json:"name"`
	Model             *string `json:"model,omitempty"`
	Pooling           *string `json:"pooling,omitempty"`
	CaseSensitive     *bool   `json:"case_sensitive,omitempty"`
	UseLookupTables   *bool   `json:"use_lookup_tables,omitempty"`
	UseRegexes        *bool   `json:"use_regexes,omitempty"`
	UseWordBoundaries *bool   `json:"use_word_boundaries,omitempty"`
	Epochs            *int64  `json:"epochs,omitempty"`
	Analyzer          *string `json:"analyzer,omitempty"`
	MinNgram          *int64  `json:"min_ngram,omitempty"`
	MaxNgram          *int64  `json:"max_ngram,omitempty"`
}

type Policy struct {
	Name       string `json:"name"`
	MaxHistory *int64 `json:"max_history,omitempty"`
	Epochs     *int64 `json:"epochs,omitempty"`
}
