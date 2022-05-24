package command

import (
	"log"
	"time"

	"github.com/spf13/cobra"

	prom "github.com/prometheus/client_golang/prometheus"
	tally "github.com/uber-go/tally/v4"
	prometheus "github.com/uber-go/tally/v4/prometheus"
	tallyhandler "go.temporal.io/sdk/contrib/tally"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/MatthiasScholz/temporal-tfsecret/internal/pkg/temporal"

	// TODO Implement
	//"github.com/MatthiasScholz/temporal-tfsecret/activities"

	"github.com/MatthiasScholz/temporal-tfsecret/internal/pkg/workflows"
)

// workerCmd represents the worker command, running in the cluster
var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Run worker",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := temporal.NewClient(client.Options{
			MetricsHandler: tallyhandler.NewMetricsHandler(newPrometheusScope(prometheus.Configuration{
				ListenAddress: "0.0.0.0:8001",
				TimerType:     "histogram",
			})),
		})
		if err != nil {
			log.Fatalf("client error: %v", err)
		}
		defer c.Close()

		w := worker.New(c, workflows.TFSecretTaskQueue, worker.Options{})
		workflows.Register(w)

		err = w.Run(worker.InterruptCh())
		if err != nil {
			log.Fatalf("worker exited: %v", err)
		}
	},
}

func newPrometheusScope(c prometheus.Configuration) tally.Scope {
	reporter, err := c.NewReporter(
		prometheus.ConfigurationOptions{
			Registry: prom.NewRegistry(),
			OnError: func(err error) {
				log.Println("error in prometheus reporter", err)
			},
		},
	)
	if err != nil {
		log.Fatalln("error creating prometheus reporter", err)
	}
	scopeOpts := tally.ScopeOptions{
		CachedReporter: reporter,
		Separator:      prometheus.DefaultSeparator,
	}
	scope, _ := tally.NewRootScope(scopeOpts, time.Second)

	return scope
}

func init() {
	rootCmd.AddCommand(workerCmd)
}
