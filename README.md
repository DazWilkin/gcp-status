# GCP Status Exporter

Converts Google's status dashboard into a series of `up` metrics (by services) for Prometheus consumption.

## Run

```bash
docker run \
--interactive --tty --rm \
ghcr.io/dazwilkin/gcp-status:67cdec5e09b6c53b3c39dca32ba7f6ca8edf73fa
```

## Notes

```console
Array.from(document.getElementsByClassName("service-status")).forEach(e => console.log(`${e.innerText}`))
```