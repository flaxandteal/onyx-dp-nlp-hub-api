package mock

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ONSdigital/dp-nlp-hub/models"
)

func CreateCategoryServer(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Return a mocked category response
		categories := models.Category{
			{
				Code:  []string{"mocked category"},
				Score: 1.2,
			},
		}
		response, err := json.Marshal(categories)
		if err != nil {
			http.Error(w, "Error marshaling category response", http.StatusInternalServerError)
			return
		}

		if _, err := w.Write(response); err != nil {
			t.Error(err)
		}
	}))
}
