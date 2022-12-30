package cache

import (
	"testing"
	"time"

	"github.com/13excite/bvg-info/pkg/store"
	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {
	cases := []struct {
		// name string
		stops                 store.CachedStops
		wantStopNameFromCache string
		wantStopIDFromCache   string
		wantWhenFromCache     time.Time
	}{
		{
			stops: store.CachedStops{
				Name: store.Sudostallee_Kongisheide,
				Departes: []store.StopDepartures{{
					Stop: struct {
						Type     string "json:\"type\""
						ID       string "json:\"id\""
						Name     string "json:\"name\""
						Location struct {
							Type      string  "json:\"type\""
							ID        string  "json:\"id\""
							Latitude  float64 "json:\"latitude\""
							Longitude float64 "json:\"longitude\""
						} "json:\"location\""
						Products struct {
							Suburban bool "json:\"suburban\""
							Subway   bool "json:\"subway\""
							Tram     bool "json:\"tram\""
							Bus      bool "json:\"bus\""
							Ferry    bool "json:\"ferry\""
							Express  bool "json:\"express\""
							Regional bool "json:\"regional\""
						} "json:\"products\""
						StationDHID string "json:\"stationDHID\""
					}{
						Type: "stop",
						ID:   "900000194519",
						Name: "Südostallee/Königsheide",
					},
					When:        time.Date(2022, 12, 29, 15, 14, 0, 0, time.UTC),
					PlannedWhen: time.Date(2022, 12, 29, 15, 14, 0, 0, time.UTC),
				}},
			},
			wantStopIDFromCache:   "900000194519",
			wantStopNameFromCache: "Südostallee/Königsheide",
			wantWhenFromCache:     time.Date(2022, 12, 29, 15, 14, 0, 0, time.UTC),
		},
		// add test for testing "key not found"
	}

	gCache := NewGCache()

	for _, tc := range cases {
		err := gCache.update(tc.stops)
		require.Equal(t, nil, err, "Got error from cache update")

		got, err := gCache.read(tc.stops.Name)

		require.Equal(t, nil, err, "Got error from cache read: Key: ", tc.stops.Name)

		require.Equal(t, tc.wantStopIDFromCache, got.Departes[0].Stop.ID, "Stop.ID is incorrect. Case: ")
	}

}
