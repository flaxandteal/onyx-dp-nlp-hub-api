package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/ONSdigital/dp-nlp-hub/clients"
	"github.com/ONSdigital/dp-nlp-hub/config"
	"github.com/ONSdigital/log.go/v2/log"
	"github.com/gorilla/mux"
)

// API provides a struct to wrap the api around
type API struct {
	Router *mux.Router
}

// Setup function sets up the api and returns an api
func Setup(ctx context.Context, r *mux.Router, cfg *config.Config) *API {
	api := &API{
		Router: r,
	}

	r.HandleFunc("/nlp-hub/search", HubHandler(cfg)).Methods("GET").Name("HubHandler")
	return api
}

func MakeRequest(ctx context.Context, url string, params interface{}, resp interface{}) error {
	cl := clients.New(url, params)

	log.Info(ctx, fmt.Sprintf("Making a request to: %s with query: %s", url, params))

	res, err := cl.DoRequest()
	if err != nil {
		return fmt.Errorf("building request for %s has failed with err: %s", url, err.Error())
	}
	defer res.Body.Close()

	log.Info(ctx, "request successful")

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("parsing response body for %s has failed with err: %s", url, err.Error())
	}

	if err := json.Unmarshal(b, &resp); err != nil {
		return fmt.Errorf("unmarshaling resp to the given interface for %s has failed with err: %s", url, err.Error())
	}

	return nil
}
