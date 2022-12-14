package jobs

import (
	"context"
	"fmt"
	"time"

	"github.com/13excite/bvg-info/pkg/store"
	"go.uber.org/zap"
)

type httpClient interface {
	GetNearbyDepartes(int) ([]store.StopDepartures, error)
}

type Job struct {
	runInterval    int
	logger         *zap.SugaredLogger
	hClient        httpClient
	departuresList []string // TODO: should we pass it???
}

func New(runInterval int, httphttpClient httpClient) *Job {
	return &Job{
		runInterval: runInterval,
		logger:      zap.S().With("package", "job"),
		hClient:     httphttpClient,
	}
}

func (j *Job) httpClientRunner() {
	// TODO: pass arg to the function
	stops, err := j.hClient.GetNearbyDepartes(123)
	if err != nil {
		j.logger.Error("httpJob failed. Got error from client", "error", err)

	}
	fmt.Println(stops)
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
