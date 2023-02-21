package params

import (
	"net/url"
)

type BerlinParams struct {
	Query string `url:"q,omitempty"`
	State string `url:"state,omitempty"`
}

func GetBerlinParams(query url.Values) *BerlinParams {
	result := BerlinParams{
		Query: "",
		State: "",
	}

	if len(query["q"]) >= 1 {
		result.Query = query["q"][0]
	}

	if len(query["state"]) >= 1 {
		result.State = query["state"][0]
	}

	return &result
}
