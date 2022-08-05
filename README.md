# GCP Status Exporter

[![build-container](https://github.com/DazWilkin/gcp-status/actions/workflows/build-container.yml/badge.svg)](https://github.com/DazWilkin/gcp-status/actions/workflows/build-container.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/DazWilkin/gcp-status.svg)](https://pkg.go.dev/github.com/DazWilkin/gcp-status)
[![Go Report Card](https://goreportcard.com/badge/github.com/dazwilkin/gcp-status)](https://goreportcard.com/report/github.com/dazwilkin/gcp-status)

Converts Google [Cloud Status Dashboard](https://status.cloud.google.com/) into a series of `up` metrics (by services) for Prometheus consumption.

```console
# HELP gcp_status_build_info A metric with a constant '1' value labeled by OS version, Go version, and the Git commit of the exporter
# TYPE gcp_status_build_info counter
gcp_status_build_info{git_commit="",go_version="go1.18",os_version=""} 1
# HELP gcp_status_services Count of GCP service availability
# TYPE gcp_status_services gauge
gcp_status_services{region="Americas"} 94
gcp_status_services{region="Asia Pacific"} 85
gcp_status_services{region="Europe"} 95
gcp_status_services{region="Global"} 61
gcp_status_services{region="Multi-regions"} 21
# HELP gcp_status_services_total Count of GCP services
# TYPE gcp_status_services_total gauge
gcp_status_services_total 138
# HELP gcp_status_start_time Exporter start time in Unix epoch seconds
# TYPE gcp_status_start_time gauge
gcp_status_start_time 1.659727534e+09
# HELP gcp_status_up Status of GCP service (1=Available; 0=Unavailable)
# TYPE gcp_status_up gauge
gcp_status_up{region="Americas",service="Artifact Registry"} 1
gcp_status_up{region="Asia Pacific",service="Artifact Registry"} 1
gcp_status_up{region="Europe",service="Artifact Registry"} 1
gcp_status_up{region="Multi-regions",service="Artifact Registry"} 1
gcp_status_up{region="Americas",service="Google Kubernetes Engine"} 1
gcp_status_up{region="Asia Pacific",service="Google Kubernetes Engine"} 1
gcp_status_up{region="Europe",service="Google Kubernetes Engine"} 1
```

## Metrics

All metric names are prefix `gcp_status_`

|Name|Type|Labels|Description|
|----|----|------|-----------|
|`build_info`|Counter|`git_commit`,`go_version`,`os_version`|The status of the Exporter (1=available)|
|`services`|Gauge|`region`|The count of Google Cloud services by region| 
|`services_total`|Gauge||The count of Google Cloud services|
|`up`|Gauge|`region`,`service`|The status of the `service` in the `region` (1=available;0=down)|

## Run

### Go

```bash
go run .
```

### Docker

```bash
docker run \
--interactive --tty --rm \
ghcr.io/dazwilkin/gcp-status:fb358c9f3cab26a481ceca9698beea595b9ff459 \
--endpoint=:9989 \
--path=/metrics
```

### Docker Compose

```YAML
gcp-exporter:
  image: ghcr.io/dazwilkin/gcp-status:fb358c9f3cab26a481ceca9698beea595b9ff459
  container_name: gcp-status
  expose:
  - "9989" # GCP Status port registered on Prometheus Wiki
  ports:
  - 9989:9989
```

### Kubernetes

```YAML

```

## Notes

```console
Array.from(document.getElementsByClassName("service-status")).forEach(e => console.log(`${e.innerText}`))
```