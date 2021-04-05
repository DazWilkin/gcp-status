ARG GOLANG_VERSION=1.16
ARG PROJECT="gcp-status"

FROM golang:${GOLANG_VERSION} as build

ARG VERSION=""
ARG COMMIT=""

ARG PROJECT

WORKDIR /${PROJECT}

COPY go.* ./
COPY *.go .

RUN CGO_ENABLED=0 GOOS=linux \
    go build \
    -ldflags "-X main.OSVersion=${VERSION} -X main.GitCommit=${COMMIT}" \
    -a -installsuffix cgo \
    -o /go/bin/${PROJECT} \
    .

FROM scratch

ARG PROJECT

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /go/bin/${PROJECT} /exporter

EXPOSE 8080

ENTRYPOINT ["/exporter"]
