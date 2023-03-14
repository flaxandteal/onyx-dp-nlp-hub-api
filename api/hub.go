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

		var result models.Hub

		scrubberCh := make(chan error)
		berlinCh := make(chan error)
		categoryCh := make(chan error)

		var wg sync.WaitGroup
		wg.Add(3)

		go func() {
			defer wg.Done()
			err := MakeRequest(r.Context(), cfg.ScrubberBase, models.GetScrubberParams(r.URL.Query()), &result.Scrubber)
			scrubberCh <- err
		}()

		go func() {
			defer wg.Done()
			err := MakeRequest(r.Context(), cfg.BerlinBase, models.GetBerlinParams(r.URL.Query()), &result.Berlin)
			berlinCh <- err
		}()

		go func() {
			defer wg.Done()
			err := MakeRequest(r.Context(), cfg.CategoryBase, models.GetCategoryParams(r.URL.Query()), &result.Category)
			categoryCh <- err
		}()

		wg.Wait()
		close(scrubberCh)
		close(berlinCh)
		close(categoryCh)

		// Check for errors in each channel
		for err := range scrubberCh {
			if err != nil {
				log.Warn(ctx, fmt.Sprintf("Scrubber error: %s", err.Error()))
			}
		}

		for err := range berlinCh {
			if err != nil {
				log.Warn(ctx, fmt.Sprintf("Berlin error: %s", err.Error()))
			}
		}

		for err := range categoryCh {
			if err != nil {
				log.Warn(ctx, fmt.Sprintf("Category error: %s", err.Error()))
			}
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
