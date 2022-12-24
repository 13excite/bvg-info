package bvv

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"testing"
)

var exampleJSON string = `{
	"id": 2355,
	"time": "10:20:30",
	"from": "Strendamm",
	"tram": true
  }
]`

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

func TestIcsClientGetData(t *testing.T) {
	cases := []struct {
		name           string
		statusCode     int
		response       *http.Response
		apiUrl         string
		wantStatusCode int
		wantBody       []byte
		wantURL        string
		wantErr        error
	}{
		{
			name: "200 OK response",
			response: &http.Response{
				StatusCode: http.StatusOK,
				// Send response to be tested
				Body: ioutil.NopCloser(bytes.NewBufferString(exampleJSON)),
				// Must be set to non-nil value or it panics
				Header: make(http.Header),
			},
			statusCode:     http.StatusOK,
			apiUrl:         "http://v5.vbb.transport.rest",
			wantStatusCode: http.StatusOK,
			wantBody:       []byte(exampleJSON),
			wantURL:        "http://v5.vbb.transport.rest/stops/123/departures",
			wantErr:        nil,
		},
	}

	for _, tc := range cases {
		hClient := NewTestClient(func(req *http.Request) *http.Response {
			// check req url
			require.Equal(t, tc.wantURL, req.URL.String(), "URL from request is incorrect")
			return tc.response
		})

		bvvClient := NewClent(tc.apiUrl)
		bvvClient.SetHTTPClient(hClient)

		body, err := bvvClient.GetNearbyDepartes()

		require.Equal(t, tc.wantErr, err, "Test error: "+tc.name)

		require.Equal(t, tc.wantBody, body, "Test body: "+tc.name)
	}
}
