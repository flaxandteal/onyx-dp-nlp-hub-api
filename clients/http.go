package clients

import (
	"fmt"
	"io"
	"net/http"

	"github.com/dghubble/sling"
)

type HttpClient struct {
	Sl *sling.Sling
}

func New(url string, params interface{}) *HttpClient {
	sl := sling.New().Base(url).QueryStruct(params)
	return &HttpClient{
		Sl: sl,
	}
}

func (cl *HttpClient) DoRequest() (*http.Response, error) {
	req, err := cl.Sl.Request()
	if err != nil {
		return nil, fmt.Errorf("building request %s has failed", req.URL)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to make do request: %v", err.Error())
	}

	// use http.statusok and between 200 and 300
	if res.StatusCode != 200 {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("invalid request: status code: %d \n Response body: %s\n Client sling: %s", res.StatusCode, body, cl.Sl)
	}

	return res, nil
}
