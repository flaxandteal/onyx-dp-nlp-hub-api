package payloads

type ScrubberSearchSchemaJson struct {
	Query   string  `json:"query"`
	Results Results `json:"results"`
	Time    string  `json:"time"`
}

type Results struct {
	Areas      []AreaResp     `json:"areas"`
	Industries []IndustryResp `json:"industries"`
}

type AreaResp struct {
	Codes      map[string]string `json:"codes"`
	Name       string            `json:"name"`
	Region     string            `json:"region"`
	RegionCode string            `json:"region_code"`
}

type IndustryResp struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
