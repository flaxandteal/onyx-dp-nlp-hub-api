package models

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBerlinParams(t *testing.T) {
	query := url.Values{}
	query.Set("q", "Berlin")
	query.Set("state", "DE")

	berlinParams := GetBerlinParams(query)

	assert.Equal(t, "Berlin", berlinParams.Query)
	assert.Equal(t, "DE", berlinParams.State)
}

func TestGetScrubberParams(t *testing.T) {
	query := url.Values{}
	query.Set("q", "Berlin")

	berlinParams := GetScrubberParams(query)

	assert.Equal(t, "Berlin", berlinParams.Query)
}
