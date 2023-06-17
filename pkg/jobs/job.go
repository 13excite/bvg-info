package jobs

import (
	"context"
	"time"

	"github.com/13excite/bvg-info/pkg/cache"
	"github.com/13excite/bvg-info/pkg/store"

	"go.uber.org/zap"
)

type httpClient interface {
	GetNearbyDepartes(int) (*store.Departures, error)
}

type Job struct {
	runInterval    int
	logger         *zap.SugaredLogger
	hClient        httpClient
	departuresList []string // TODO: should we pass it???
	gCache         cache.Cache
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
	// TODO: pass arg to the function
	stops, err := j.hClient.GetNearbyDepartes(733612)
	if err != nil {
		j.logger.Error("httpJob failed. Got error from client", "error", err)

	}

	resultByStop := []store.CachedStop{}
	for _, stop := range stops.Departures {
		resultByStop = append(resultByStop, store.CachedStop{
			ID:          stop.Stop.ID,
			Name:        stop.Stop.Name,
			Time:        stop.When,
			PlannedTime: stop.PlannedWhen,
			Direction:   stop.Direction,
			LineName:    stop.Line.Name,
			ProductName: stop.Line.ProductName,
		})
	}

	j.gCache.Update(store.Sudostallee_Kongisheide, resultByStop)

	//fmt.Println(stops)
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
