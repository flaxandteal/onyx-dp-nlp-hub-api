package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ONSdigital/dp-nlp-hub/config"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	// Create a mock config
	cfg := &config.Config{}

	// Create a mock router
	r := mux.NewRouter()

	// Setup the API
	api := Setup(context.Background(), r, cfg)

	// Assert that the Router was set correctly
	assert.Equal(t, r, api.Router)

	// Assert that the "/search" route was added
	route := r.Get("HubHandler")
	assert.NotNil(t, route, "Expected HubHandler to be added")
}

func TestMakeRequest(t *testing.T) {
	// Create a test server to simulate the API
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "OK"}`))
	}))
	defer server.Close()

	params := struct {
		query string `url:"q,omitempty"`
	}{
		query: "querytynabratovchedagosho",
	}

	// Define the expected response
	resp := struct {
		Message string `json:"message"`
	}{}

	// Make a request using the MakeRequest function
	MakeRequest(context.Background(), server.URL, params, &resp)

	// Check the response is as expected
	assert.Equal(t, "OK", resp.Message)
}
