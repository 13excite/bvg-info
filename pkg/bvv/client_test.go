package bvv

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

var exampleJSON string = `
{
	"departures": [
	  {
		"tripId": "1|1124378|2|80|13062023",
		"stop": {
		  "type": "stop",
		  "id": "733612",
		  "name": "Südostallee/Königsheide, Berlin",
		  "location": {
			"type": "location",
			"id": "733612",
			"latitude": 52.456579,
			"longitude": 13.500638
		  },
		  "products": {
			"nationalExpress": false,
			"national": false,
			"regionalExpress": false,
			"regional": false,
			"suburban": false,
			"bus": true,
			"ferry": false,
			"subway": false,
			"tram": false,
			"taxi": false
		  }
		},
		"when": "2023-06-13T10:39:00+02:00",
		"plannedWhen": "2023-06-13T10:39:00+02:00",
		"delay": 0,
		"platform": null,
		"plannedPlatform": null,
		"prognosisType": "prognosed",
		"direction": "S Schöneweide",
		"provenance": null,
		"line": {
		  "type": "line",
		  "id": "5-vbbbvb-166",
		  "fahrtNr": "22483",
		  "name": "Bus 166",
		  "public": true,
		  "adminCode": "vbbBVB",
		  "productName": "Bus",
		  "mode": "bus",
		  "product": "bus",
		  "operator": {
			"type": "operator",
			"id": "nahreisezug",
			"name": "Nahreisezug"
		  }
		},
		"remarks": [],
		"origin": null,
		"destination": {
		  "type": "stop",
		  "id": "733587",
		  "name": "Schöneweide (S)/Sterndamm, Berlin",
		  "location": {
			"type": "location",
			"id": "733587",
			"latitude": 52.453397,
			"longitude": 13.509618
		  },
		  "products": {
			"nationalExpress": false,
			"national": false,
			"regionalExpress": false,
			"regional": true,
			"suburban": true,
			"bus": true,
			"ferry": false,
			"subway": false,
			"tram": true,
			"taxi": false
		  },
		  "station": {
			"type": "station",
			"id": "8010041",
			"name": "Berlin-Schöneweide",
			"location": {
			  "type": "location",
			  "id": "8010041",
			  "latitude": 52.455204,
			  "longitude": 13.508773
			},
			"products": {
			  "nationalExpress": false,
			  "national": false,
			  "regionalExpress": false,
			  "regional": true,
			  "suburban": true,
			  "bus": true,
			  "ferry": false,
			  "subway": false,
			  "tram": true,
			  "taxi": false
			}
		  }
		},
		"currentTripPosition": {
		  "type": "location",
		  "latitude": 52.459213,
		  "longitude": 13.496458
		}
	  }
	]
}
`

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
			stopId: 733612,
			response: &http.Response{
				StatusCode: http.StatusOK,
				// Send response to be tested
				Body: ioutil.NopCloser(bytes.NewBufferString(exampleJSON)),
				// Must be set to non-nil value or it panics
				Header: make(http.Header),
			},
			statusCode:     http.StatusOK,
			apiUrl:         "https://v6.db.transport.rest",
			wantStatusCode: http.StatusOK,
			wantStopName:   "Südostallee/Königsheide, Berlin",
			wantURL:        "https://v6.db.transport.rest/stops/733612/departures",
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

		require.Equal(t, tc.wantStopName, departes.Departures[0].Stop.Name, "Test struct decoding: "+tc.name)
	}
}
