package api

import (
	"encoding/json"
	"net/http"

	"github.com/ONSdigital/log.go/v2/log"

	"github.com/ONSdigital/dp-nlp-hub/config"
	"github.com/ONSdigital/dp-nlp-hub/models"
)

func HubHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ctx := r.Context()

		var result models.Hub

		// Gets the scrubber response
		err := MakeRequest(ctx, cfg.ScrubberBase, models.GetScrubberParams(r.URL.Query()), &result.Scrubber)
		if err != nil {
			log.Warn(ctx, "There was an error making request to Scrubber: "+err.Error())
		}
		// Gets the berlin response using a filter from url params and a query from scrubber
		err = MakeRequest(ctx, cfg.BerlinBase, models.GetBerlinParams(r.URL.Query()), &result.Berlin)
		if err != nil {
			log.Warn(ctx, "There was an error making request to Berlin: "+err.Error())
		}
		// Gets the category response using berlin normalized query
		err = MakeRequest(ctx, cfg.CategoryBase, models.GetCategoryParams(result.Berlin.Query.Normalized), &result.Category)
		if err != nil {
			log.Warn(ctx, "There was an error making request to Category: "+err.Error())
		}
		// This is here for testing purposes
		allResponses, err := json.Marshal(result)
		if err != nil {
			w.Write([]byte("issue with marshaling"))
			return
		}

		w.Write([]byte(allResponses))
	}
}
