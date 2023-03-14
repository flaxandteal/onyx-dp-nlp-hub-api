package clients

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoRequest(t *testing.T) {
	// Create a test server to simulate the API
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(`{"message": "OK"}`)); err != nil {
			t.Error(err)
		}
	}))
	defer server.Close()

	// Create an instance of the HttpClient struct
	client := New(server.URL, nil)

	// Send a request to the test server using the HttpClient struct
	resp, err := client.DoRequest()
	assert.NoError(t, err)

	// Check that the HTTP response has the expected status code
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Check that the response body is as expected
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, []byte(`{"message": "OK"}`), body)
}
