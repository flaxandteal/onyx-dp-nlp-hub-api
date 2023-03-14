package mock

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ONSdigital/dp-nlp-hub/models"
)

func CreateBerlinServer(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Return a mocked berlin response
		response := models.Berlin{
			Query: models.SearchTermJson{
				Codes:           []string{"test1", "test2"},
				ExactMatches:    []string{"Berlin"},
				Normalized:      "Berlin",
				NotExactMatches: []string{"berlin", "berlin"},
				Raw:             "Berlin",
				StopWords:       []string{"in", "a"},
			},
			Results: []models.SearchResult{
				{
					Loc: models.LocJson{
						Codes:    []string{"test1", "test2"},
						Encoding: "UTF-8",
						Id:       "9",
						Key:      "9",
						Names:    []string{"a"},
						State:    []string{"a"},
					},
					Score: 100,
				},
				{
					Loc: models.LocJson{
						Codes:    []string{"BRE", "DEU"},
						Encoding: "UTF-8",
						Id:       "10",
						Key:      "10",
						Names:    []string{"b"},
						State:    []string{"b"},
					},
					Score: 50,
				},
			},
			Time: "10",
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := w.Write(jsonResponse); err != nil {
			t.Error(err)
		}
	}))
}
