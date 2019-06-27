package main

import (
	"fmt"
	"net/http"
	"time"

	messagebird "github.com/messagebird/go-rest-api"
	"github.com/messagebird/go-rest-api/balance"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"gopkg.in/alecthomas/kingpin.v2"

	log "github.com/sirupsen/logrus"
)

var (
	credits = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "messagebird_balance_credits",
		Help: "The total number of credits remaining in the MessageBird account",
	})

	client *messagebird.Client
)

func recordMetrics() {
	go func() {
		for {
			// Request the balance information
			balance, err := balance.Read(client)
			if err != nil {
				switch errResp := err.(type) {
				case messagebird.ErrorResponse:
					for _, mbError := range errResp.Errors {
						log.Error(fmt.Sprintf("%#v", mbError))
					}
				}
			}

			credits.Set(float64(balance.Amount))

			time.Sleep(10 * time.Second)
		}
	}()

}

func main() {
	var (
		listenAddress = kingpin.Flag(
			"web.listen-address",
			"Address on which to expose metrics and web interface.",
		).Default(":9601").String()
		metricsPath = kingpin.Flag(
			"web.telemetry-path",
			"Path under which to expose metrics.",
		).Default("/metrics").String()
		apiKey = kingpin.Flag(
			"messagebird.api-key",
			"MessageBird Api Key.",
		).Required().String()
	)

	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	// Initialize a new MessageBird client.
	client = messagebird.New(*apiKey)

	registry := prometheus.NewRegistry()
	registry.MustRegister(credits)
	handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})

	log.WithFields(log.Fields{
		"version": "unknown",
	}).Info("Starting MessageBird Exporter")

	recordMetrics()

	http.Handle(*metricsPath, handler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
				<head><title>MessageBird Exporter</title></head>
				<body>
				<h1>MessageBird Exporter</h1>
				<p><a href="` + *metricsPath + `">Metrics</a></p>
				</body>
				</html>`))
	})

	log.WithFields(log.Fields{
		"listen_address": *listenAddress,
		"metrics_path":   *metricsPath,
	}).Info("Listening on " + *listenAddress)
	if err := http.ListenAndServe(*listenAddress, nil); err != nil {
		log.Fatal(err)
	}
}
