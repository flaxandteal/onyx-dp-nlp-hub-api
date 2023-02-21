package mock

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/ONSdigital/dp-nlp-hub/payloads"
)

func CreateCategoryServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Return a mocked category response
		categories := payloads.Category{
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
		w.Write(response)
	}))
}
