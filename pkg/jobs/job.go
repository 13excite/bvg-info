package jobs

import (
	"context"
	"strconv"
	"time"

	"github.com/13excite/bvg-info/pkg/cache"
	"github.com/13excite/bvg-info/pkg/store"

	"go.uber.org/zap"
)

type httpClient interface {
	GetNearbyDepartes(int) (*store.Departures, error)
}

type Job struct {
	runInterval int
	logger      *zap.SugaredLogger
	hClient     httpClient
	// departuresList []string // TODO: should we pass it???
	gCache cache.Cache
}

func New(runInterval int, httphttpClient httpClient, cache cache.Cache) *Job {
	return &Job{
		runInterval: runInterval,
		logger:      zap.S().With("package", "job"),
		hClient:     httphttpClient,
		gCache:      cache,
	}
}

func (j *Job) httpClientRunner() {

	for stopKey, stop := range store.NearbyDepartures() {
		stopId, _ := strconv.Atoi(stop.ID)

		stops, err := j.hClient.GetNearbyDepartes(stopId)
		if err != nil {
			j.logger.Error("httpJob failed. Got error from client", "error", err)
			continue
		}
		resultByStop := []store.CachedStop{}
		for _, departe := range stops.Departures {
			resultByStop = append(resultByStop, store.CachedStop{
				ID:          departe.Stop.ID,
				Name:        departe.Stop.Name,
				Time:        departe.When,
				PlannedTime: departe.PlannedWhen,
				Direction:   departe.Direction,
				LineName:    departe.Line.Name,
				ProductName: departe.Line.ProductName,
				Remarks:     departe.Remarks,
			})
		}
		err = j.gCache.Update(stopKey, resultByStop)
		if err != nil {
			j.logger.Errorw("could not update values in cache", "key_name", stopKey, "error", err.Error())
		}
	}
}

func (j *Job) RunBackgroundHTTPJob(ctx context.Context) error {

	j.logger.Info("Background httpJob is starting....")

	ticker := time.NewTicker(time.Duration(j.runInterval) * time.Second)

	for {
		select {
		case <-ticker.C:

			j.httpClientRunner()

		case <-ctx.Done():
			return nil
		}
	}
}
