package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	messagebird "github.com/messagebird/go-rest-api"
	"github.com/messagebird/go-rest-api/balance"

	"github.com/alecthomas/kingpin/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/exporter-toolkit/web"
	webflag "github.com/prometheus/exporter-toolkit/web/kingpinflag"

	"github.com/prometheus/common/promslog"
	"github.com/prometheus/common/version"
)

var (
	metricsPath = kingpin.Flag(
		"web.telemetry-path",
		"Path under which to expose metrics.",
	).Default("/metrics").String()
	toolkitFlags = webflag.AddFlags(kingpin.CommandLine, ":9601")
	apiKey       = kingpin.Flag(
		"messagebird.api-key",
		"MessageBird Api Key.",
	).Required().String()

	credits = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "messagebird_balance_credits",
		Help: "The total number of credits remaining in the MessageBird account",
	})

	client *messagebird.Client
	logger *slog.Logger
)

func recordMetrics() {
	go func() {
		for {
			// Request the balance information.
			balance, err := balance.Read(client)
			if err != nil {
				switch errResp := err.(type) {
				case messagebird.ErrorResponse:
					for _, err := range errResp.Errors {
						logger.Error("Error retrieving balance from the MessageBird account", "err", err)
					}
				}
			} else {
				credits.Set(float64(balance.Amount))
			}

			time.Sleep(10 * time.Second)
		}
	}()
}

func main() {
	promslogConfig := &promslog.Config{}
	kingpin.Version(version.Print("messagebird_exporter"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
	logger = promslog.New(promslogConfig)

	// Initialize a new MessageBird client using the provided API key.
	client = messagebird.New(*apiKey)

	logger.Info("Starting messagebird_exporter", "version", version.Info())
	logger.Info("operational information", "build_context", version.BuildContext())

	recordMetrics()

	// Define the metrics page based on the metrics path.
	http.Handle(*metricsPath, promhttp.Handler())

	// Define the landing page if the metrics path is not set to the root.
	if *metricsPath != "/" {
		landingConfig := web.LandingConfig{
			Name:        "MessageBird Exporter",
			Description: "MessageBird Exporter",
			Version:     version.Info(),
			Links: []web.LandingLinks{
				{
					Address: *metricsPath,
					Text:    "Metrics",
				},
			},
		}
		landingPage, err := web.NewLandingPage(landingConfig)
		if err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
		http.Handle("/", landingPage)
	}

	// Start the HTTP server.
	srv := &http.Server{}
	if err := web.ListenAndServe(srv, toolkitFlags, logger); err != nil {
		logger.Error("Error starting HTTP server", "err", err)
		os.Exit(1)
	}
}
