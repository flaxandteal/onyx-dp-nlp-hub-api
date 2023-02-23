package params

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetScrubberParams(t *testing.T) {
	query := url.Values{}
	query.Set("q", "Berlin")

	berlinParams := GetScrubberParams(query)

	assert.Equal(t, "Berlin", berlinParams.Query)
}
