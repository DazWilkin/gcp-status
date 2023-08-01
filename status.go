package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	url string = "https://status.cloud.google.com"
)

// StatusCollector implements prometheus.Collector
var _ prometheus.Collector = (*StatusCollector)(nil)

// StatusCollector represents GCP status dashboard
type StatusCollector struct {
	// Prometheus Metric representing the number of Google Cloud services available
	Regions *prometheus.Desc
	// Prometheus Metric representing the total number of Google Cloud services
	Services *prometheus.Desc
	// Prometheus Metric representing the status of each Google Cloud service
	Up *prometheus.Desc
}

// NewStatusCollector returns a new StatusCollector
func NewStatusCollector() *StatusCollector {
	fqName := name("status")
	return &StatusCollector{
		Services: prometheus.NewDesc(
			fqName("services_total"),
			"Count of GCP services",
			[]string{},
			nil,
		),
		Regions: prometheus.NewDesc(
			fqName("services"),
			"Count of GCP service availability",
			[]string{
				"region",
			},
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
	dashboard := NewDashboard(url)
	node, err := dashboard.Open()
	if err != nil {
		return
	}

	services := parse(node)

	// Service totals
	ch <- prometheus.MustNewConstMetric(
		c.Services,
		prometheus.GaugeValue,
		float64(len(services)),
		[]string{}...,
	)
	// Service counts by region
	byRegion := services.ByRegion()
	for r := Americas; r <= Global; r++ {
		if count, ok := byRegion[r]; ok {
			ch <- prometheus.MustNewConstMetric(
				c.Regions,
				prometheus.GaugeValue,
				float64(count),
				[]string{
					r.String(),
				}...,
			)
		}
	}

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
