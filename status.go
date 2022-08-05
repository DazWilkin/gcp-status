package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

// StatusCollector implements prometheus.Collector
var _ prometheus.Collector = (*StatusCollector)(nil)

// StatusCollector represents GCP status dashboard
type StatusCollector struct {
	// Prometheus Metric representing the number of GCP services parsed
	Services *prometheus.Desc
	// Prometheus Metric representing the status of each GCP service
	Up *prometheus.Desc
}

// NewStatusCollector returns a new StatusCollector
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
				"region",
			},
			nil,
		),
	}
}

// Collect implements Prometheus' Collector interface and is used to collect metrics
func (c *StatusCollector) Collect(ch chan<- prometheus.Metric) {
	dashboard := Dashboard{}
	services := dashboard.Parse()

	ch <- prometheus.MustNewConstMetric(
		c.Services,
		prometheus.GaugeValue,
		float64(len(services)),
		[]string{}...,
	)

	// Each Google Cloud service
	for _, service := range services {
		// Comprises a status for a combination of Google Regions
		// e.g. Americas, Multi-regions, Global
		for region, up := range service.Regions {
			ch <- prometheus.MustNewConstMetric(
				c.Up,
				prometheus.GaugeValue,
				func(up bool) float64 {
					if up {
						return 1.0
					}
					return 0.0
				}(up),
				[]string{
					service.Name,
					region.String(),
				}...,
			)
		}
	}

}

// Describe implements Prometheus' Collector interface and is used to describe metrics
func (c *StatusCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.Services
	ch <- c.Up
}
