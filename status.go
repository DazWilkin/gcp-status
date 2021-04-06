package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/net/html"
)

// StatusCollector implements prometheus.Collector
var _ prometheus.Collector = (*StatusCollector)(nil)

type StatusCollector struct {
	// Prometheus Metric representing the number of GCP services parsed
	Services *prometheus.Desc
	// Prometheus Metric representing the status of each GCP service
	Up *prometheus.Desc
}

func NewStatusCollector() *StatusCollector {
	fqName := name("status")
	return &StatusCollector{
		Services: prometheus.NewDesc(
			fqName("services"),
			"Count of GCP services",
			[]string{},
			nil,
		),

		Up: prometheus.NewDesc(
			fqName("up"),
			"Status of GCP service (1=Available; 0=Unavailable)",
			[]string{
				"service",
			},
			nil,
		),
	}
}
func (c *StatusCollector) Collect(ch chan<- prometheus.Metric) {
	resp, err := http.Get(dashboard)
	if err != nil {
		log.Fatal("Unable to GET status dashboard")
	}
	// defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal("Unable to HTML parse body")
	}
	services := extractServices(doc)
	ch <- prometheus.MustNewConstMetric(
		c.Services,
		prometheus.GaugeValue,
		float64(len(services)),
		[]string{}...,
	)
	for _, service := range services {
		ch <- prometheus.MustNewConstMetric(
			c.Up,
			prometheus.GaugeValue,
			service.Up,
			[]string{
				service.Name,
			}...,
		)
	}

}
func (c *StatusCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.Services
	ch <- c.Up
}
