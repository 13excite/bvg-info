package bvv

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

var exampleJSON string = `[
	{
		"tripId": "1|20914|3|86|24122022",
		"stop": {
			"type": "stop",
			"id": "900000194519",
			"name": "Südostallee/Königsheide",
			"location": {
				"type": "location",
				"id": "900194519",
				"latitude": 52.456516,
				"longitude": 13.500755
			},
			"products": {
				"suburban": false,
				"subway": false,
				"tram": false,
				"bus": true,
				"ferry": false,
				"express": false,
				"regional": false
			},
			"stationDHID": "de:11000:900194519"
		},
		"when": "2022-12-24T17:10:00+01:00",
		"plannedWhen": "2022-12-24T17:10:00+01:00",
		"delay": 0,
		"platform": null,
		"plannedPlatform": null,
		"prognosisType": "prognosed",
		"direction": "U Boddinstr.",
		"provenance": null,
		"origin": null,
		"destination": {
			"type": "stop",
			"id": "900000079152",
			"name": "Fontanestr./Flughafenstr.",
			"location": {
				"type": "location",
				"id": "900079152",
				"latitude": 52.480257,
				"longitude": 13.421165
			},
			"products": {
				"suburban": false,
				"subway": false,
				"tram": false,
				"bus": true,
				"ferry": false,
				"express": false,
				"regional": false
			},
			"stationDHID": "de:11000:900079152"
		}
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
		stopId         int
		response       *http.Response
		apiUrl         string
		wantStatusCode int
		wantStopName   string
		wantURL        string
		wantErr        error
	}{
		{
			name:   "200 OK response",
			stopId: 900000194519,
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
			wantStopName:   "Südostallee/Königsheide",
			wantURL:        "http://v5.vbb.transport.rest/stops/900000194519/departures",
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

		departes, err := bvvClient.GetNearbyDepartes(tc.stopId)

		require.Equal(t, tc.wantErr, err, "Test error: "+tc.name)

		require.Equal(t, tc.wantStopName, departes[0].Stop.Name, "Test struct decoding: "+tc.name)
	}
}
