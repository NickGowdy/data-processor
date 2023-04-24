package client

import (
	"io"
	"net/http"
)

// Client used to communicate to external system
type SomeHttpClient struct {
	Client Client
}

const _ = "/some_url" // endpoint for saving data

// Send data via POST (HTTP) request
func (h *SomeHttpClient) Post(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return resp, nil
}
