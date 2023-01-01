package cache

import (
	"errors"
	"testing"
	"time"

	"github.com/13excite/bvg-info/pkg/store"
	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {
	cases := []struct {
		name                  string
		getterCacheKey        string
		stops                 store.CachedStops
		wantStopNameFromCache string
		wantStopIDFromCache   string
		wantWhenFromCache     time.Time
		wantError             error
	}{
		{
			name:           "Set/get value from cache",
			getterCacheKey: store.Sudostallee_Kongisheide,
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
			wantError:             nil,
		},
		// add test for testing "key not found"
		{
			name:           "Value not found in cache",
			getterCacheKey: "NotFoundKey",
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
			// wantStopIDFromCache:   "900000194519",
			// wantStopNameFromCache: "NotFoundValue",
			// wantWhenFromCache:     time.Date(2022, 12, 29, 15, 14, 0, 0, time.UTC),
			wantError: errors.New("the stop isn't in cache"),
		},
	}

	gCache := NewGCache()

	for _, tc := range cases {
		err := gCache.update(tc.stops)
		require.Equal(t, nil, err, "Got error from cache update. Test case: "+tc.name)

		got, err := gCache.read(tc.getterCacheKey)

		require.Equal(t, tc.wantError, err, "Got error from cache read: Key: "+tc.stops.Name+"Case: "+tc.name)

		// TODO: skip or add condition when testing error
		require.Equal(t, tc.wantStopIDFromCache, got.Departes[0].Stop.ID, "Stop.ID is incorrect. Test case: "+tc.name)
		require.Equal(t, tc.wantStopNameFromCache, got.Departes[0].Stop.Name, "Stop.Name field is incorrect. Test case: "+tc.name)
		require.Equal(t, tc.wantWhenFromCache, got.Departes[0].When, "When field is incorrect. Test case: "+tc.name)

	}
}
