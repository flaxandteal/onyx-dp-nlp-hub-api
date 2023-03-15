package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/ONSdigital/dp-nlp-hub/config"
	"github.com/ONSdigital/dp-nlp-hub/models"
	"github.com/ONSdigital/log.go/v2/log"
)

func HubHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		ctx := r.Context()

		var scrubber *models.Scrubber
		var berlin *models.Berlin
		var category *models.Category

		var wg sync.WaitGroup
		wg.Add(3)

		go func() {
			defer wg.Done()
			err := MakeRequest(r.Context(), cfg.ScrubberBase, models.GetScrubberParams(r.URL.Query()), scrubber)
			log.Warn(ctx, fmt.Sprintf("Scrubber error: %s", err.Error()))
		}()

		go func() {
			defer wg.Done()
			err := MakeRequest(r.Context(), cfg.BerlinBase, models.GetBerlinParams(r.URL.Query()), berlin)
			log.Warn(ctx, fmt.Sprintf("Berlin error: %s", err.Error()))
		}()

		go func() {
			defer wg.Done()
			err := MakeRequest(r.Context(), cfg.CategoryBase, models.GetCategoryParams(r.URL.Query()), category)
			log.Warn(ctx, fmt.Sprintf("Category error: %s", err.Error()))
		}()

		wg.Wait()

		result := models.Hub{
			Scrubber: *scrubber,
			Berlin:   *berlin,
			Category: *category,
		}

		if err := json.NewEncoder(w).Encode(result); err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			log.Error(ctx, "error encoding result", err)

			errObj := ErrorResp{
				Errors: []Errors{
					{
						Error_code: "", // to be added once Nathan finished the error-codes lib
						Message:    "An unexpected error occurred while processing your request",
					},
				},
				Trace_id: getRequestId(ctx),
			}

			if err := json.NewEncoder(w).Encode(errObj); err != nil {
				log.Error(ctx, "error encoding errObj", err)
			}
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
