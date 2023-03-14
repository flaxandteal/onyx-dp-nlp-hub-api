package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ONSdigital/dp-nlp-hub/config"
	"github.com/ONSdigital/dp-nlp-hub/models"
)

func HubHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var result models.Hub

		// Gets the scrubber response
		MakeRequest(r.Context(), cfg.ScrubberBase, models.GetScrubberParams(r.URL.Query()), &result.Scrubber)

		// Gets the berlin response using a filter from url params and a query from scrubber
		MakeRequest(r.Context(), cfg.BerlinBase, models.GetBerlinParams(r.URL.Query()), &result.Berlin)

		// Gets the category response using berlin normalized query
		MakeRequest(r.Context(), cfg.CategoryBase, models.GetCategoryParams(result.Berlin.Query.Normalized), &result.Category)

		// This is here for testing purposes
		if err := json.NewEncoder(w).Encode(result); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("An unexpected error occurred while processing your request: " + err.Error()))
		}
	}
}
