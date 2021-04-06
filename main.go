package main

import (
	"flag"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	dashboard = "https://status.cloud.google.com"
)

var (
	// GitCommit is the git commit value and is expected to be set during build
	GitCommit string
	// GoVersion is the Golang runtime version
	GoVersion = runtime.Version()
	// OSVersion is the OS version (uname --kernel-release) and is expected to be set during build
	OSVersion string
	// StartTime is the start time of the exporter represented as a UNIX epoch
	StartTime = time.Now().Unix()
)

var (
	endpoint    = flag.String("endpoint", ":9989", "The endpoint of the HTTP server")
	metricsPath = flag.String("path", "/metrics", "The path on which Prometheus metrics will be served")
)

func handleHealthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
func main() {
	flag.Parse()

	if GitCommit == "" {
		log.Println("[main] GitCommit value unchanged: expected to be set during build")
	}
	if OSVersion == "" {
		log.Println("[main] OSVersion value unchanged: expected to be set during build")
	}

	registry := prometheus.NewRegistry()
	registry.MustRegister(NewExporterCollector(OSVersion, GoVersion, GitCommit, StartTime))
	registry.MustRegister(NewStatusCollector())

	mux := http.NewServeMux()
	mux.Handle("/healthz", http.HandlerFunc(handleHealthz))
	mux.Handle(*metricsPath, promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))

	log.Printf("[main] Server starting (%s)", *endpoint)
	log.Printf("[main] metrics served on: %s", *metricsPath)
	log.Fatal(http.ListenAndServe(*endpoint, mux))
}
