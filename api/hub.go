package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ONSdigital/dp-nlp-hub/config"
	"github.com/ONSdigital/dp-nlp-hub/models"
	"github.com/ONSdigital/log.go/v2/log"
)

func HubHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		ctx := r.Context()

		var result models.Hub

		// Gets the scrubber response
		err := MakeRequest(r.Context(), cfg.ScrubberBase, models.GetScrubberParams(r.URL.Query()), &result.Scrubber)
		if err != nil {
			log.Warn(ctx, err.Error())
		}

		// Gets the berlin response using a filter from url params and a query from scrubber
		err = MakeRequest(r.Context(), cfg.BerlinBase, models.GetBerlinParams(r.URL.Query()), &result.Berlin)
		if err != nil {
			log.Warn(ctx, err.Error())
		}

		// Gets the category response using berlin normalized query
		err = MakeRequest(r.Context(), cfg.CategoryBase, models.GetCategoryParams(r.URL.Query()), &result.Category)
		if err != nil {
			log.Warn(ctx, err.Error())
		}

		if err := json.NewEncoder(w).Encode(result); err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			errObj := ErrorResp{
				errors: []Errors{
					{
						error_code: "", // to be added once Nathan finished the error-codes lib
						message:    "An unexpected error occurred while processing your request",
					},
				},
				trace_id: getRequestId(ctx),
			}

			json.NewEncoder(w).Encode(errObj)
		}
	}
}

func getRequestId(ctx context.Context) string {
	requestID := ctx.Value("request-id")

	correlationID, ok := requestID.(string)
	if !ok {
		return ""
	}

	return correlationID
}
