package cmd

import (
	"context"

	cli "github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/13excite/bvg-info/pkg/api"
	"github.com/13excite/bvg-info/pkg/bvv"
	"github.com/13excite/bvg-info/pkg/cache"
	"github.com/13excite/bvg-info/pkg/conf"
	"github.com/13excite/bvg-info/pkg/jobs"
)

func init() {
	rootCmd.AddCommand(apiCmd)
}

var (
	apiCmd = &cli.Command{
		Use:   "api",
		Short: "Start API",
		Long:  `Start API`,
		Run: func(cmd *cli.Command, args []string) { // Initialize the database
			conf.C.Defaults()
			if configFile != "" {
				conf.C.ReadConfigFile(configFile)
			}
			conf.InitLogger(&conf.C)

			logger = zap.S().With("package", "cmd")

			// Cache
			cache := cache.NewGCache()

			httpClient := bvv.NewClent(conf.C.VBB.API)

			// Create the server
			s := api.New(&conf.C, cache)
			s.Router().HandleFunc("/test", s.GetData).Methods("GET")

			if err := s.ListenAndServe(&conf.C); err != nil {
				logger.Fatalw("Could not start server", "error", err)
			}

			job := jobs.New(conf.C.VBB.ScanSecInterval, httpClient, cache)
			err := job.RunBackgroundHTTPJob(context.TODO())
			if err != nil {
				logger.Fatalw("Could not start background job", "error", err)
			}

			// conf.S.InitInterrupt()
			// <-conf.Stop.Chan() // Wait until Stop
			// conf.Stop.Wait()   // Wait until everyone cleans up
			_ = zap.L().Sync() // Flush the logger

		},
	}
)
