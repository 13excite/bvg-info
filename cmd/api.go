package cmd

import (
	cli "github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/13excite/bvg-info/pkg/conf"

	// "github.com/13excite/bvg-info/pkg/store"
	"github.com/13excite/bvg-info/pkg/bvv"
	"github.com/13excite/bvg-info/pkg/cache"
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
			if &configFile != nil && configFile != "" {
				conf.C.ReadConfigFile(configFile)
			}
			conf.InitLogger(&conf.C)

			logger = zap.S().With("package", "cmd")

			// Cache
			gcache := cache.NewGCache()

			// Create job and httpClient
			hClient := bvv.NewClent(conf.C.VBB.API)
			job := jobs.New(conf.C.VBB.ScanSecInterval, hClient)


		},
	}
)
