# GCP Status Exporter

[![build-container](https://github.com/DazWilkin/gcp-status/actions/workflows/build-container.yml/badge.svg)](https://github.com/DazWilkin/gcp-status/actions/workflows/build-container.yml)

Converts Google's status dashboard into a series of `up` metrics (by services) for Prometheus consumption.

```console
# HELP services Count of GCP services
# TYPE services gauge
services 41
# HELP up Status of GCP service
# TYPE up gauge
up{service="Apigee"} 1
up{service="Cloud Asset Inventory"} 1
up{service="Cloud Data Fusion"} 1
up{service="Cloud Developer Tools"} 1
up{service="Cloud Endpoints"} 1
up{service="Cloud Filestore"} 1
up{service="Cloud Firestore"} 1
up{service="Cloud Key Management Service"} 1
up{service="Cloud Machine Learning"} 1
up{service="Cloud Memorystore"} 1
up{service="Cloud Run"} 1
up{service="Cloud Security Command Center"} 1
up{service="Cloud Spanner"} 1
up{service="Cloud Talent Solution - Job Search"} 1
up{service="Cloud Workflows"} 1
up{service="Eventarc"} 1
up{service="Google App Engine"} 1
up{service="Google BigQuery"} 1
up{service="Google Cloud Bigtable"} 1
up{service="Google Cloud Composer"} 1
up{service="Google Cloud Console"} 1
up{service="Google Cloud DNS"} 1
up{service="Google Cloud Dataflow"} 1
up{service="Google Cloud Dataproc"} 1
up{service="Google Cloud Datastore"} 1
up{service="Google Cloud Functions"} 1
up{service="Google Cloud Infrastructure Components"} 1
up{service="Google Cloud IoT"} 1
up{service="Google Cloud Networking"} 1
up{service="Google Cloud Pub/Sub"} 1
up{service="Google Cloud SQL"} 1
up{service="Google Cloud Scheduler"} 1
up{service="Google Cloud Storage"} 1
up{service="Google Cloud Support"} 1
up{service="Google Cloud Tasks"} 1
up{service="Google Compute Engine"} 1
up{service="Google Kubernetes Engine"} 1
up{service="Healthcare and Life Sciences"} 1
up{service="Identity and Access Management"} 1
up{service="Operations"} 1
up{service="Secret Manager"} 1
```

## Run

```bash
docker run \
--interactive --tty --rm \
ghcr.io/dazwilkin/gcp-status:be5308701567385d30ec6d57b54da1dd7b94cd1c
```

## Notes

```console
Array.from(document.getElementsByClassName("service-status")).forEach(e => console.log(`${e.innerText}`))
```