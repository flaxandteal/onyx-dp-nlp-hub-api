package payloads

type BerlinSearchSchemaJson struct {
	Query   SearchTermJson `json:"query"`
	Results []SearchResult `json:"results"`
	Time    string         `json:"time"`
}

type SearchTermJson struct {
	Codes           []string    `json:"codes"`
	ExactMatches    []string    `json:"exact_matches"`
	Normalized      string      `json:"normalized"`
	NotExactMatches []string    `json:"not_exact_matches"`
	Raw             string      `json:"raw"`
	StateFilter     interface{} `json:"state_filter,omitempty"`
	StopWords       []string    `json:"stop_words"`
}

type SearchResult struct {
	Loc   LocJson `json:"loc"`
	Score int     `json:"score"`
}

type LocJson struct {
	Codes    []string    `json:"codes"`
	Encoding string      `json:"encoding"`
	Id       string      `json:"id"`
	Key      string      `json:"key"`
	Names    []string    `json:"names"`
	State    []string    `json:"state"`
	Subdiv   interface{} `json:"subdiv,omitempty"`
}
