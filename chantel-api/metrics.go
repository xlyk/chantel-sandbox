package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
)

var (
	DateRequestCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "chantel_api_date_request_counter",
		Help: "Total number of API requests to the date endpoint",
	})

	ErrorCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "chantel_api_error_counter",
		Help: "Total number of API errors.",
	})
)

func ServeMetrics() {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Info().Msg("serving prometheus metrics on port 1234")

		server := &http.Server{
			Addr:              ":1234",
			ReadHeaderTimeout: 3 * time.Second,
		}

		if err := server.ListenAndServe(); err != nil {
			log.Error().Err(err).Msg("")
		}
	}()
}
