package cache

import (
	"errors"
	"testing"
	"time"

	"github.com/13excite/bvg-info/pkg/store"
	"github.com/stretchr/testify/require"
	//"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	cases := []struct {
		name                  string
		getterCacheKey        string // key name for getting date from cache
		stops                 store.CachedStops
		wantStopNameFromCache string
		wantStopIDFromCache   string
		wantWhenFromCache     time.Time
		wantError             error
	}{
		{
			name:                  "Set/get value from cache",
			getterCacheKey:        store.Sudostallee_Kongisheide,
			wantStopIDFromCache:   "733612",
			wantStopNameFromCache: "Südostallee/Königsheide",
			wantWhenFromCache:     time.Date(2022, 12, 29, 15, 14, 0, 0, time.UTC),
			wantError:             nil,
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
							NationalExpress bool "json:\"nationalExpress\""
							National        bool "json:\"national\""
							RegionalExpress bool "json:\"regionalExpress\""
							Regional        bool "json:\"regional\""
							Suburban        bool "json:\"suburban\""
							Bus             bool "json:\"bus\""
							Ferry           bool "json:\"ferry\""
							Subway          bool "json:\"subway\""
							Tram            bool "json:\"tram\""
							Taxi            bool "json:\"taxi\""
						} "json:\"products\""
						Station struct {
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
								NationalExpress bool "json:\"nationalExpress\""
								National        bool "json:\"national\""
								RegionalExpress bool "json:\"regionalExpress\""
								Regional        bool "json:\"regional\""
								Suburban        bool "json:\"suburban\""
								Bus             bool "json:\"bus\""
								Ferry           bool "json:\"ferry\""
								Subway          bool "json:\"subway\""
								Tram            bool "json:\"tram\""
								Taxi            bool "json:\"taxi\""
							} "json:\"products\""
						} "json:\"station\""
					}{
						Type: "stop",
						ID:   "733612",
						Name: "Südostallee/Königsheide",
					},
					When:        time.Date(2022, 12, 29, 15, 14, 0, 0, time.UTC),
					PlannedWhen: time.Date(2022, 12, 29, 15, 14, 0, 0, time.UTC),
				}},
			},
		},
		{
			name:           "Value not found in cache",
			getterCacheKey: "NotFoundKey",
			wantError:      errors.New("the stop isn't in cache"),
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
							NationalExpress bool "json:\"nationalExpress\""
							National        bool "json:\"national\""
							RegionalExpress bool "json:\"regionalExpress\""
							Regional        bool "json:\"regional\""
							Suburban        bool "json:\"suburban\""
							Bus             bool "json:\"bus\""
							Ferry           bool "json:\"ferry\""
							Subway          bool "json:\"subway\""
							Tram            bool "json:\"tram\""
							Taxi            bool "json:\"taxi\""
						} "json:\"products\""
						Station struct {
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
								NationalExpress bool "json:\"nationalExpress\""
								National        bool "json:\"national\""
								RegionalExpress bool "json:\"regionalExpress\""
								Regional        bool "json:\"regional\""
								Suburban        bool "json:\"suburban\""
								Bus             bool "json:\"bus\""
								Ferry           bool "json:\"ferry\""
								Subway          bool "json:\"subway\""
								Tram            bool "json:\"tram\""
								Taxi            bool "json:\"taxi\""
							} "json:\"products\""
						} "json:\"station\""
					}{
						Type: "stop",
						ID:   "733612",
						Name: "Südostallee/Königsheide",
					},
					When:        time.Date(2022, 12, 29, 15, 14, 0, 0, time.UTC),
					PlannedWhen: time.Date(2022, 12, 29, 15, 14, 0, 0, time.UTC),
				}},
			},
		},
	}
	gCache := NewGCache()

	for _, tc := range cases {
		err := gCache.update(tc.stops)
		require.Equal(t, nil, err, "Got error from cache update. Test case: "+tc.name)

		got, err := gCache.read(tc.getterCacheKey)

		if err != nil {
			require.Equal(t, tc.wantError, err, "Got error from cache read: Key: "+tc.stops.Name+". Case: "+tc.name)

		} else {
			require.Equal(t, tc.wantStopIDFromCache, got.Departes[0].Stop.ID, "Stop.ID is incorrect. Test case: "+tc.name)
			require.Equal(t, tc.wantStopNameFromCache, got.Departes[0].Stop.Name, "Stop.Name field is incorrect. Test case: "+tc.name)
			require.Equal(t, tc.wantWhenFromCache, got.Departes[0].When, "When field is incorrect. Test case: "+tc.name)
		}

	}
}
