package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ONSdigital/dp-nlp-hub/api/mock"
	"github.com/ONSdigital/dp-nlp-hub/config"
	"github.com/ONSdigital/dp-nlp-hub/models"
	"github.com/stretchr/testify/assert"
)

func TestHubHandler(t *testing.T) {
	// Create mock servers for the hub to call
	berlin := mock.CreateBerlinServer(t)
	defer berlin.Close()

	scrubber := mock.CreateScrubberServer(t)
	defer scrubber.Close()

	category := mock.CreateCategoryServer(t)
	defer category.Close()

	cfg := config.Config{
		ScrubberBase: scrubber.URL,
		BerlinBase:   berlin.URL,
		CategoryBase: category.URL,
	}

	// Create new test server to handle the http request
	server := httptest.NewServer(http.HandlerFunc(HubHandler(&cfg)))
	defer server.Close()

	// Create a test request to the server
	req, err := http.NewRequest("GET", server.URL+"?q=value", nil)
	assert.NoError(t, err)

	// Send the request and record the response
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()

	// Check that the response has a status code of 200 OK
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Decode the response JSON into a Hub object and check that it contains expected data
	var hub models.Hub
	err = json.NewDecoder(resp.Body).Decode(&hub)
	assert.NoError(t, err)
	t.Log(hub.Scrubber)
	t.Log(hub.Berlin)
	t.Log(hub.Category)
	assert.NotEmpty(t, hub.Scrubber)
	assert.NotEmpty(t, hub.Berlin)
	assert.NotEmpty(t, hub.Category)
}
