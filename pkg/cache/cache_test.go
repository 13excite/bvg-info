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
				Departes: []store.CachedStop{
					{
						ID:          "733612",
						Name:        "Südostallee/Königsheide",
						Time:        time.Date(2022, 12, 29, 15, 14, 0, 0, time.UTC),
						PlannedTime: time.Date(2022, 12, 29, 15, 14, 0, 0, time.UTC),
					},
				},
			},
		},
		{
			name:           "Value not found in cache",
			getterCacheKey: "NotFoundKey",
			wantError:      errors.New("the stop isn't in cache"),
			stops: store.CachedStops{
				Name: store.Sudostallee_Kongisheide,
				Departes: []store.CachedStop{
					{
						ID:          "733612",
						Name:        "Südostallee/Königsheide",
						Time:        time.Date(2022, 12, 29, 15, 14, 0, 0, time.UTC),
						PlannedTime: time.Date(2022, 12, 29, 15, 14, 0, 0, time.UTC),
					},
				},
			},
		},
	}
	gCache := NewGCache()

	for _, tc := range cases {
		err := gCache.Update(tc.stops.Name, tc.stops.Departes)
		require.Equal(t, nil, err, "Got error from cache update. Test case: "+tc.name)

		got, err := gCache.Read(tc.getterCacheKey)

		if err != nil {
			require.Equal(t, tc.wantError, err, "Got error from cache read: Key: "+tc.stops.Name+". Case: "+tc.name)

		} else {
			require.Equal(t, tc.wantStopIDFromCache, got[0].ID, "Stop.ID is incorrect. Test case: "+tc.name)
			require.Equal(t, tc.wantStopNameFromCache, got[0].Name, "Stop.Name field is incorrect. Test case: "+tc.name)
			require.Equal(t, tc.wantWhenFromCache, got[0].Time, "When field is incorrect. Test case: "+tc.name)
		}

	}
}
