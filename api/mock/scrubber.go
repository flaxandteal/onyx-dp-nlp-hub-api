package mock

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ONSdigital/dp-nlp-hub/models"
)

func CreateScrubberServer(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Return a mocked scrubber response
		response := models.Scrubber{
			Query: "mocked query",
			Results: models.Results{
				Areas: []models.AreaResp{
					{
						Codes:      map[string]string{"TC1": "TC1"},
						Name:       "TN1",
						Region:     "TR1",
						RegionCode: "TRC",
					},
				},
				Industries: []models.IndustryResp{
					{
						Code: "1111",
						Name: "IT",
					},
				},
			},
			Time: "2023-02-21T15:30:00Z",
		}

		responseBytes, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := w.Write(responseBytes); err != nil {
			t.Error(err)
		}
	}))
}
